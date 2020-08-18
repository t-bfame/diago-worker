// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package internal

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WorkerClient is the client API for Worker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WorkerClient interface {
	Coordinate(ctx context.Context, opts ...grpc.CallOption) (Worker_CoordinateClient, error)
}

type workerClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkerClient(cc grpc.ClientConnInterface) WorkerClient {
	return &workerClient{cc}
}

func (c *workerClient) Coordinate(ctx context.Context, opts ...grpc.CallOption) (Worker_CoordinateClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Worker_serviceDesc.Streams[0], "/Worker/Coordinate", opts...)
	if err != nil {
		return nil, err
	}
	x := &workerCoordinateClient{stream}
	return x, nil
}

type Worker_CoordinateClient interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ClientStream
}

type workerCoordinateClient struct {
	grpc.ClientStream
}

func (x *workerCoordinateClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *workerCoordinateClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WorkerServer is the server API for Worker service.
// All implementations must embed UnimplementedWorkerServer
// for forward compatibility
type WorkerServer interface {
	Coordinate(Worker_CoordinateServer) error
	mustEmbedUnimplementedWorkerServer()
}

// UnimplementedWorkerServer must be embedded to have forward compatible implementations.
type UnimplementedWorkerServer struct {
}

func (*UnimplementedWorkerServer) Coordinate(Worker_CoordinateServer) error {
	return status.Errorf(codes.Unimplemented, "method Coordinate not implemented")
}
func (*UnimplementedWorkerServer) mustEmbedUnimplementedWorkerServer() {}

func RegisterWorkerServer(s *grpc.Server, srv WorkerServer) {
	s.RegisterService(&_Worker_serviceDesc, srv)
}

func _Worker_Coordinate_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(WorkerServer).Coordinate(&workerCoordinateServer{stream})
}

type Worker_CoordinateServer interface {
	Send(*Message) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type workerCoordinateServer struct {
	grpc.ServerStream
}

func (x *workerCoordinateServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func (x *workerCoordinateServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Worker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Worker",
	HandlerType: (*WorkerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Coordinate",
			Handler:       _Worker_Coordinate_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "internal/worker.proto",
}
