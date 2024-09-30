package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"grpc/bidirectional/pb"
)

type MessageServiceServer struct {
	pb.UnimplementedMessageServiceServer
}

func (s *MessageServiceServer) BidirectionalStreamMessages(stream pb.MessageService_BidirectionalStreamMessagesServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		fmt.Printf("Received message: %s\n", req.Message)
		res := &pb.MessageResponse{Reply: "Received: " + req.Message}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessageServiceServer(grpcServer, &MessageServiceServer{})

	fmt.Println("Bidirectional Streaming gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
