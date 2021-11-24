// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GroupsClient is the client API for Groups service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupsClient interface {
	GetAll(ctx context.Context, in *Groups, opts ...grpc.CallOption) (*Groups, error)
	Get(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Group, error)
	GetByName(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Group, error)
	GetGraph(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Groups, error)
	Add(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Group, error)
	AddSubGroup(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Group, error)
	Update(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Group, error)
}

type groupsClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupsClient(cc grpc.ClientConnInterface) GroupsClient {
	return &groupsClient{cc}
}

func (c *groupsClient) GetAll(ctx context.Context, in *Groups, opts ...grpc.CallOption) (*Groups, error) {
	out := new(Groups)
	err := c.cc.Invoke(ctx, "/grpc.groups/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) Get(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/grpc.groups/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) GetByName(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/grpc.groups/GetByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) GetGraph(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Groups, error) {
	out := new(Groups)
	err := c.cc.Invoke(ctx, "/grpc.groups/GetGraph", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) Add(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/grpc.groups/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) AddSubGroup(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/grpc.groups/AddSubGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupsClient) Update(ctx context.Context, in *Group, opts ...grpc.CallOption) (*Group, error) {
	out := new(Group)
	err := c.cc.Invoke(ctx, "/grpc.groups/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupsServer is the server API for Groups service.
// All implementations must embed UnimplementedGroupsServer
// for forward compatibility
type GroupsServer interface {
	GetAll(context.Context, *Groups) (*Groups, error)
	Get(context.Context, *Group) (*Group, error)
	GetByName(context.Context, *Group) (*Group, error)
	GetGraph(context.Context, *emptypb.Empty) (*Groups, error)
	Add(context.Context, *Group) (*Group, error)
	AddSubGroup(context.Context, *Group) (*Group, error)
	Update(context.Context, *Group) (*Group, error)
	mustEmbedUnimplementedGroupsServer()
}

// UnimplementedGroupsServer must be embedded to have forward compatible implementations.
type UnimplementedGroupsServer struct {
}

func (UnimplementedGroupsServer) GetAll(context.Context, *Groups) (*Groups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedGroupsServer) Get(context.Context, *Group) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGroupsServer) GetByName(context.Context, *Group) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedGroupsServer) GetGraph(context.Context, *emptypb.Empty) (*Groups, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGraph not implemented")
}
func (UnimplementedGroupsServer) Add(context.Context, *Group) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedGroupsServer) AddSubGroup(context.Context, *Group) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSubGroup not implemented")
}
func (UnimplementedGroupsServer) Update(context.Context, *Group) (*Group, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGroupsServer) mustEmbedUnimplementedGroupsServer() {}

// UnsafeGroupsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupsServer will
// result in compilation errors.
type UnsafeGroupsServer interface {
	mustEmbedUnimplementedGroupsServer()
}

func RegisterGroupsServer(s grpc.ServiceRegistrar, srv GroupsServer) {
	s.RegisterService(&Groups_ServiceDesc, srv)
}

func _Groups_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Groups)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.groups/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).GetAll(ctx, req.(*Groups))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.groups/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).Get(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.groups/GetByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).GetByName(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_GetGraph_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).GetGraph(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.groups/GetGraph",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).GetGraph(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.groups/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).Add(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_AddSubGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).AddSubGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.groups/AddSubGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).AddSubGroup(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _Groups_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.groups/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupsServer).Update(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

// Groups_ServiceDesc is the grpc.ServiceDesc for Groups service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Groups_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.groups",
	HandlerType: (*GroupsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _Groups_GetAll_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Groups_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _Groups_GetByName_Handler,
		},
		{
			MethodName: "GetGraph",
			Handler:    _Groups_GetGraph_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Groups_Add_Handler,
		},
		{
			MethodName: "AddSubGroup",
			Handler:    _Groups_AddSubGroup_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Groups_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "groups.proto",
}
