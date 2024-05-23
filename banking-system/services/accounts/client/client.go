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

	// create 요청 테스트
	createCtx, createCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer createCancel()

	createResp, err := c.CreateAccount(createCtx, &pb.CreateAccountRequest{
		OwnerName:      "John Doe",
		InitialDeposit: 100000,
		Currency:       "KRW",
	})
	if err != nil {
		log.Fatalf("could not create account: %v", err)
	}
	log.Printf("account created: %v", createResp.GetAccount())

	CreatedAccountID := createResp.Account.AccountId

	// get 요청 테스트
	getCtx, getCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer getCancel()
	getResp, err := c.GetAccount(getCtx, &pb.GetAccountRequest{
		AccountId: CreatedAccountID,
	})
	if err != nil {
		log.Fatalf("could not get account: %v", err)
	}
	log.Printf("account got: %v", getResp.GetAccount())

	// update 요청 테스트
	updateCtx, updateCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer updateCancel()
	updateResp, err := c.UpdateAccount(updateCtx, &pb.UpdateAccountRequest{
		AccountId: CreatedAccountID,
		OwnerName: "John Doe Jr",
	})
	if err != nil {
		log.Fatalf("could not update account: %v", err)
	}
	log.Printf("account updated: %v", updateResp.GetAccount())

	// delete 요청 테스트
	deleteCtx, deleteCancel := context.WithTimeout(context.Background(), time.Second*5)
	defer deleteCancel()
	deleteResp, err := c.DeleteAccount(deleteCtx, &pb.DeleteAccountRequest{
		AccountId: CreatedAccountID,
	})
	if err != nil {
		log.Fatalf("could not delete account: %v", err)
	}
	log.Printf("account deleted: %v", deleteResp.AccountId)
}
