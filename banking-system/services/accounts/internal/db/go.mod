module github.com/onedaydev/myBank/banking-system/services/accounts/internal/db

go 1.22.1

require github.com/onedaydev/myBank/banking-system/services/accounts/api v0.0.0-00010101000000-000000000000

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/grpc v1.63.2 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)

replace github.com/onedaydev/myBank/banking-system/services/accounts/api => ../../api
