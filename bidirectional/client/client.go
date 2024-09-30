package main

import (
	"context"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"grpc/bidirectional/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)

	stream, err := client.BidirectionalStreamMessages(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	messages := []string{"Message 1", "Message 2", "Message 3"}

	go func() {
		for _, msg := range messages {
			if err := stream.Send(&pb.MessageRequest{Message: msg}); err != nil {
				log.Fatalf("Error sending message: %v", err)
			}
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
		log.Printf("Server replied: %s\n", res.Reply)
	}
}
