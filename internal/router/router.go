package router

import (
	"github.com/sproutbro/parserx/internal/parser"
	"google.golang.org/grpc"
)

func NewParserRouter() *grpc.Server {
	// grpc 서버 생성
	s := grpc.NewServer()

	// 각서비스 등록
	parser.New(s)

	return s
}
