// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: management/resource/v1alpha1/listwatch.proto

package v1alpha1

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

// ResourceListWatchManagementClient is the client API for ResourceListWatchManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceListWatchManagementClient interface {
	ListWatch(ctx context.Context, in *ListWatchRequest, opts ...grpc.CallOption) (ResourceListWatchManagement_ListWatchClient, error)
}

type resourceListWatchManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceListWatchManagementClient(cc grpc.ClientConnInterface) ResourceListWatchManagementClient {
	return &resourceListWatchManagementClient{cc}
}

func (c *resourceListWatchManagementClient) ListWatch(ctx context.Context, in *ListWatchRequest, opts ...grpc.CallOption) (ResourceListWatchManagement_ListWatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &ResourceListWatchManagement_ServiceDesc.Streams[0], "/vink.kubevm.io.apis.management.resource.v1alpha1.ResourceListWatchManagement/ListWatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &resourceListWatchManagementListWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ResourceListWatchManagement_ListWatchClient interface {
	Recv() (*ListWatchResponse, error)
	grpc.ClientStream
}

type resourceListWatchManagementListWatchClient struct {
	grpc.ClientStream
}

func (x *resourceListWatchManagementListWatchClient) Recv() (*ListWatchResponse, error) {
	m := new(ListWatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ResourceListWatchManagementServer is the server API for ResourceListWatchManagement service.
// All implementations must embed UnimplementedResourceListWatchManagementServer
// for forward compatibility
type ResourceListWatchManagementServer interface {
	ListWatch(*ListWatchRequest, ResourceListWatchManagement_ListWatchServer) error
	mustEmbedUnimplementedResourceListWatchManagementServer()
}

// UnimplementedResourceListWatchManagementServer must be embedded to have forward compatible implementations.
type UnimplementedResourceListWatchManagementServer struct {
}

func (UnimplementedResourceListWatchManagementServer) ListWatch(*ListWatchRequest, ResourceListWatchManagement_ListWatchServer) error {
	return status.Errorf(codes.Unimplemented, "method ListWatch not implemented")
}
func (UnimplementedResourceListWatchManagementServer) mustEmbedUnimplementedResourceListWatchManagementServer() {
}

// UnsafeResourceListWatchManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceListWatchManagementServer will
// result in compilation errors.
type UnsafeResourceListWatchManagementServer interface {
	mustEmbedUnimplementedResourceListWatchManagementServer()
}

func RegisterResourceListWatchManagementServer(s grpc.ServiceRegistrar, srv ResourceListWatchManagementServer) {
	s.RegisterService(&ResourceListWatchManagement_ServiceDesc, srv)
}

func _ResourceListWatchManagement_ListWatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListWatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourceListWatchManagementServer).ListWatch(m, &resourceListWatchManagementListWatchServer{stream})
}

type ResourceListWatchManagement_ListWatchServer interface {
	Send(*ListWatchResponse) error
	grpc.ServerStream
}

type resourceListWatchManagementListWatchServer struct {
	grpc.ServerStream
}

func (x *resourceListWatchManagementListWatchServer) Send(m *ListWatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ResourceListWatchManagement_ServiceDesc is the grpc.ServiceDesc for ResourceListWatchManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceListWatchManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "vink.kubevm.io.apis.management.resource.v1alpha1.ResourceListWatchManagement",
	HandlerType: (*ResourceListWatchManagementServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListWatch",
			Handler:       _ResourceListWatchManagement_ListWatch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "management/resource/v1alpha1/listwatch.proto",
}
