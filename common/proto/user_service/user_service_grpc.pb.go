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
	GetByUsername(ctx context.Context, in *GetByUsernameRequest, opts ...grpc.CallOption) (*GetByUsernameResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetPublicUsers(ctx context.Context, in *GetPublicUsersRequest, opts ...grpc.CallOption) (*GetPublicUsersResponse, error)
	Insert(ctx context.Context, in *InsertUserRequest, opts ...grpc.CallOption) (*InsertUserResponse, error)
	EditUser(ctx context.Context, in *InsertUserRequest, opts ...grpc.CallOption) (*EditUserResponse, error)
	GetEducation(ctx context.Context, in *GetEducationRequest, opts ...grpc.CallOption) (*GetEducationResponse, error)
	AddEducation(ctx context.Context, in *AddEducationRequest, opts ...grpc.CallOption) (*AddEducationResponse, error)
	DeleteEducation(ctx context.Context, in *DeleteEducationRequest, opts ...grpc.CallOption) (*DeleteEducationResponse, error)
	GetExperience(ctx context.Context, in *GetExperienceRequest, opts ...grpc.CallOption) (*GetExperienceResponse, error)
	AddExperience(ctx context.Context, in *AddExperienceRequest, opts ...grpc.CallOption) (*AddExperienceResponse, error)
	DeleteExperience(ctx context.Context, in *DeleteExperienceRequest, opts ...grpc.CallOption) (*DeleteExperienceResponse, error)
	GetInterests(ctx context.Context, in *GetInterestsRequest, opts ...grpc.CallOption) (*GetInterestsResponse, error)
	AddInterest(ctx context.Context, in *AddInterestRequest, opts ...grpc.CallOption) (*AddInterestResponse, error)
	DeleteInterest(ctx context.Context, in *DeleteInterestRequest, opts ...grpc.CallOption) (*DeleteInterestResponse, error)
	GetSkills(ctx context.Context, in *GetSkillsRequest, opts ...grpc.CallOption) (*GetSkillsResponse, error)
	AddSkill(ctx context.Context, in *AddSkillRequest, opts ...grpc.CallOption) (*AddSkillResponse, error)
	DeleteSkill(ctx context.Context, in *DeleteSkillRequest, opts ...grpc.CallOption) (*DeleteSkillResponse, error)
	SearchProfiles(ctx context.Context, in *SearchProfilesRequest, opts ...grpc.CallOption) (*SearchProfilesResponse, error)
	SetPrivacy(ctx context.Context, in *SetPrivacyRequest, opts ...grpc.CallOption) (*SetPrivacyResponse, error)
	SetApiKey(ctx context.Context, in *SetApiKeyRequest, opts ...grpc.CallOption) (*SetApiKeyResponse, error)
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

func (c *userServiceClient) GetByUsername(ctx context.Context, in *GetByUsernameRequest, opts ...grpc.CallOption) (*GetByUsernameResponse, error) {
	out := new(GetByUsernameResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/GetByUsername", in, out, opts...)
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

func (c *userServiceClient) GetPublicUsers(ctx context.Context, in *GetPublicUsersRequest, opts ...grpc.CallOption) (*GetPublicUsersResponse, error) {
	out := new(GetPublicUsersResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/GetPublicUsers", in, out, opts...)
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

func (c *userServiceClient) GetEducation(ctx context.Context, in *GetEducationRequest, opts ...grpc.CallOption) (*GetEducationResponse, error) {
	out := new(GetEducationResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/GetEducation", in, out, opts...)
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

func (c *userServiceClient) GetExperience(ctx context.Context, in *GetExperienceRequest, opts ...grpc.CallOption) (*GetExperienceResponse, error) {
	out := new(GetExperienceResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/GetExperience", in, out, opts...)
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

func (c *userServiceClient) GetInterests(ctx context.Context, in *GetInterestsRequest, opts ...grpc.CallOption) (*GetInterestsResponse, error) {
	out := new(GetInterestsResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/GetInterests", in, out, opts...)
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

func (c *userServiceClient) GetSkills(ctx context.Context, in *GetSkillsRequest, opts ...grpc.CallOption) (*GetSkillsResponse, error) {
	out := new(GetSkillsResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/GetSkills", in, out, opts...)
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

func (c *userServiceClient) SearchProfiles(ctx context.Context, in *SearchProfilesRequest, opts ...grpc.CallOption) (*SearchProfilesResponse, error) {
	out := new(SearchProfilesResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/SearchProfiles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetPrivacy(ctx context.Context, in *SetPrivacyRequest, opts ...grpc.CallOption) (*SetPrivacyResponse, error) {
	out := new(SetPrivacyResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/SetPrivacy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SetApiKey(ctx context.Context, in *SetApiKeyRequest, opts ...grpc.CallOption) (*SetApiKeyResponse, error) {
	out := new(SetApiKeyResponse)
	err := c.cc.Invoke(ctx, "/users.UserService/SetApiKey", in, out, opts...)
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
	GetByUsername(context.Context, *GetByUsernameRequest) (*GetByUsernameResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetPublicUsers(context.Context, *GetPublicUsersRequest) (*GetPublicUsersResponse, error)
	Insert(context.Context, *InsertUserRequest) (*InsertUserResponse, error)
	EditUser(context.Context, *InsertUserRequest) (*EditUserResponse, error)
	GetEducation(context.Context, *GetEducationRequest) (*GetEducationResponse, error)
	AddEducation(context.Context, *AddEducationRequest) (*AddEducationResponse, error)
	DeleteEducation(context.Context, *DeleteEducationRequest) (*DeleteEducationResponse, error)
	GetExperience(context.Context, *GetExperienceRequest) (*GetExperienceResponse, error)
	AddExperience(context.Context, *AddExperienceRequest) (*AddExperienceResponse, error)
	DeleteExperience(context.Context, *DeleteExperienceRequest) (*DeleteExperienceResponse, error)
	GetInterests(context.Context, *GetInterestsRequest) (*GetInterestsResponse, error)
	AddInterest(context.Context, *AddInterestRequest) (*AddInterestResponse, error)
	DeleteInterest(context.Context, *DeleteInterestRequest) (*DeleteInterestResponse, error)
	GetSkills(context.Context, *GetSkillsRequest) (*GetSkillsResponse, error)
	AddSkill(context.Context, *AddSkillRequest) (*AddSkillResponse, error)
	DeleteSkill(context.Context, *DeleteSkillRequest) (*DeleteSkillResponse, error)
	SearchProfiles(context.Context, *SearchProfilesRequest) (*SearchProfilesResponse, error)
	SetPrivacy(context.Context, *SetPrivacyRequest) (*SetPrivacyResponse, error)
	SetApiKey(context.Context, *SetApiKeyRequest) (*SetApiKeyResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedUserServiceServer) GetByUsername(context.Context, *GetByUsernameRequest) (*GetByUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUsername not implemented")
}
func (UnimplementedUserServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUserServiceServer) GetPublicUsers(context.Context, *GetPublicUsersRequest) (*GetPublicUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublicUsers not implemented")
}
func (UnimplementedUserServiceServer) Insert(context.Context, *InsertUserRequest) (*InsertUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Insert not implemented")
}
func (UnimplementedUserServiceServer) EditUser(context.Context, *InsertUserRequest) (*EditUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUser not implemented")
}
func (UnimplementedUserServiceServer) GetEducation(context.Context, *GetEducationRequest) (*GetEducationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEducation not implemented")
}
func (UnimplementedUserServiceServer) AddEducation(context.Context, *AddEducationRequest) (*AddEducationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEducation not implemented")
}
func (UnimplementedUserServiceServer) DeleteEducation(context.Context, *DeleteEducationRequest) (*DeleteEducationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEducation not implemented")
}
func (UnimplementedUserServiceServer) GetExperience(context.Context, *GetExperienceRequest) (*GetExperienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExperience not implemented")
}
func (UnimplementedUserServiceServer) AddExperience(context.Context, *AddExperienceRequest) (*AddExperienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddExperience not implemented")
}
func (UnimplementedUserServiceServer) DeleteExperience(context.Context, *DeleteExperienceRequest) (*DeleteExperienceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteExperience not implemented")
}
func (UnimplementedUserServiceServer) GetInterests(context.Context, *GetInterestsRequest) (*GetInterestsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetInterests not implemented")
}
func (UnimplementedUserServiceServer) AddInterest(context.Context, *AddInterestRequest) (*AddInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddInterest not implemented")
}
func (UnimplementedUserServiceServer) DeleteInterest(context.Context, *DeleteInterestRequest) (*DeleteInterestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInterest not implemented")
}
func (UnimplementedUserServiceServer) GetSkills(context.Context, *GetSkillsRequest) (*GetSkillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkills not implemented")
}
func (UnimplementedUserServiceServer) AddSkill(context.Context, *AddSkillRequest) (*AddSkillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSkill not implemented")
}
func (UnimplementedUserServiceServer) DeleteSkill(context.Context, *DeleteSkillRequest) (*DeleteSkillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSkill not implemented")
}
func (UnimplementedUserServiceServer) SearchProfiles(context.Context, *SearchProfilesRequest) (*SearchProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchProfiles not implemented")
}
func (UnimplementedUserServiceServer) SetPrivacy(context.Context, *SetPrivacyRequest) (*SetPrivacyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPrivacy not implemented")
}
func (UnimplementedUserServiceServer) SetApiKey(context.Context, *SetApiKeyRequest) (*SetApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetApiKey not implemented")
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

func _UserService_GetByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetByUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByUsername(ctx, req.(*GetByUsernameRequest))
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

func _UserService_GetPublicUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetPublicUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetPublicUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetPublicUsers(ctx, req.(*GetPublicUsersRequest))
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

func _UserService_GetEducation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEducationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetEducation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetEducation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetEducation(ctx, req.(*GetEducationRequest))
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

func _UserService_GetExperience_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExperienceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetExperience(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetExperience",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetExperience(ctx, req.(*GetExperienceRequest))
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

func _UserService_GetInterests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInterestsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetInterests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetInterests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetInterests(ctx, req.(*GetInterestsRequest))
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

func _UserService_GetSkills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSkillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetSkills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/GetSkills",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetSkills(ctx, req.(*GetSkillsRequest))
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

func _UserService_SearchProfiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchProfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SearchProfiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/SearchProfiles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SearchProfiles(ctx, req.(*SearchProfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetPrivacy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPrivacyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetPrivacy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/SetPrivacy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetPrivacy(ctx, req.(*SetPrivacyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SetApiKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SetApiKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/users.UserService/SetApiKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SetApiKey(ctx, req.(*SetApiKeyRequest))
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
			MethodName: "GetByUsername",
			Handler:    _UserService_GetByUsername_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "GetPublicUsers",
			Handler:    _UserService_GetPublicUsers_Handler,
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
			MethodName: "GetEducation",
			Handler:    _UserService_GetEducation_Handler,
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
			MethodName: "GetExperience",
			Handler:    _UserService_GetExperience_Handler,
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
			MethodName: "GetInterests",
			Handler:    _UserService_GetInterests_Handler,
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
			MethodName: "GetSkills",
			Handler:    _UserService_GetSkills_Handler,
		},
		{
			MethodName: "AddSkill",
			Handler:    _UserService_AddSkill_Handler,
		},
		{
			MethodName: "DeleteSkill",
			Handler:    _UserService_DeleteSkill_Handler,
		},
		{
			MethodName: "SearchProfiles",
			Handler:    _UserService_SearchProfiles_Handler,
		},
		{
			MethodName: "SetPrivacy",
			Handler:    _UserService_SetPrivacy_Handler,
		},
		{
			MethodName: "SetApiKey",
			Handler:    _UserService_SetApiKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/proto/user_service/user_service.proto",
}
