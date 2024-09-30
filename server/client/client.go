package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"grpc/server/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)

	req := &pb.MessageRequest{Message: "Server Streaming RPC"}
	stream, err := client.StreamMessages(context.Background(), req)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("Server replied: %s\n", res.Reply)
	}
}
