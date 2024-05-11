package server

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	pb "github.com/onedaydev/myBank/banking-system/services/accounts/api"
	db "github.com/onedaydev/myBank/banking-system/services/accounts/internal/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountServer struct {
	pb.UnimplementedAccountServiceServer
	db *sql.DB
}

func NewAccountServer(db *sql.DB) *AccountServer {
	return &AccountServer{db: db}
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
	accountData := db.AccountData{
		AccountID: id.String(),
		OwnerName: req.OwnerName,
		Balance:   req.InitialDeposit,
		Currency:  req.Currency,
	}
	conn, err := db.ConnectToDB()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to connect to database: %v", err)
	}
	defer conn.Close()

	err = db.CreateAccount(conn, accountData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create account: %v", err)
	}

	return &pb.CreateAccountResponse{
		Account: &pb.AccountInfo{
			AccountId: accountData.AccountID,
			OwnerName: accountData.OwnerName,
			Balance:   accountData.Balance,
			Currency:  accountData.Currency,
		},
	}, nil
}

func (s *AccountServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	conn, err := db.ConnectToDB()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to connect to database: %v", err)
	}
	account, err := db.GetRow(conn, req.AccountId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Account not found: %v", err)
	}
	return &pb.GetAccountResponse{
		Account: &pb.AccountInfo{
			AccountId: account.AccountID,
			OwnerName: account.OwnerName,
			Balance:   account.Balance,
			Currency:  account.Currency,
		},
	}, nil
}

func (s *AccountServer) UpdateAccount(ctx context.Context, req *pb.UpdateAccountRequest) (*pb.UpdateAccountResponse, error) {
	conn, err := db.ConnectToDB()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to connect to database: %v", err)
	}
	account, err := db.UpdateOwnerName(conn, req.AccountId, req.OwnerName)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Account not found: %v", err)
	}
	return &pb.UpdateAccountResponse{
		Account: &pb.AccountInfo{
			AccountId: account.AccountID,
			OwnerName: account.OwnerName,
			Balance:   account.Balance,
			Currency:  account.Currency,
		},
	}, nil
}

// func (s *AccountServer) UpdateBalanceAccount(ctx context.Context, req *pb.UpdateAccountBalanceRequest) (*pb.UpdateAccountBalanceResponse, error) {
// 	conn, err := db.ConnectToDB()
// 	if err != nil {
// 		return nil, status.Errorf(codes.Internal, "failed to connect to database: %v", err)
// 	}
// 	account, err := db.UpdateRow(conn, req.AccountId, req.Balance)
// 	if err != nil {
// 		return nil, status.Errorf(codes.NotFound, "Account not found: %v", err)
// 	}
// 	return &pb.UpdateAccountBalanceResponse{
// 		Account: &pb.AccountInfo{
// 			AccountId: account.AccountID,
// 			OwnerName: account.OwnerName,
// 			Balance:   account.Balance,
// 			Currency:  account.Currency,
// 		},
// 	}, nil
// }

func (s *AccountServer) DeleteAccount(ctx context.Context, req *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	conn, err := db.ConnectToDB()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to connect to database: %v", err)
	}
	err = db.DeleteRow(conn, req.AccountId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Account not found: %v", err)
	}
	return &pb.DeleteAccountResponse{
		AccountId: req.AccountId,
	}, nil
}
