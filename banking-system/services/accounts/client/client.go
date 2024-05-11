package main

import (
	"context"
	"log"
	"time"

	pb "github.com/onedaydev/myBank/banking-system/services/accounts/api"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAccountServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	r, err := c.CreateAccount(ctx, &pb.CreateAccountRequest{
		OwnerName:      "John Doe",
		InitialDeposit: 100000,
		Currency:       "KRW",
	})
	if err != nil {
		log.Fatalf("could not create account: %v", err)
	}
	log.Printf("account created: %v", r.GetAccount())
}
