// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: common/proto/authentication_service/authentication_service.proto

package authentications

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

// AuthenticationServiceClient is the client API for AuthenticationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticationServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	PasswordlessLogin(ctx context.Context, in *PasswordlessLoginRequest, opts ...grpc.CallOption) (*PasswordlessLoginResponse, error)
	GenerateVerificationToken(ctx context.Context, in *GenerateVerificationTokenRequest, opts ...grpc.CallOption) (*GenerateVerificationTokenResponse, error)
	AccountRecovery(ctx context.Context, in *AccountRecoveryRequest, opts ...grpc.CallOption) (*AccountRecoveryResponse, error)
	ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error)
	Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error)
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	EditUsername(ctx context.Context, in *EditUsernameRequest, opts ...grpc.CallOption) (*EditUsernameResponse, error)
	GetTwoFactorAuth(ctx context.Context, in *GetTwoFactorAuthRequest, opts ...grpc.CallOption) (*GetTwoFactorAuthResponse, error)
	ChangeTwoFactorAuth(ctx context.Context, in *ChangeTwoFactorAuthRequest, opts ...grpc.CallOption) (*ChangeTwoFactorAuthResponse, error)
	VerifyTwoFactorAuthToken(ctx context.Context, in *VerifyTwoFactorAuthTokenRequest, opts ...grpc.CallOption) (*VerifyTwoFactorAuthTokenResponse, error)
}

type authenticationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticationServiceClient(cc grpc.ClientConnInterface) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/authentications.AuthenticationService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) PasswordlessLogin(ctx context.Context, in *PasswordlessLoginRequest, opts ...grpc.CallOption) (*PasswordlessLoginResponse, error) {
	out := new(PasswordlessLoginResponse)
	err := c.cc.Invoke(ctx, "/authentications.AuthenticationService/PasswordlessLogin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) GenerateVerificationToken(ctx context.Context, in *GenerateVerificationTokenRequest, opts ...grpc.CallOption) (*GenerateVerificationTokenResponse, error) {
	out := new(GenerateVerificationTokenResponse)
	err := c.cc.Invoke(ctx, "/authentications.AuthenticationService/GenerateVerificationToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) AccountRecovery(ctx context.Context, in *AccountRecoveryRequest, opts ...grpc.CallOption) (*AccountRecoveryResponse, error) {
	out := new(AccountRecoveryResponse)
	err := c.cc.Invoke(ctx, "/authentications.AuthenticationService/AccountRecovery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) ChangePassword(ctx context.Context, in *ChangePasswordRequest, opts ...grpc.CallOption) (*ChangePasswordResponse, error) {
	out := new(ChangePasswordResponse)
	err := c.cc.Invoke(ctx, "/authentications.AuthenticationService/ChangePassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) Validate(ctx context.Context, in *ValidateRequest, opts ...grpc.CallOption) (*ValidateResponse, error) {
	out := new(ValidateResponse)
	err := c.cc.Invoke(ctx, "/authentications.AuthenticationService/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/authentications.AuthenticationService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) EditUsername(ctx context.Context, in *EditUsernameRequest, opts ...grpc.CallOption) (*EditUsernameResponse, error) {
	out := new(EditUsernameResponse)
	err := c.cc.Invoke(ctx, "/authentications.AuthenticationService/EditUsername", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) GetTwoFactorAuth(ctx context.Context, in *GetTwoFactorAuthRequest, opts ...grpc.CallOption) (*GetTwoFactorAuthResponse, error) {
	out := new(GetTwoFactorAuthResponse)
	err := c.cc.Invoke(ctx, "/authentications.AuthenticationService/GetTwoFactorAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) ChangeTwoFactorAuth(ctx context.Context, in *ChangeTwoFactorAuthRequest, opts ...grpc.CallOption) (*ChangeTwoFactorAuthResponse, error) {
	out := new(ChangeTwoFactorAuthResponse)
	err := c.cc.Invoke(ctx, "/authentications.AuthenticationService/ChangeTwoFactorAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticationServiceClient) VerifyTwoFactorAuthToken(ctx context.Context, in *VerifyTwoFactorAuthTokenRequest, opts ...grpc.CallOption) (*VerifyTwoFactorAuthTokenResponse, error) {
	out := new(VerifyTwoFactorAuthTokenResponse)
	err := c.cc.Invoke(ctx, "/authentications.AuthenticationService/VerifyTwoFactorAuthToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticationServiceServer is the server API for AuthenticationService service.
// All implementations must embed UnimplementedAuthenticationServiceServer
// for forward compatibility
type AuthenticationServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	PasswordlessLogin(context.Context, *PasswordlessLoginRequest) (*PasswordlessLoginResponse, error)
	GenerateVerificationToken(context.Context, *GenerateVerificationTokenRequest) (*GenerateVerificationTokenResponse, error)
	AccountRecovery(context.Context, *AccountRecoveryRequest) (*AccountRecoveryResponse, error)
	ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error)
	Validate(context.Context, *ValidateRequest) (*ValidateResponse, error)
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	EditUsername(context.Context, *EditUsernameRequest) (*EditUsernameResponse, error)
	GetTwoFactorAuth(context.Context, *GetTwoFactorAuthRequest) (*GetTwoFactorAuthResponse, error)
	ChangeTwoFactorAuth(context.Context, *ChangeTwoFactorAuthRequest) (*ChangeTwoFactorAuthResponse, error)
	VerifyTwoFactorAuthToken(context.Context, *VerifyTwoFactorAuthTokenRequest) (*VerifyTwoFactorAuthTokenResponse, error)
	mustEmbedUnimplementedAuthenticationServiceServer()
}

// UnimplementedAuthenticationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthenticationServiceServer struct {
}

func (UnimplementedAuthenticationServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthenticationServiceServer) PasswordlessLogin(context.Context, *PasswordlessLoginRequest) (*PasswordlessLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PasswordlessLogin not implemented")
}
func (UnimplementedAuthenticationServiceServer) GenerateVerificationToken(context.Context, *GenerateVerificationTokenRequest) (*GenerateVerificationTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateVerificationToken not implemented")
}
func (UnimplementedAuthenticationServiceServer) AccountRecovery(context.Context, *AccountRecoveryRequest) (*AccountRecoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccountRecovery not implemented")
}
func (UnimplementedAuthenticationServiceServer) ChangePassword(context.Context, *ChangePasswordRequest) (*ChangePasswordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePassword not implemented")
}
func (UnimplementedAuthenticationServiceServer) Validate(context.Context, *ValidateRequest) (*ValidateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedAuthenticationServiceServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthenticationServiceServer) EditUsername(context.Context, *EditUsernameRequest) (*EditUsernameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUsername not implemented")
}
func (UnimplementedAuthenticationServiceServer) GetTwoFactorAuth(context.Context, *GetTwoFactorAuthRequest) (*GetTwoFactorAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTwoFactorAuth not implemented")
}
func (UnimplementedAuthenticationServiceServer) ChangeTwoFactorAuth(context.Context, *ChangeTwoFactorAuthRequest) (*ChangeTwoFactorAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeTwoFactorAuth not implemented")
}
func (UnimplementedAuthenticationServiceServer) VerifyTwoFactorAuthToken(context.Context, *VerifyTwoFactorAuthTokenRequest) (*VerifyTwoFactorAuthTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyTwoFactorAuthToken not implemented")
}
func (UnimplementedAuthenticationServiceServer) mustEmbedUnimplementedAuthenticationServiceServer() {}

// UnsafeAuthenticationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticationServiceServer will
// result in compilation errors.
type UnsafeAuthenticationServiceServer interface {
	mustEmbedUnimplementedAuthenticationServiceServer()
}

func RegisterAuthenticationServiceServer(s grpc.ServiceRegistrar, srv AuthenticationServiceServer) {
	s.RegisterService(&AuthenticationService_ServiceDesc, srv)
}

func _AuthenticationService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentications.AuthenticationService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_PasswordlessLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PasswordlessLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).PasswordlessLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentications.AuthenticationService/PasswordlessLogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).PasswordlessLogin(ctx, req.(*PasswordlessLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_GenerateVerificationToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateVerificationTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).GenerateVerificationToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentications.AuthenticationService/GenerateVerificationToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).GenerateVerificationToken(ctx, req.(*GenerateVerificationTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_AccountRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountRecoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).AccountRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentications.AuthenticationService/AccountRecovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).AccountRecovery(ctx, req.(*AccountRecoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_ChangePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ChangePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentications.AuthenticationService/ChangePassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ChangePassword(ctx, req.(*ChangePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentications.AuthenticationService/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).Validate(ctx, req.(*ValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentications.AuthenticationService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_EditUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).EditUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentications.AuthenticationService/EditUsername",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).EditUsername(ctx, req.(*EditUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_GetTwoFactorAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTwoFactorAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).GetTwoFactorAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentications.AuthenticationService/GetTwoFactorAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).GetTwoFactorAuth(ctx, req.(*GetTwoFactorAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_ChangeTwoFactorAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeTwoFactorAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).ChangeTwoFactorAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentications.AuthenticationService/ChangeTwoFactorAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).ChangeTwoFactorAuth(ctx, req.(*ChangeTwoFactorAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthenticationService_VerifyTwoFactorAuthToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyTwoFactorAuthTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticationServiceServer).VerifyTwoFactorAuthToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authentications.AuthenticationService/VerifyTwoFactorAuthToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticationServiceServer).VerifyTwoFactorAuthToken(ctx, req.(*VerifyTwoFactorAuthTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthenticationService_ServiceDesc is the grpc.ServiceDesc for AuthenticationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthenticationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authentications.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthenticationService_Login_Handler,
		},
		{
			MethodName: "PasswordlessLogin",
			Handler:    _AuthenticationService_PasswordlessLogin_Handler,
		},
		{
			MethodName: "GenerateVerificationToken",
			Handler:    _AuthenticationService_GenerateVerificationToken_Handler,
		},
		{
			MethodName: "AccountRecovery",
			Handler:    _AuthenticationService_AccountRecovery_Handler,
		},
		{
			MethodName: "ChangePassword",
			Handler:    _AuthenticationService_ChangePassword_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _AuthenticationService_Validate_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _AuthenticationService_Register_Handler,
		},
		{
			MethodName: "EditUsername",
			Handler:    _AuthenticationService_EditUsername_Handler,
		},
		{
			MethodName: "GetTwoFactorAuth",
			Handler:    _AuthenticationService_GetTwoFactorAuth_Handler,
		},
		{
			MethodName: "ChangeTwoFactorAuth",
			Handler:    _AuthenticationService_ChangeTwoFactorAuth_Handler,
		},
		{
			MethodName: "VerifyTwoFactorAuthToken",
			Handler:    _AuthenticationService_VerifyTwoFactorAuthToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "common/proto/authentication_service/authentication_service.proto",
}
