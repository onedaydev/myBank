package server

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/onedaydev/myBank/banking-system/services/alarm/api"
	"github.com/onedaydev/myBank/banking-system/services/alarm/internal/kafka"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

type AlarmServer struct {
	pb.UnimplementedAlarmServiceServer
	producer *kafka.Producer
}

func NewAlarmServer(producer *kafka.Producer) *AlarmServer {
	return &AlarmServer{producer: producer}
}

func (s *AlarmServer) GetAlarm(req *pb.AlarmGetRequest, stream pb.AlarmService_GetAlarmServer) error {
	for {
		alarm := &pb.AlarmGetReply{
			AlarmId:   generateID(),
			UserId:    req.UserId,
			EventType: randomEventType(req.EventTypes),
			Message:   "This is a test alarm",
			Timestamp: time.Now().Unix(),
		}

		err := s.producer.SendMessage("alarms", fmt.Sprintf("%v", alarm))
		if err != nil {
			return fmt.Errorf("failed to send message to Kafka: %v", err)
		}

		if err := stream.Send(alarm); err != nil {
			return err
		}

		time.Sleep(time.Second * 2)
	}
}

func generateID() string {
	return fmt.Sprintf("%d", rand.Int())
}

func randomEventType(eventTypes []string) string {
	if len(eventTypes) == 0 {
		return "default"
	}
	return eventTypes[rand.Intn(len(eventTypes))]
}

func StartServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	producer, err := kafka.NewProducer([]string{"localhost:9092"})
	if err != nil {
		log.Fatalf("failed to start Kafka producer: %v", err)
	}
	defer producer.Close()

	s := grpc.NewServer()
	pb.RegisterAlarmServiceServer(s, NewAlarmServer(producer))
	reflection.Register(s)

	log.Printf("Server is running on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failded to serve: %v", err)
	}
}
