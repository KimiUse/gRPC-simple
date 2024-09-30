package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"grpc/client/pb"
)

type MessageServiceServer struct {
	pb.UnimplementedMessageServiceServer
}

func (s *MessageServiceServer) ClientStreamMessages(stream pb.MessageService_ClientStreamMessagesServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageResponse{Reply: "Received: " + fmt.Sprint(messages)})
		}
		if err != nil {
			return err
		}
		messages = append(messages, req.Message)
		fmt.Printf("Received message: %s\n", req.Message)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessageServiceServer(grpcServer, &MessageServiceServer{})

	fmt.Println("Client Streaming gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
