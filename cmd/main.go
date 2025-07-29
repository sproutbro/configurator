package main

import (
	"log"
	"net"

	"github.com/sproutbro/parserx/internal/router"
)

func main() {

	listener, err := net.Listen("tcp", ":55554")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 각서비스 등록
	server := router.NewParserRouter()

	log.Println("🚀 gRPC server listening on :55554")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("서버 종료 중 오류: %v", err)
	}
}
