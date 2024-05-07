/banking-system          # 프로젝트 루트 디렉토리
|-- /services            # 각 마이크로서비스 디렉토리
|   |-- /accounts        # 계좌 관리 서비스
|   |   |-- /api         # gRPC 프로토콜 파일 및 생성된 소스 코드
|   |   |   |-- account.proto   # gRPC 프로토콜 정의 파일
|   |   |   `-- account.pb.go   # 프로토콜 버퍼 컴파일러에 의해 생성된 Go 소스 파일
|   |   |-- /cmd         # 애플리케이션 실행을 위한 메인 패키지
|   |   |   `-- main.go  # 서비스 시작을 위한 메인 애플리케이션 파일
|   |   `-- /internal    # 서비스 내부 구현
|   |       |-- /server  # gRPC 서버 구현
|   |       `-- /db      # 데이터베이스 접근 구현
|   |-- /transactions    # 거래 관리 서비스
|   |-- /credit          # 신용 평가 서비스
|   `-- /customers       # 고객 서비스
|-- /common              # 공통 라이브러리 및 유틸리티
|   |-- /lib             # 공통 라이브러리 파일
|   `-- /utils           # 공통 유틸리티 함수
|-- /deploy              # 쿠버네티스 배포 스크립트 및 매니페스트 파일
|   |-- k8s              # 쿠버네티스 YAML 파일
|   `-- scripts          # 배포 관련 스크립트
`-- README.md            # 프로젝트 설명 및 사용 가이드

github.com/onedaydev/myBank/banking-system/services/accounts/

protoc --go_out=. --go_opt=paths=source_relative \
--go-grpc_out=. --go-grpc_opt=paths=source_relative \
users.proto repositories.proto

go get google.golang.org/grpc@v1.37.0

replace github.com/onedaydev/myBank/banking-system/services/accounts/ => ../../api