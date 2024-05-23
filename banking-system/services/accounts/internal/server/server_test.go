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

func createBufconnDialer(l *bufconn.Listener) func(ctx context.Context, addr string) (net.Conn, error) {
	return func(ctx context.Context, addr string) (net.Conn, error) {
		return l.Dial()
	}
}

func TestAccountCreateService(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefulStop()

	bufconnDialer := createBufconnDialer(l)

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

	bufconnDialer := createBufconnDialer(l)

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

func TestAccountUpdateService(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefulStop()

	bufconnDialer := createBufconnDialer(l)

	client, err := grpc.DialContext(
		context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(bufconnDialer),
	)

	if err != nil {
		t.Fatal(err)
	}

	accountClient := pb.NewAccountServiceClient(client)
	resp, err := accountClient.UpdateAccount(
		context.Background(),
		&pb.UpdateAccountRequest{
			AccountId: "23a18029-0d16-11ef-a43f-9c6b003d2771",
			OwnerName: "John Doe Jr", // RowAffected 에러 주의. 테스트 실행마다 변경할 것
		},
	)

	if err != nil {
		t.Fatal(err)
	}
	if resp.Account.OwnerName != "John Doe Jr" { // RowAffected 에러 주의. 테스트 실행마다 변경할 것
		t.Errorf(
			"expected Currency to be: John Doe Jr, Got: %v",
			resp.Account.OwnerName,
		)
	}
}

func TestAccountDeleteService(t *testing.T) {
	s, l := startTestGrpcServer()
	defer s.GracefulStop()

	bufconnDialer := createBufconnDialer(l)
	client, err := grpc.DialContext(
		context.Background(), "", grpc.WithInsecure(), grpc.WithContextDialer(bufconnDialer),
	)
	if err != nil {
		t.Fatal(err)
	}

	DeleteTestID := "23a18029-0d16-11ef-a43f-9c6b003d2771" // Test마다 설정하기

	accountClient := pb.NewAccountServiceClient(client)
	resp, err := accountClient.DeleteAccount(
		context.Background(),
		&pb.DeleteAccountRequest{
			AccountId: DeleteTestID,
		},
	)

	if err != nil {
		t.Fatal(err)
	}
	if resp.AccountId != DeleteTestID {
		t.Errorf(
			"expected Deleted AccountID to be: %s, Got: %v",
			DeleteTestID,
			resp.AccountId,
		)
	}
}

// type AccountDB interface {
// 	CreateAccount(ctx context.Context, ownerName string, initialDeposit float64, currency string) (*Account, error)
// }

// type MockAccountDB struct {
// 	mock.Mock
// }

// func (m *MockAccountDB) CreateAccount(ctx context.Context, ownerName string, initialDeposit float64, currency string) (*pb.AccountInfo, error) {
// 	args := m.Called(ctx, ownerName, initialDeposit, currency)
// 	return args.Get(0).(*pb.AccountInfo), args.Error(1)
// }

// func startTestGrpcServer_Mock() (*grpc.Server, *bufconn.Listener) {
// 	l := bufconn.Listen(10)
// 	s := grpc.NewServer()
// 	mockAccountDB := &MockAccountDB{}
// 	mockAccountDB.On("CreateAccount", "John Doe", 1000.0, "KRW").Return(&pb.AccountInfo{OwnerName: "John Doe", Balance: 1000.0, Currency: "KRW"}, nil)

// 	pb.RegisterAccountServiceServer(s, &NewAccountServer{}) // MockDB 용으로 Server구조체가 필요함.

// 	go func() {
// 		err := s.Serve(l)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}()
// 	return s, l
// }
