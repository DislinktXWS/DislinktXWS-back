// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: common/proto/user_service/user_service.proto

package users

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_user_service_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_user_service_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_user_service_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_user_service_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_user_service_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_user_service_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_user_service_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_user_service_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_user_service_user_service_proto_rawDescGZIP(), []int{2}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_user_service_user_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_user_service_user_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_user_service_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetAllResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type InsertUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *InsertUserRequest) Reset() {
	*x = InsertUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_user_service_user_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertUserRequest) ProtoMessage() {}

func (x *InsertUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_user_service_user_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertUserRequest.ProtoReflect.Descriptor instead.
func (*InsertUserRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_user_service_user_service_proto_rawDescGZIP(), []int{4}
}

func (x *InsertUserRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type InsertUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InsertUserResponse) Reset() {
	*x = InsertUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_user_service_user_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertUserResponse) ProtoMessage() {}

func (x *InsertUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_user_service_user_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertUserResponse.ProtoReflect.Descriptor instead.
func (*InsertUserResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_user_service_user_service_proto_rawDescGZIP(), []int{5}
}

type AddEducationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Education *Education `protobuf:"bytes,1,opt,name=education,proto3" json:"education,omitempty"`
	Id        string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddEducationRequest) Reset() {
	*x = AddEducationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_user_service_user_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEducationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEducationRequest) ProtoMessage() {}

func (x *AddEducationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_user_service_user_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEducationRequest.ProtoReflect.Descriptor instead.
func (*AddEducationRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_user_service_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *AddEducationRequest) GetEducation() *Education {
	if x != nil {
		return x.Education
	}
	return nil
}

func (x *AddEducationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AddEducationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddEducationResponse) Reset() {
	*x = AddEducationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_user_service_user_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEducationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEducationResponse) ProtoMessage() {}

func (x *AddEducationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_user_service_user_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEducationResponse.ProtoReflect.Descriptor instead.
func (*AddEducationResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_user_service_user_service_proto_rawDescGZIP(), []int{7}
}

type DeleteEducationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Index uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *DeleteEducationRequest) Reset() {
	*x = DeleteEducationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_user_service_user_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEducationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEducationRequest) ProtoMessage() {}

func (x *DeleteEducationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_user_service_user_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEducationRequest.ProtoReflect.Descriptor instead.
func (*DeleteEducationRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_user_service_user_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteEducationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteEducationRequest) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type DeleteEducationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteEducationResponse) Reset() {
	*x = DeleteEducationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_user_service_user_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEducationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEducationResponse) ProtoMessage() {}

func (x *DeleteEducationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_user_service_user_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEducationResponse.ProtoReflect.Descriptor instead.
func (*DeleteEducationResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_user_service_user_service_proto_rawDescGZIP(), []int{9}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname     string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Username    string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Password    string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	DateOfBirth string `protobuf:"bytes,6,opt,name=dateOfBirth,proto3" json:"dateOfBirth,omitempty"`
	Gender      string `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Email       string `protobuf:"bytes,8,opt,name=email,proto3" json:"email,omitempty"`
	Phone       string `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_user_service_user_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_user_service_user_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_common_proto_user_service_user_service_proto_rawDescGZIP(), []int{10}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetDateOfBirth() string {
	if x != nil {
		return x.DateOfBirth
	}
	return ""
}

func (x *User) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type Education struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	School       string  `protobuf:"bytes,1,opt,name=school,proto3" json:"school,omitempty"`
	Degree       string  `protobuf:"bytes,2,opt,name=degree,proto3" json:"degree,omitempty"`
	FieldOfStudy string  `protobuf:"bytes,3,opt,name=fieldOfStudy,proto3" json:"fieldOfStudy,omitempty"`
	StartDate    string  `protobuf:"bytes,4,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate      string  `protobuf:"bytes,5,opt,name=endDate,proto3" json:"endDate,omitempty"`
	Grade        float32 `protobuf:"fixed32,6,opt,name=grade,proto3" json:"grade,omitempty"`
	Description  string  `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Education) Reset() {
	*x = Education{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_user_service_user_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Education) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Education) ProtoMessage() {}

func (x *Education) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_user_service_user_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Education.ProtoReflect.Descriptor instead.
func (*Education) Descriptor() ([]byte, []int) {
	return file_common_proto_user_service_user_service_proto_rawDescGZIP(), []int{11}
}

func (x *Education) GetSchool() string {
	if x != nil {
		return x.School
	}
	return ""
}

func (x *Education) GetDegree() string {
	if x != nil {
		return x.Degree
	}
	return ""
}

func (x *Education) GetFieldOfStudy() string {
	if x != nil {
		return x.FieldOfStudy
	}
	return ""
}

func (x *Education) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Education) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *Education) GetGrade() float32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *Education) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_common_proto_user_service_user_service_proto protoreflect.FileDescriptor

