package main

import (
	"context"
	"log"

	pb "github.com/onedaydev/myBank/banking-system/services/alarm/api"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAlarmServiceClient(conn)

	req := &pb.AlarmGetRequest{
		UserId:     "user_test_1",
		EventTypes: []string{"event1", "event2"},
	}

	stream, err := c.GetAlarm(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get alarm: %v", err)
	}

	for {
		alarm, err := stream.Recv()
		if err != nil {
			log.Fatalf("error receiving alarm: %v", err)
		}
		log.Printf("Recived alarm: %v", alarm)
	}
}
