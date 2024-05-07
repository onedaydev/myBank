package server

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/onedaydev/myBank/banking-system/services/accounts/api"
	db "github.com/onedaydev/myBank/banking-system/services/accounts/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountServer struct {
	pb.UnimplementedAccountServiceServer
}

func NewAccountServer() *AccountServer {
	return &AccountServer{}
}

func (s *AccountServer) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (
	*pb.CreateAccountResponse, error,
) {
	if req.InitialDeposit < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Initial deposit cannot be negative")
	}
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to generate an account ID")
	}
	account := &pb.AccountInfo{
		AccountId: id.String(),
		OwnerName: req.OwnerName,
		Balance:   req.InitialDeposit,
		Currency:  req.Currency,
	}
	conn, err := db.ConnectToDB()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to connect to database: %v", err)
	}
	defer conn.Close()

	err = db.CreateAccount(conn, account)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account: %v", err)
	}

	return &pb.CreateAccountResponse{Account: account}, nil
}

// func (s *AccountServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
// 	account, err := fetchAccountFromDB(req.AccountId)
// 	if err != nil {
// 		return nil, status.Errorf(codes.NotFound, "Account not found: %v", err)
// 	}

// 	return &pb.GetAccountResponse{Account: account}, nil
// }

// func (s *AccountServer) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {

// }

// func (s *AccountServer) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {

// }
