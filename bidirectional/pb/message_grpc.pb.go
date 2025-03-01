// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: message.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MessageService_BidirectionalStreamMessages_FullMethodName = "/pb.MessageService/BidirectionalStreamMessages"
)

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessageServiceClient interface {
	BidirectionalStreamMessages(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[MessageRequest, MessageResponse], error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) BidirectionalStreamMessages(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[MessageRequest, MessageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MessageService_ServiceDesc.Streams[0], MessageService_BidirectionalStreamMessages_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[MessageRequest, MessageResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MessageService_BidirectionalStreamMessagesClient = grpc.BidiStreamingClient[MessageRequest, MessageResponse]

// MessageServiceServer is the server API for MessageService service.
// All implementations must embed UnimplementedMessageServiceServer
// for forward compatibility.
type MessageServiceServer interface {
	BidirectionalStreamMessages(grpc.BidiStreamingServer[MessageRequest, MessageResponse]) error
	mustEmbedUnimplementedMessageServiceServer()
}

// UnimplementedMessageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMessageServiceServer struct{}

func (UnimplementedMessageServiceServer) BidirectionalStreamMessages(grpc.BidiStreamingServer[MessageRequest, MessageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamMessages not implemented")
}
func (UnimplementedMessageServiceServer) mustEmbedUnimplementedMessageServiceServer() {}
func (UnimplementedMessageServiceServer) testEmbeddedByValue()                        {}

// UnsafeMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessageServiceServer will
// result in compilation errors.
type UnsafeMessageServiceServer interface {
	mustEmbedUnimplementedMessageServiceServer()
}

func RegisterMessageServiceServer(s grpc.ServiceRegistrar, srv MessageServiceServer) {
	// If the following call pancis, it indicates UnimplementedMessageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MessageService_ServiceDesc, srv)
}

func _MessageService_BidirectionalStreamMessages_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).BidirectionalStreamMessages(&grpc.GenericServerStream[MessageRequest, MessageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MessageService_BidirectionalStreamMessagesServer = grpc.BidiStreamingServer[MessageRequest, MessageResponse]

// MessageService_ServiceDesc is the grpc.ServiceDesc for MessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BidirectionalStreamMessages",
			Handler:       _MessageService_BidirectionalStreamMessages_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}
