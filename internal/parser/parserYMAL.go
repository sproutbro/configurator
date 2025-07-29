package parser

import (
	"context"
	"encoding/json"
	"os"

	pb "github.com/sproutbro/parserx/proto/pbparser"
	"gopkg.in/yaml.v3"
)

func (p *svc) YAML(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	data, err := os.ReadFile(req.Key)
	if err != nil {
		return &pb.Resp{}, err
	}

	var config map[string]interface{}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return &pb.Resp{}, err
	}

	jsonBytes, err := json.Marshal(config)
	if err != nil {
		return &pb.Resp{}, err
	}
	return &pb.Resp{Json: string(jsonBytes), Error: ""}, nil
}
