// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: subscribes.proto

package subs

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

const (
	Subscribe_AddGroup_FullMethodName                = "/subs.Subscribe/AddGroup"
	Subscribe_DeleteGroup_FullMethodName             = "/subs.Subscribe/DeleteGroup"
	Subscribe_GetGroupList_FullMethodName            = "/subs.Subscribe/GetGroupList"
	Subscribe_SetGroups_FullMethodName               = "/subs.Subscribe/SetGroups"
	Subscribe_UpdateGroupName_FullMethodName         = "/subs.Subscribe/UpdateGroupName"
	Subscribe_Subscribe_FullMethodName               = "/subs.Subscribe/Subscribe"
	Subscribe_UnSubscribe_FullMethodName             = "/subs.Subscribe/UnSubscribe"
	Subscribe_VerifySubscribe_FullMethodName         = "/subs.Subscribe/VerifySubscribe"
	Subscribe_QueryAttentionByGroup_FullMethodName   = "/subs.Subscribe/QueryAttentionByGroup"
	Subscribe_QueryFollowList_FullMethodName         = "/subs.Subscribe/QueryFollowList"
	Subscribe_RemoveFromAttentionList_FullMethodName = "/subs.Subscribe/RemoveFromAttentionList"
)

// SubscribeClient is the client API for Subscribe service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscribeClient interface {
	// 添加关注分组
	AddGroup(ctx context.Context, in *AddGroupReq, opts ...grpc.CallOption) (*AddGroupResp, error)
	// 删除关注分组
	DeleteGroup(ctx context.Context, in *DeleteGroupReq, opts ...grpc.CallOption) (*Empty, error)
	// 获取用户的所有关注分组
	GetGroupList(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*GroupList, error)
	// 为一个订阅设置分组
	SetGroups(ctx context.Context, in *SetGroupsReq, opts ...grpc.CallOption) (*Empty, error)
	// 修改分组名称
	UpdateGroupName(ctx context.Context, in *UpdateGroupNameReq, opts ...grpc.CallOption) (*Empty, error)
	// 订阅
	Subscribe(ctx context.Context, in *SubscribeReq, opts ...grpc.CallOption) (*Empty, error)
	// 取消订阅
	UnSubscribe(ctx context.Context, in *UnSubscribeReq, opts ...grpc.CallOption) (*Empty, error)
	// 查询一个用户是否订阅另一个用户
	VerifySubscribe(ctx context.Context, in *VerifySubscriptionReq, opts ...grpc.CallOption) (*Empty, error)
	// 查询分组下的订阅列表 分页
	QueryAttentionByGroup(ctx context.Context, in *QueryAttentionByGroupReq, opts ...grpc.CallOption) (*QueryAttentionResp, error)
	// 查询粉丝列表 分页
	QueryFollowList(ctx context.Context, in *QueryFollowerList, opts ...grpc.CallOption) (*QueryFollowResp, error)
	// 从关注列表移除-软删除 设置status
	RemoveFromAttentionList(ctx context.Context, in *RemoveAttentionReq, opts ...grpc.CallOption) (*Empty, error)
}

type subscribeClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscribeClient(cc grpc.ClientConnInterface) SubscribeClient {
	return &subscribeClient{cc}
}

