module github.com/onedaydev/myBank/banking-system/services/alarm/cmd

go 1.22.1

replace github.com/onedaydev/myBank/banking-system/services/alarm/api => ../api

replace github.com/onedaydev/myBank/banking-system/services/alarm/internal/server => ../internal/server

replace github.com/onedaydev/myBank/banking-system/services/alarm/internal/kafka => ../internal/kafka

require github.com/onedaydev/myBank/banking-system/services/alarm/internal/server v0.0.0-00010101000000-000000000000

require (
	github.com/IBM/sarama v1.43.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/eapache/go-resiliency v1.6.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20230731223053-c322873962e3 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hashicorp/errwrap v1.0.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-uuid v1.0.3 // indirect
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.7.6 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.4 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/klauspost/compress v1.17.8 // indirect
	github.com/onedaydev/myBank/banking-system/services/alarm/api v0.0.0-00010101000000-000000000000 // indirect
	github.com/onedaydev/myBank/banking-system/services/alarm/internal/kafka v0.0.0-00010101000000-000000000000 // indirect
	github.com/pierrec/lz4/v4 v4.1.21 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	golang.org/x/crypto v0.22.0 // indirect
	golang.org/x/net v0.24.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240318140521-94a12d6c2237 // indirect
	google.golang.org/grpc v1.64.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)
