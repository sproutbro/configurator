package parser

import (
	"context"
	"encoding/json"
	"os"

	pb "github.com/sproutbro/parserx/proto/pbparser"
)

// (ctx context.Context, req *pbcfg.ConfigRequest) (*pbcfg.ConfigResponse, error)
func (p *svc) JSON(ctx context.Context, req *pb.Req) (*pb.Resp, error) {

	data, err := os.ReadFile(req.Key)
	if err != nil {
		return &pb.Resp{}, err
	}

	var config map[string]interface{}
	if err := json.Unmarshal(data, &config); err != nil {
		return &pb.Resp{}, err
	}

	return &pb.Resp{Json: string(data), Error: "err"}, nil
}