func (c *subscribeClient) AddGroup(ctx context.Context, in *AddGroupReq, opts ...grpc.CallOption) (*AddGroupResp, error) {
	out := new(AddGroupResp)
	err := c.cc.Invoke(ctx, Subscribe_AddGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) DeleteGroup(ctx context.Context, in *DeleteGroupReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Subscribe_DeleteGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) GetGroupList(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*GroupList, error) {
	out := new(GroupList)
	err := c.cc.Invoke(ctx, Subscribe_GetGroupList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) SetGroups(ctx context.Context, in *SetGroupsReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Subscribe_SetGroups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) UpdateGroupName(ctx context.Context, in *UpdateGroupNameReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Subscribe_UpdateGroupName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) Subscribe(ctx context.Context, in *SubscribeReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Subscribe_Subscribe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) UnSubscribe(ctx context.Context, in *UnSubscribeReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Subscribe_UnSubscribe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) VerifySubscribe(ctx context.Context, in *VerifySubscriptionReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Subscribe_VerifySubscribe_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) QueryAttentionByGroup(ctx context.Context, in *QueryAttentionByGroupReq, opts ...grpc.CallOption) (*QueryAttentionResp, error) {
	out := new(QueryAttentionResp)
	err := c.cc.Invoke(ctx, Subscribe_QueryAttentionByGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) QueryFollowList(ctx context.Context, in *QueryFollowerList, opts ...grpc.CallOption) (*QueryFollowResp, error) {
	out := new(QueryFollowResp)
	err := c.cc.Invoke(ctx, Subscribe_QueryFollowList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscribeClient) RemoveFromAttentionList(ctx context.Context, in *RemoveAttentionReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, Subscribe_RemoveFromAttentionList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscribeServer is the server API for Subscribe service.
// All implementations must embed UnimplementedSubscribeServer
// for forward compatibility
type SubscribeServer interface {
	// 添加关注分组
	AddGroup(context.Context, *AddGroupReq) (*AddGroupResp, error)
	// 删除关注分组
	DeleteGroup(context.Context, *DeleteGroupReq) (*Empty, error)
	// 获取用户的所有关注分组
	GetGroupList(context.Context, *UserId) (*GroupList, error)
	// 为一个订阅设置分组
	SetGroups(context.Context, *SetGroupsReq) (*Empty, error)
	// 修改分组名称
	UpdateGroupName(context.Context, *UpdateGroupNameReq) (*Empty, error)
	// 订阅
	Subscribe(context.Context, *SubscribeReq) (*Empty, error)
	// 取消订阅
	UnSubscribe(context.Context, *UnSubscribeReq) (*Empty, error)
	// 查询一个用户是否订阅另一个用户
	VerifySubscribe(context.Context, *VerifySubscriptionReq) (*Empty, error)
	// 查询分组下的订阅列表 分页
	QueryAttentionByGroup(context.Context, *QueryAttentionByGroupReq) (*QueryAttentionResp, error)
	// 查询粉丝列表 分页
	QueryFollowList(context.Context, *QueryFollowerList) (*QueryFollowResp, error)
	// 从关注列表移除-软删除 设置status
	RemoveFromAttentionList(context.Context, *RemoveAttentionReq) (*Empty, error)
	mustEmbedUnimplementedSubscribeServer()
}

// UnimplementedSubscribeServer must be embedded to have forward compatible implementations.
type UnimplementedSubscribeServer struct {
}

func (UnimplementedSubscribeServer) AddGroup(context.Context, *AddGroupReq) (*AddGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroup not implemented")
}
func (UnimplementedSubscribeServer) DeleteGroup(context.Context, *DeleteGroupReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroup not implemented")
}
func (UnimplementedSubscribeServer) GetGroupList(context.Context, *UserId) (*GroupList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupList not implemented")
}
func (UnimplementedSubscribeServer) SetGroups(context.Context, *SetGroupsReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGroups not implemented")
}
func (UnimplementedSubscribeServer) UpdateGroupName(context.Context, *UpdateGroupNameReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupName not implemented")
}
func (UnimplementedSubscribeServer) Subscribe(context.Context, *SubscribeReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSubscribeServer) UnSubscribe(context.Context, *UnSubscribeReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnSubscribe not implemented")
}
func (UnimplementedSubscribeServer) VerifySubscribe(context.Context, *VerifySubscriptionReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifySubscribe not implemented")
}
func (UnimplementedSubscribeServer) QueryAttentionByGroup(context.Context, *QueryAttentionByGroupReq) (*QueryAttentionResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryAttentionByGroup not implemented")
}
func (UnimplementedSubscribeServer) QueryFollowList(context.Context, *QueryFollowerList) (*QueryFollowResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryFollowList not implemented")
}
func (UnimplementedSubscribeServer) RemoveFromAttentionList(context.Context, *RemoveAttentionReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveFromAttentionList not implemented")
}
func (UnimplementedSubscribeServer) mustEmbedUnimplementedSubscribeServer() {}

// UnsafeSubscribeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscribeServer will
// result in compilation errors.
type UnsafeSubscribeServer interface {
	mustEmbedUnimplementedSubscribeServer()
}

func RegisterSubscribeServer(s grpc.ServiceRegistrar, srv SubscribeServer) {
	s.RegisterService(&Subscribe_ServiceDesc, srv)
}

func _Subscribe_AddGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).AddGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Subscribe_AddGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).AddGroup(ctx, req.(*AddGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_DeleteGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).DeleteGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Subscribe_DeleteGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).DeleteGroup(ctx, req.(*DeleteGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_GetGroupList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).GetGroupList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Subscribe_GetGroupList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).GetGroupList(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_SetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).SetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Subscribe_SetGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).SetGroups(ctx, req.(*SetGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_UpdateGroupName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupNameReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).UpdateGroupName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Subscribe_UpdateGroupName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).UpdateGroupName(ctx, req.(*UpdateGroupNameReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Subscribe_Subscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).Subscribe(ctx, req.(*SubscribeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_UnSubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnSubscribeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).UnSubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Subscribe_UnSubscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).UnSubscribe(ctx, req.(*UnSubscribeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_VerifySubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifySubscriptionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).VerifySubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Subscribe_VerifySubscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).VerifySubscribe(ctx, req.(*VerifySubscriptionReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_QueryAttentionByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAttentionByGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).QueryAttentionByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Subscribe_QueryAttentionByGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).QueryAttentionByGroup(ctx, req.(*QueryAttentionByGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_QueryFollowList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryFollowerList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).QueryFollowList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Subscribe_QueryFollowList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).QueryFollowList(ctx, req.(*QueryFollowerList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscribe_RemoveFromAttentionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAttentionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscribeServer).RemoveFromAttentionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Subscribe_RemoveFromAttentionList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscribeServer).RemoveFromAttentionList(ctx, req.(*RemoveAttentionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Subscribe_ServiceDesc is the grpc.ServiceDesc for Subscribe service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Subscribe_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subs.Subscribe",
	HandlerType: (*SubscribeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddGroup",
			Handler:    _Subscribe_AddGroup_Handler,
		},
		{
			MethodName: "DeleteGroup",
			Handler:    _Subscribe_DeleteGroup_Handler,
		},
		{
			MethodName: "GetGroupList",
			Handler:    _Subscribe_GetGroupList_Handler,
		},
		{
			MethodName: "SetGroups",
			Handler:    _Subscribe_SetGroups_Handler,
		},
		{
			MethodName: "UpdateGroupName",
			Handler:    _Subscribe_UpdateGroupName_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _Subscribe_Subscribe_Handler,
		},
		{
			MethodName: "UnSubscribe",
			Handler:    _Subscribe_UnSubscribe_Handler,
		},
		{
			MethodName: "VerifySubscribe",
			Handler:    _Subscribe_VerifySubscribe_Handler,
		},
		{
			MethodName: "QueryAttentionByGroup",
			Handler:    _Subscribe_QueryAttentionByGroup_Handler,
		},
		{
			MethodName: "QueryFollowList",
			Handler:    _Subscribe_QueryFollowList_Handler,
		},
		{
			MethodName: "RemoveFromAttentionList",
			Handler:    _Subscribe_RemoveFromAttentionList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscribes.proto",
}
