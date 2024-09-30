package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"grpc/server/pb"
)

type MessageServiceServer struct {
	pb.UnimplementedMessageServiceServer
}

func (s *MessageServiceServer) StreamMessages(req *pb.MessageRequest, stream pb.MessageService_StreamMessagesServer) error {
	for i := 1; i <= 5; i++ {
		reply := fmt.Sprintf("Reply #%d to %s", i, req.Message)
		res := &pb.MessageResponse{Reply: reply}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(time.Second)
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMessageServiceServer(grpcServer, &MessageServiceServer{})

	fmt.Println("Server Streaming gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
