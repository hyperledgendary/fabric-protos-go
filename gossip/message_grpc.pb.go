// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gossip

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GossipClient is the client API for Gossip service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GossipClient interface {
	// GossipStream is the gRPC stream used for sending and receiving messages
	GossipStream(ctx context.Context, opts ...grpc.CallOption) (Gossip_GossipStreamClient, error)
	// Ping is used to probe a remote peer's aliveness
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
}

type gossipClient struct {
	cc grpc.ClientConnInterface
}

func NewGossipClient(cc grpc.ClientConnInterface) GossipClient {
	return &gossipClient{cc}
}

func (c *gossipClient) GossipStream(ctx context.Context, opts ...grpc.CallOption) (Gossip_GossipStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Gossip_ServiceDesc.Streams[0], "/gossip.Gossip/GossipStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gossipGossipStreamClient{stream}
	return x, nil
}

type Gossip_GossipStreamClient interface {
	Send(*Envelope) error
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type gossipGossipStreamClient struct {
	grpc.ClientStream
}

func (x *gossipGossipStreamClient) Send(m *Envelope) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gossipGossipStreamClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gossipClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/gossip.Gossip/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GossipServer is the server API for Gossip service.
// All implementations must embed UnimplementedGossipServer
// for forward compatibility
type GossipServer interface {
	// GossipStream is the gRPC stream used for sending and receiving messages
	GossipStream(Gossip_GossipStreamServer) error
	// Ping is used to probe a remote peer's aliveness
	Ping(context.Context, *Empty) (*Empty, error)
	mustEmbedUnimplementedGossipServer()
}

// UnimplementedGossipServer must be embedded to have forward compatible implementations.
type UnimplementedGossipServer struct {
}

func (UnimplementedGossipServer) GossipStream(Gossip_GossipStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GossipStream not implemented")
}
func (UnimplementedGossipServer) Ping(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedGossipServer) mustEmbedUnimplementedGossipServer() {}

// UnsafeGossipServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GossipServer will
// result in compilation errors.
type UnsafeGossipServer interface {
	mustEmbedUnimplementedGossipServer()
}

func RegisterGossipServer(s grpc.ServiceRegistrar, srv GossipServer) {
	s.RegisterService(&Gossip_ServiceDesc, srv)
}

func _Gossip_GossipStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GossipServer).GossipStream(&gossipGossipStreamServer{stream})
}

type Gossip_GossipStreamServer interface {
	Send(*Envelope) error
	Recv() (*Envelope, error)
	grpc.ServerStream
}

type gossipGossipStreamServer struct {
	grpc.ServerStream
}

func (x *gossipGossipStreamServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gossipGossipStreamServer) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Gossip_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GossipServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gossip.Gossip/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GossipServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Gossip_ServiceDesc is the grpc.ServiceDesc for Gossip service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gossip_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gossip.Gossip",
	HandlerType: (*GossipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Gossip_Ping_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GossipStream",
			Handler:       _Gossip_GossipStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "gossip/message.proto",
}
