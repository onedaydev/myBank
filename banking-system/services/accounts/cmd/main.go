package main

import (
	"log"
	"net"

	pb "github.com/onedaydev/myBank/banking-system/services/accounts/api"
	db "github.com/onedaydev/myBank/banking-system/services/accounts/internal/db"
	"github.com/onedaydev/myBank/banking-system/services/accounts/internal/server"
	"google.golang.org/grpc"
)

const (
	port = ":500051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// DB 연결
	dbConn, err := db.ConnectToDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	s := grpc.NewServer()

	// 계좌 관리 서비스 등록
	pb.RegisterAccountServiceServer(s, server.NewAccountServer(dbConn))

	// gRPC 서버 리플렉션 활성화
	// reflection.Register(s)

	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
