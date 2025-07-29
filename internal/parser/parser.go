package parser

import (
	"context"

	pb "github.com/sproutbro/parserx/proto/pbparser"
	"google.golang.org/grpc"
)

type SVC interface {
	YAML(ctx context.Context, req *pb.Req) *pb.Resp
	JSON(ctx context.Context, req *pb.Req) *pb.Resp
	ENV(ctx context.Context, req *pb.Req) *pb.Resp
}

type svc struct {
	pb.UnimplementedParserxServer
}

func New(s *grpc.Server) *svc {
	if s == nil {
		return &svc{}
	}
	pb.RegisterParserxServer(s, &svc{})
	return nil
}
