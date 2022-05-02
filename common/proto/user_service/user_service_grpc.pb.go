// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: common/proto/user_service/user_service.proto

package users

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Insert(ctx context.Context, in *InsertUserRequest, opts ...grpc.CallOption) (*InsertUserResponse, error)
	EditUser(ctx context.Context, in *InsertUserRequest, opts ...grpc.CallOption) (*EditUserResponse, error)
	AddEducation(ctx context.Context, in *AddEducationRequest, opts ...grpc.CallOption) (*AddEducationResponse, error)
	DeleteEducation(ctx context.Context, in *DeleteEducationRequest, opts ...grpc.CallOption) (*DeleteEducationResponse, error)
	AddExperience(ctx context.Context, in *AddExperienceRequest, opts ...grpc.CallOption) (*AddExperienceResponse, error)
	DeleteExperience(ctx context.Context, in *DeleteExperienceRequest, opts ...grpc.CallOption) (*DeleteExperienceResponse, error)
	AddInterest(ctx context.Context, in *AddInterestRequest, opts ...grpc.CallOption) (*AddInterestResponse, error)
	DeleteInterest(ctx context.Context, in *DeleteInterestRequest, opts ...grpc.CallOption) (*DeleteInterestResponse, error)
	AddSkill(ctx context.Context, in *AddSkillRequest, opts ...grpc.CallOption) (*AddSkillResponse, error)
	DeleteSkill(ctx context.Context, in *DeleteSkillRequest, opts ...grpc.CallOption) (*DeleteSkillResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/GetAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Insert(ctx context.Context, in *InsertUserRequest, opts ...grpc.CallOption) (*InsertUserResponse, error) {
	out := new(InsertUserResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/Insert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) EditUser(ctx context.Context, in *InsertUserRequest, opts ...grpc.CallOption) (*EditUserResponse, error) {
	out := new(EditUserResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/EditUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddEducation(ctx context.Context, in *AddEducationRequest, opts ...grpc.CallOption) (*AddEducationResponse, error) {
	out := new(AddEducationResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/AddEducation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteEducation(ctx context.Context, in *DeleteEducationRequest, opts ...grpc.CallOption) (*DeleteEducationResponse, error) {
	out := new(DeleteEducationResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/DeleteEducation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddExperience(ctx context.Context, in *AddExperienceRequest, opts ...grpc.CallOption) (*AddExperienceResponse, error) {
	out := new(AddExperienceResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/AddExperience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteExperience(ctx context.Context, in *DeleteExperienceRequest, opts ...grpc.CallOption) (*DeleteExperienceResponse, error) {
	out := new(DeleteExperienceResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/DeleteExperience", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddInterest(ctx context.Context, in *AddInterestRequest, opts ...grpc.CallOption) (*AddInterestResponse, error) {
	out := new(AddInterestResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/AddInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteInterest(ctx context.Context, in *DeleteInterestRequest, opts ...grpc.CallOption) (*DeleteInterestResponse, error) {
	out := new(DeleteInterestResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/DeleteInterest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) AddSkill(ctx context.Context, in *AddSkillRequest, opts ...grpc.CallOption) (*AddSkillResponse, error) {
	out := new(AddSkillResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/AddSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteSkill(ctx context.Context, in *DeleteSkillRequest, opts ...grpc.CallOption) (*DeleteSkillResponse, error) {
	out := new(DeleteSkillResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/DeleteSkill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	Insert(context.Context, *InsertUserRequest) (*InsertUserResponse, error)
	EditUser(context.Context, *InsertUserRequest) (*EditUserResponse, error)
	AddEducation(context.Context, *AddEducationRequest) (*AddEducationResponse, error)
	DeleteEducation(context.Context, *DeleteEducationRequest) (*DeleteEducationResponse, error)
	AddExperience(context.Context, *AddExperienceRequest) (*AddExperienceResponse, error)
	DeleteExperience(context.Context, *DeleteExperienceRequest) (*DeleteExperienceResponse, error)
	AddInterest(context.Context, *AddInterestRequest) (*AddInterestResponse, error)
	DeleteInterest(context.Context, *DeleteInterestRequest) (*DeleteInterestResponse, error)
	AddSkill(context.Context, *AddSkillRequest) (*AddSkillResponse, error)
	DeleteSkill(context.Context, *DeleteSkillRequest) (*DeleteSkillResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUserServiceServer) Insert(context.Context, *InsertUserRequest) (*InsertUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedUserServiceServer) EditUser(context.Context, *InsertUserRequest) (*EditUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUser not implemented")
}
func (UnimplementedUserServiceServer) AddEducation(context.Context, *AddEducationRequest) (*AddEducationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEducation not implemented")
}
func (UnimplementedUserServiceServer) DeleteEducation(context.Context, *DeleteEducationRequest) (*DeleteEducationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEducation not implemented")
}
func (UnimplementedUserServiceServer) AddExperience(context.Context, *AddExperienceRequest) (*AddExperienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExperience not implemented")
}
func (UnimplementedUserServiceServer) DeleteExperience(context.Context, *DeleteExperienceRequest) (*DeleteExperienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExperience not implemented")
}
func (UnimplementedUserServiceServer) AddInterest(context.Context, *AddInterestRequest) (*AddInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInterest not implemented")
}
func (UnimplementedUserServiceServer) DeleteInterest(context.Context, *DeleteInterestRequest) (*DeleteInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInterest not implemented")
}
func (UnimplementedUserServiceServer) AddSkill(context.Context, *AddSkillRequest) (*AddSkillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSkill not implemented")
}
func (UnimplementedUserServiceServer) DeleteSkill(context.Context, *DeleteSkillRequest) (*DeleteSkillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSkill not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Insert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Insert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/Insert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Insert(ctx, req.(*InsertUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_EditUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).EditUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/EditUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).EditUser(ctx, req.(*InsertUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddEducation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEducationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddEducation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/AddEducation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddEducation(ctx, req.(*AddEducationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteEducation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEducationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteEducation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/DeleteEducation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteEducation(ctx, req.(*DeleteEducationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddExperience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddExperienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddExperience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/AddExperience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddExperience(ctx, req.(*AddExperienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteExperience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteExperienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteExperience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/DeleteExperience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteExperience(ctx, req.(*DeleteExperienceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/AddInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddInterest(ctx, req.(*AddInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteInterest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInterestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteInterest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/DeleteInterest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteInterest(ctx, req.(*DeleteInterestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_AddSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).AddSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/AddSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).AddSkill(ctx, req.(*AddSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteSkill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSkillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteSkill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/DeleteSkill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteSkill(ctx, req.(*DeleteSkillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "users.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _UserService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "Insert",
			Handler:    _UserService_Insert_Handler,
		},
		{
			MethodName: "EditUser",
			Handler:    _UserService_EditUser_Handler,
		},
		{
			MethodName: "AddEducation",
			Handler:    _UserService_AddEducation_Handler,
		},
		{
			MethodName: "DeleteEducation",
			Handler:    _UserService_DeleteEducation_Handler,
		},
		{
			MethodName: "AddExperience",
			Handler:    _UserService_AddExperience_Handler,
		},
		{
			MethodName: "DeleteExperience",
			Handler:    _UserService_DeleteExperience_Handler,
		},
		{
			MethodName: "AddInterest",
			Handler:    _UserService_AddInterest_Handler,
		},
		{
			MethodName: "DeleteInterest",
			Handler:    _UserService_DeleteInterest_Handler,
		},
		{
			MethodName: "AddSkill",
			Handler:    _UserService_AddSkill_Handler,
		},
		{
			MethodName: "DeleteSkill",
			Handler:    _UserService_DeleteSkill_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/proto/user_service/user_service.proto",
}
