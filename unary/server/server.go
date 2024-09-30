package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc/unary/pb"
)

type MessageServiceServer struct {
	pb.UnimplementedMessageServiceServer
}

func (s *MessageServiceServer) SendMessage(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	fmt.Printf("Received message: %s\n", req.Message)
	return &pb.MessageResponse{Reply: "Hello, " + req.Message}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessageServiceServer(grpcServer, &MessageServiceServer{})

	fmt.Println("Unary gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
