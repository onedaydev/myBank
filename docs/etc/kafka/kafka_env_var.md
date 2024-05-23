docker-compose 내용 중 환경 변수 설정 부분에 대한 설명과 예시

#### Kafka 브로커가 클라이언트 및 다른 브로커에게 자신을 알리는 방법을 설정
KAFKA_ADVERTISED_LISTENERS: INSIDE://kafka:9093,OUTSIDE://localhost:9092
      
#### 각 리스너의 보안 프로토콜을 정의
KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: INSIDE:PLAINTEXT,OUTSIDE:PLAINTEXT
      
#### Kafka 브로커가 실제로 바인딩하고 리스닝할 네트워크 인터페이스와 포트를 설정
KAFKA_LISTENERS: INSIDE://0.0.0.0:9093,OUTSIDE://0.0.0.0:9092
      
#### Kafka 브로커 간 통신에 사용할 리스너를 지정
KAFKA_INTER_BROKER_LISTENER_NAME: INSIDE
      
#### Kafka 브로커가 연결할 Zookeeper 인스턴스를 지정
KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181