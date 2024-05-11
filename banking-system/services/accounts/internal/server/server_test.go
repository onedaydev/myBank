package server

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/onedaydev/myBank/banking-system/services/accounts/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
)

func startTestGrpcServer() (*grpc.Server, *bufconn.Listener) {
	l := bufconn.Listen(10)
	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, &AccountServer{})

	go func() {
		err := s.Serve(l)
		if err != nil {
			log.Fatal(err)
		}
	}()
	return s, l
}

func TestAccountCreateService(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefulStop()

	bufconnDialer := func(
		ctx context.Context, addr string,
	) (net.Conn, error) {
		return l.Dial()
	}

	client, err := grpc.DialContext(
		context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(bufconnDialer),
	)

	if err != nil {
		t.Fatal(err)
	}

	accountClient := pb.NewAccountServiceClient(client)
	resp, err := accountClient.CreateAccount(
		context.Background(),
		&pb.CreateAccountRequest{
			OwnerName:      "John Doe",
			InitialDeposit: 1000.0,
			Currency:       "KRW",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	if resp.Account.OwnerName != "John Doe" {
		t.Errorf(
			"expected Currency to be: John Doe, Got: %v",
			resp.Account.OwnerName,
		)
	}
	if resp.Account.Balance != 1000.0 {
		t.Errorf(
			"expected Currency to be: 1000.0, Got: %v",
			resp.Account.Balance,
		)
	}
	if resp.Account.Currency != "KRW" {
		t.Errorf(
			"expected Currency to be: KRW, Got: %v",
			resp.Account.Currency,
		)
	}
}

func TestAccountGetService(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefulStop()

	bufconnDialer := func(
		ctx context.Context, addr string,
	) (net.Conn, error) {
		return l.Dial()
	}

	client, err := grpc.DialContext(
		context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(bufconnDialer),
	)

	if err != nil {
		t.Fatal(err)
	}

	accountClient := pb.NewAccountServiceClient(client)
	resp, err := accountClient.GetAccount(
		context.Background(),
		&pb.GetAccountRequest{
			AccountId: "23a18029-0d16-11ef-a43f-9c6b003d2771",
		},
	)

	if err != nil {
		t.Fatal(err)
	}
	if resp.Account.Currency != "KRW" {
		t.Errorf(
			"expected Currency to be: KRW, Got: %v",
			resp.Account.Currency,
		)
	}
	if resp.Account.OwnerName != "John Doe" {
		t.Errorf(
			"expected Currency to be: John Doe, Got: %v",
			resp.Account.OwnerName,
		)
	}
	if resp.Account.Balance != 100000.00 {
		t.Errorf(
			"expected Currency to be: 100000.00, Got: %v",
			resp.Account.Balance,
		)
	}
}

// func TestAccountUpdateService(t *testing.T) {
// 	s, l := startTestGrpcServer()
// 	defer s.GracefulStop()

// 	bufconnDialer := func(
// 		ctx context.Context, addr string,
// 	) (net.Conn, error) {
// 		return l.Dial()
// 	}

// 	client, err := grpc.DialContext(
// 		context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(bufconnDialer),
// 	)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	accountClient := pb.NewAccountServiceClient(client)
// 	resp, err := accountClient.GetAccount(
// 		context.Background(),
// 		&pb.GetAccountRequest{
// 			AccountId: "",
// 		},
// 	)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if resp.Account.Currency != "10000.0" {
// 		t.Errorf(
// 			"expected Currency to be: 10000.0, Got: %v",
// 			resp.Account.Currency,
// 		)
// 	}
// 	if resp.Account.OwnerName != "10000.0" {
// 		t.Errorf(
// 			"expected Currency to be: 10000.0, Got: %v",
// 			resp.Account.OwnerName,
// 		)
// 	}
// 	if resp.Account.Balance != 10000.0 {
// 		t.Errorf(
// 			"expected Currency to be: 10000.0, Got: %v",
// 			resp.Account.Balance,
// 		)
// 	}
// }

// func TestAccountDeleteService(t *testing.T) {
// 	s, l := startTestGrpcServer()
// 	defer s.GracefulStop()

// 	bufconnDialer := func(
// 		ctx context.Context, addr string,
// 	) (net.Conn, error) {
// 		return l.Dial()
// 	}

// 	client, err := grpc.DialContext(
// 		context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(bufconnDialer),
// 	)

// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	accountClient := pb.NewAccountServiceClient(client)
// 	resp, err := accountClient.GetAccount(
// 		context.Background(),
// 		&pb.GetAccountRequest{
// 			AccountId: "",
// 		},
// 	)

// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if resp.Account.Currency != "10000.0" {
// 		t.Errorf(
// 			"expected Currency to be: 10000.0, Got: %v",
// 			resp.Account.Currency,
// 		)
// 	}

// }
