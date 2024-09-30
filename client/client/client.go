package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"grpc/client/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)

	stream, err := client.ClientStreamMessages(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	messages := []string{"Hello", "from", "Client Streaming RPC"}

	for _, msg := range messages {
		if err := stream.Send(&pb.MessageRequest{Message: msg}); err != nil {
			log.Fatalf("Error sending message: %v", err)
		}
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving response: %v", err)
	}

	log.Printf("Server replied: %s\n", res.Reply)
}