var file_common_proto_user_service_user_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x2e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x33, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x34, 0x0a, 0x11, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x14, 0x0a,
	0x12, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x55, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2e, 0x0a, 0x09, 0x65, 0x64,
	0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x09, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x41, 0x64,
	0x64, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x3e, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x64, 0x75, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xe2, 0x01,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x09, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x63, 0x68, 0x6f, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x67, 0x72,
	0x65, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x66, 0x53, 0x74, 0x75, 0x64, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x66, 0x53,
	0x74, 0x75, 0x64, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61,
	0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0xc8, 0x04, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x11, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4a, 0x0a, 0x06,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x12, 0x58, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65,
	0x72, 0x74, 0x12, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22,
	0x0b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x04, 0x75, 0x73,
	0x65, 0x72, 0x12, 0x5a, 0x0a, 0x08, 0x45, 0x64, 0x69, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x1a, 0x0b, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x76,
	0x0a, 0x0c, 0x41, 0x64, 0x64, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x73, 0x2e, 0x41, 0x64, 0x64, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x27, 0x1a,
	0x1a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x65, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x09, 0x65, 0x64, 0x75,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x77, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x2a, 0x1d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x65, 0x64, 0x75, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x7d, 0x42,
	0x15, 0x5a, 0x13, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_user_service_user_service_proto_rawDescOnce sync.Once
	file_common_proto_user_service_user_service_proto_rawDescData = file_common_proto_user_service_user_service_proto_rawDesc
)

func file_common_proto_user_service_user_service_proto_rawDescGZIP() []byte {
	file_common_proto_user_service_user_service_proto_rawDescOnce.Do(func() {
		file_common_proto_user_service_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_user_service_user_service_proto_rawDescData)
	})
	return file_common_proto_user_service_user_service_proto_rawDescData
}

var file_common_proto_user_service_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_common_proto_user_service_user_service_proto_goTypes = []interface{}{
	(*GetRequest)(nil),              // 0: users.GetRequest
	(*GetResponse)(nil),             // 1: users.GetResponse
	(*GetAllRequest)(nil),           // 2: users.GetAllRequest
	(*GetAllResponse)(nil),          // 3: users.GetAllResponse
	(*InsertUserRequest)(nil),       // 4: users.InsertUserRequest
	(*InsertUserResponse)(nil),      // 5: users.InsertUserResponse
	(*AddEducationRequest)(nil),     // 6: users.AddEducationRequest
	(*AddEducationResponse)(nil),    // 7: users.AddEducationResponse
	(*DeleteEducationRequest)(nil),  // 8: users.DeleteEducationRequest
	(*DeleteEducationResponse)(nil), // 9: users.DeleteEducationResponse
	(*User)(nil),                    // 10: users.User
	(*Education)(nil),               // 11: users.Education
}
var file_common_proto_user_service_user_service_proto_depIdxs = []int32{
	10, // 0: users.GetResponse.user:type_name -> users.User
	10, // 1: users.GetAllResponse.users:type_name -> users.User
	10, // 2: users.InsertUserRequest.user:type_name -> users.User
	11, // 3: users.AddEducationRequest.education:type_name -> users.Education
	0,  // 4: users.UserService.Get:input_type -> users.GetRequest
	2,  // 5: users.UserService.GetAll:input_type -> users.GetAllRequest
	4,  // 6: users.UserService.Insert:input_type -> users.InsertUserRequest
	4,  // 7: users.UserService.EditUser:input_type -> users.InsertUserRequest
	6,  // 8: users.UserService.AddEducation:input_type -> users.AddEducationRequest
	8,  // 9: users.UserService.DeleteEducation:input_type -> users.DeleteEducationRequest
	1,  // 10: users.UserService.Get:output_type -> users.GetResponse
	3,  // 11: users.UserService.GetAll:output_type -> users.GetAllResponse
	5,  // 12: users.UserService.Insert:output_type -> users.InsertUserResponse
	5,  // 13: users.UserService.EditUser:output_type -> users.InsertUserResponse
	7,  // 14: users.UserService.AddEducation:output_type -> users.AddEducationResponse
	9,  // 15: users.UserService.DeleteEducation:output_type -> users.DeleteEducationResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_common_proto_user_service_user_service_proto_init() }
func file_common_proto_user_service_user_service_proto_init() {
	if File_common_proto_user_service_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_user_service_user_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_user_service_user_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_user_service_user_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_user_service_user_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_user_service_user_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertUserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_user_service_user_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertUserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_user_service_user_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEducationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_user_service_user_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEducationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_user_service_user_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEducationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_user_service_user_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEducationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_user_service_user_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_common_proto_user_service_user_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Education); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_user_service_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_proto_user_service_user_service_proto_goTypes,
		DependencyIndexes: file_common_proto_user_service_user_service_proto_depIdxs,
		MessageInfos:      file_common_proto_user_service_user_service_proto_msgTypes,
	}.Build()
	File_common_proto_user_service_user_service_proto = out.File
	file_common_proto_user_service_user_service_proto_rawDesc = nil
	file_common_proto_user_service_user_service_proto_goTypes = nil
	file_common_proto_user_service_user_service_proto_depIdxs = nil
}
