// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: replicationtypes/bertyreplication.proto

package replicationtypes

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

// ReplicationServiceClient is the client API for ReplicationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReplicationServiceClient interface {
	// ReplicateGroup
	ReplicateGroup(ctx context.Context, in *ReplicationServiceReplicateGroup_Request, opts ...grpc.CallOption) (*ReplicationServiceReplicateGroup_Reply, error)
	ReplicateGlobalStats(ctx context.Context, in *ReplicateGlobalStats_Request, opts ...grpc.CallOption) (*ReplicateGlobalStats_Reply, error)
	ReplicateGroupStats(ctx context.Context, in *ReplicateGroupStats_Request, opts ...grpc.CallOption) (*ReplicateGroupStats_Reply, error)
}

type replicationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicationServiceClient(cc grpc.ClientConnInterface) ReplicationServiceClient {
	return &replicationServiceClient{cc}
}

func (c *replicationServiceClient) ReplicateGroup(ctx context.Context, in *ReplicationServiceReplicateGroup_Request, opts ...grpc.CallOption) (*ReplicationServiceReplicateGroup_Reply, error) {
	out := new(ReplicationServiceReplicateGroup_Reply)
	err := c.cc.Invoke(ctx, "/berty.replication.v1.ReplicationService/ReplicateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationServiceClient) ReplicateGlobalStats(ctx context.Context, in *ReplicateGlobalStats_Request, opts ...grpc.CallOption) (*ReplicateGlobalStats_Reply, error) {
	out := new(ReplicateGlobalStats_Reply)
	err := c.cc.Invoke(ctx, "/berty.replication.v1.ReplicationService/ReplicateGlobalStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *replicationServiceClient) ReplicateGroupStats(ctx context.Context, in *ReplicateGroupStats_Request, opts ...grpc.CallOption) (*ReplicateGroupStats_Reply, error) {
	out := new(ReplicateGroupStats_Reply)
	err := c.cc.Invoke(ctx, "/berty.replication.v1.ReplicationService/ReplicateGroupStats", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicationServiceServer is the server API for ReplicationService service.
// All implementations must embed UnimplementedReplicationServiceServer
// for forward compatibility
type ReplicationServiceServer interface {
	// ReplicateGroup
	ReplicateGroup(context.Context, *ReplicationServiceReplicateGroup_Request) (*ReplicationServiceReplicateGroup_Reply, error)
	ReplicateGlobalStats(context.Context, *ReplicateGlobalStats_Request) (*ReplicateGlobalStats_Reply, error)
	ReplicateGroupStats(context.Context, *ReplicateGroupStats_Request) (*ReplicateGroupStats_Reply, error)
	mustEmbedUnimplementedReplicationServiceServer()
}

// UnimplementedReplicationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReplicationServiceServer struct {
}

func (UnimplementedReplicationServiceServer) ReplicateGroup(context.Context, *ReplicationServiceReplicateGroup_Request) (*ReplicationServiceReplicateGroup_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplicateGroup not implemented")
}
func (UnimplementedReplicationServiceServer) ReplicateGlobalStats(context.Context, *ReplicateGlobalStats_Request) (*ReplicateGlobalStats_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplicateGlobalStats not implemented")
}
func (UnimplementedReplicationServiceServer) ReplicateGroupStats(context.Context, *ReplicateGroupStats_Request) (*ReplicateGroupStats_Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplicateGroupStats not implemented")
}
func (UnimplementedReplicationServiceServer) mustEmbedUnimplementedReplicationServiceServer() {}

// UnsafeReplicationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReplicationServiceServer will
// result in compilation errors.
type UnsafeReplicationServiceServer interface {
	mustEmbedUnimplementedReplicationServiceServer()
}

func RegisterReplicationServiceServer(s grpc.ServiceRegistrar, srv ReplicationServiceServer) {
	s.RegisterService(&ReplicationService_ServiceDesc, srv)
}

func _ReplicationService_ReplicateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplicationServiceReplicateGroup_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).ReplicateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.replication.v1.ReplicationService/ReplicateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).ReplicateGroup(ctx, req.(*ReplicationServiceReplicateGroup_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicationService_ReplicateGlobalStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplicateGlobalStats_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).ReplicateGlobalStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.replication.v1.ReplicationService/ReplicateGlobalStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).ReplicateGlobalStats(ctx, req.(*ReplicateGlobalStats_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReplicationService_ReplicateGroupStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplicateGroupStats_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicationServiceServer).ReplicateGroupStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/berty.replication.v1.ReplicationService/ReplicateGroupStats",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicationServiceServer).ReplicateGroupStats(ctx, req.(*ReplicateGroupStats_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// ReplicationService_ServiceDesc is the grpc.ServiceDesc for ReplicationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReplicationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "berty.replication.v1.ReplicationService",
	HandlerType: (*ReplicationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReplicateGroup",
			Handler:    _ReplicationService_ReplicateGroup_Handler,
		},
		{
			MethodName: "ReplicateGlobalStats",
			Handler:    _ReplicationService_ReplicateGlobalStats_Handler,
		},
		{
			MethodName: "ReplicateGroupStats",
			Handler:    _ReplicationService_ReplicateGroupStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "replicationtypes/bertyreplication.proto",
}
