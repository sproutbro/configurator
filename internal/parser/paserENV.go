package parser

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// .env 파일 읽기
func (p *Manager) ENV(filename string) (map[string]interface{}, error) {
	envMap := make(map[string]interface{})

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
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
	return envMap, nil
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
