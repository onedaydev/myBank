# Banking-system

## Services
### Accounts
- 단방향 grpc : 계좌 생성, 계좌 조회, 계좌 소유자 명 변경, 계좌 삭제 구현
- Client 폴더 내 client 파일로 통합 테스트 완료
- Internal/Server 내 server_test파일로 단위 테스트 완료

### Chatting
- 양방향 스트리밍 구현

### Alarm
- 서버 스트리밍 구현

### File-Upload
- 클라이언트 스트리밍 구현

protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
tmp.proto
