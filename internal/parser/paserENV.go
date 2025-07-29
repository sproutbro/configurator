package parser

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
	"strconv"
	"strings"

	pb "github.com/sproutbro/parserx/proto/pbparser"
)

// .env 파일 읽기
func (p *svc) ENV(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	envMap := make(map[string]interface{})

	file, err := os.Open(req.Key)
	if err != nil {
		return &pb.Resp{Json: "", Error: ""}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// 주석/빈줄 무시
		if strings.HasPrefix(line, "#") || strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			val := strings.TrimSpace(parts[1])
			envMap[key] = parseValue(val)
		}
	}

	decode, err := json.Marshal(envMap)
	if err != nil {
		return &pb.Resp{Json: "", Error: ""}, err
	}

	return &pb.Resp{Json: string(decode), Error: ""}, nil
}

// .env 파일 타입추론
func parseValue(val string) interface{} {
	// 자동 타입 추론
	if val == "true" || val == "false" {
		return val == "true"
	}
	if i, err := strconv.Atoi(val); err == nil {
		return i
	}
	if f, err := strconv.ParseFloat(val, 64); err == nil {
		return f
	}
	return val
}
