// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: common/proto/business_offer_service/business_offer_service.proto

package business_offer

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

type Skill_SkillProficiency int32

const (
	Skill_novice           Skill_SkillProficiency = 0
	Skill_advancedBeginner Skill_SkillProficiency = 1
	Skill_proficient       Skill_SkillProficiency = 2
	Skill_expert           Skill_SkillProficiency = 3
	Skill_master           Skill_SkillProficiency = 4
)

// Enum value maps for Skill_SkillProficiency.
var (
	Skill_SkillProficiency_name = map[int32]string{
		0: "novice",
		1: "advancedBeginner",
		2: "proficient",
		3: "expert",
		4: "master",
	}
	Skill_SkillProficiency_value = map[string]int32{
		"novice":           0,
		"advancedBeginner": 1,
		"proficient":       2,
		"expert":           3,
		"master":           4,
	}
)

func (x Skill_SkillProficiency) Enum() *Skill_SkillProficiency {
	p := new(Skill_SkillProficiency)
	*p = x
	return p
}

func (x Skill_SkillProficiency) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Skill_SkillProficiency) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_business_offer_service_business_offer_service_proto_enumTypes[0].Descriptor()
}

func (Skill_SkillProficiency) Type() protoreflect.EnumType {
	return &file_common_proto_business_offer_service_business_offer_service_proto_enumTypes[0]
}

func (x Skill_SkillProficiency) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Skill_SkillProficiency.Descriptor instead.
func (Skill_SkillProficiency) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_business_offer_service_business_offer_service_proto_rawDescGZIP(), []int{5, 0}
}

type InsertOfferRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offer *BusinessOffer `protobuf:"bytes,1,opt,name=offer,proto3" json:"offer,omitempty"`
}

func (x *InsertOfferRequest) Reset() {
	*x = InsertOfferRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertOfferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertOfferRequest) ProtoMessage() {}

func (x *InsertOfferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertOfferRequest.ProtoReflect.Descriptor instead.
func (*InsertOfferRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_business_offer_service_business_offer_service_proto_rawDescGZIP(), []int{0}
}

func (x *InsertOfferRequest) GetOffer() *BusinessOffer {
	if x != nil {
		return x.Offer
	}
	return nil
}

type InsertOfferResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InsertOfferResponse) Reset() {
	*x = InsertOfferResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertOfferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertOfferResponse) ProtoMessage() {}

func (x *InsertOfferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertOfferResponse.ProtoReflect.Descriptor instead.
func (*InsertOfferResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_business_offer_service_business_offer_service_proto_rawDescGZIP(), []int{1}
}

type InsertSkillsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Skill *Skill `protobuf:"bytes,1,opt,name=skill,proto3" json:"skill,omitempty"`
}

func (x *InsertSkillsRequest) Reset() {
	*x = InsertSkillsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertSkillsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertSkillsRequest) ProtoMessage() {}

func (x *InsertSkillsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertSkillsRequest.ProtoReflect.Descriptor instead.
func (*InsertSkillsRequest) Descriptor() ([]byte, []int) {
	return file_common_proto_business_offer_service_business_offer_service_proto_rawDescGZIP(), []int{2}
}

func (x *InsertSkillsRequest) GetSkill() *Skill {
	if x != nil {
		return x.Skill
	}
	return nil
}

type InsertSkillsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InsertSkillsResponse) Reset() {
	*x = InsertSkillsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertSkillsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertSkillsResponse) ProtoMessage() {}

func (x *InsertSkillsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertSkillsResponse.ProtoReflect.Descriptor instead.
func (*InsertSkillsResponse) Descriptor() ([]byte, []int) {
	return file_common_proto_business_offer_service_business_offer_service_proto_rawDescGZIP(), []int{3}
}

type BusinessOffer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId    string `protobuf:"bytes,1,opt,name=authorId,proto3" json:"authorId,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Position    string `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Industry    string `protobuf:"bytes,6,opt,name=industry,proto3" json:"industry,omitempty"`
}

func (x *BusinessOffer) Reset() {
	*x = BusinessOffer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessOffer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessOffer) ProtoMessage() {}

func (x *BusinessOffer) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessOffer.ProtoReflect.Descriptor instead.
func (*BusinessOffer) Descriptor() ([]byte, []int) {
	return file_common_proto_business_offer_service_business_offer_service_proto_rawDescGZIP(), []int{4}
}

func (x *BusinessOffer) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *BusinessOffer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BusinessOffer) GetPosition() string {
	if x != nil {
		return x.Position
	}
	return ""
}

func (x *BusinessOffer) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BusinessOffer) GetIndustry() string {
	if x != nil {
		return x.Industry
	}
	return ""
}

type Skill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthorId    string                 `protobuf:"bytes,1,opt,name=authorId,proto3" json:"authorId,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Proficiency Skill_SkillProficiency `protobuf:"varint,3,opt,name=proficiency,proto3,enum=business_offer.Skill_SkillProficiency" json:"proficiency,omitempty"`
}

func (x *Skill) Reset() {
	*x = Skill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Skill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Skill) ProtoMessage() {}

func (x *Skill) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Skill.ProtoReflect.Descriptor instead.
func (*Skill) Descriptor() ([]byte, []int) {
	return file_common_proto_business_offer_service_business_offer_service_proto_rawDescGZIP(), []int{5}
}

func (x *Skill) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Skill) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Skill) GetProficiency() Skill_SkillProficiency {
	if x != nil {
		return x.Proficiency
	}
	return Skill_novice
}

var File_common_proto_business_offer_service_business_offer_service_proto protoreflect.FileDescriptor

var file_common_proto_business_offer_service_business_offer_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x49, 0x0a, 0x12, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x42, 0x0a, 0x13, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x53, 0x6b, 0x69, 0x6c,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x6b, 0x69,
	0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x52,
	0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x22, 0x16, 0x0a, 0x14, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x99,
	0x01, 0x0a, 0x0d, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4f, 0x66, 0x66, 0x65, 0x72,
	0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x79, 0x22, 0xdf, 0x01, 0x0a, 0x05, 0x53,
	0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x48, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65,
	0x6e, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x2e, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x5c,
	0x0a, 0x10, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x0a, 0x0a, 0x06, 0x6e, 0x6f, 0x76, 0x69, 0x63, 0x65, 0x10, 0x00, 0x12, 0x14,
	0x0a, 0x10, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x63, 0x69, 0x65,
	0x6e, 0x74, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x65, 0x72, 0x74, 0x10, 0x03,
	0x12, 0x0a, 0x0a, 0x06, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x10, 0x04, 0x32, 0x98, 0x02, 0x0a,
	0x15, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a, 0x13, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x12, 0x22, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x49,
	0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x23, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6f, 0x66, 0x66,
	0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22, 0x0d,
	0x2f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x3a, 0x05, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x12, 0x80, 0x01, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x4f, 0x66, 0x66, 0x65, 0x72, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x12, 0x23,
	0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6f,
	0x66, 0x66, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x53, 0x6b, 0x69, 0x6c, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x16, 0x22, 0x0d, 0x2f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x73, 0x2f, 0x73, 0x6b, 0x69, 0x6c, 0x6c,
	0x3a, 0x05, 0x73, 0x6b, 0x69, 0x6c, 0x6c, 0x42, 0x1e, 0x5a, 0x1c, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x5f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_proto_business_offer_service_business_offer_service_proto_rawDescOnce sync.Once
	file_common_proto_business_offer_service_business_offer_service_proto_rawDescData = file_common_proto_business_offer_service_business_offer_service_proto_rawDesc
)

func file_common_proto_business_offer_service_business_offer_service_proto_rawDescGZIP() []byte {
	file_common_proto_business_offer_service_business_offer_service_proto_rawDescOnce.Do(func() {
		file_common_proto_business_offer_service_business_offer_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_business_offer_service_business_offer_service_proto_rawDescData)
	})
	return file_common_proto_business_offer_service_business_offer_service_proto_rawDescData
}

var file_common_proto_business_offer_service_business_offer_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_common_proto_business_offer_service_business_offer_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_common_proto_business_offer_service_business_offer_service_proto_goTypes = []interface{}{
	(Skill_SkillProficiency)(0),  // 0: business_offer.Skill.SkillProficiency
	(*InsertOfferRequest)(nil),   // 1: business_offer.InsertOfferRequest
	(*InsertOfferResponse)(nil),  // 2: business_offer.InsertOfferResponse
	(*InsertSkillsRequest)(nil),  // 3: business_offer.InsertSkillsRequest
	(*InsertSkillsResponse)(nil), // 4: business_offer.InsertSkillsResponse
	(*BusinessOffer)(nil),        // 5: business_offer.BusinessOffer
	(*Skill)(nil),                // 6: business_offer.Skill
}
var file_common_proto_business_offer_service_business_offer_service_proto_depIdxs = []int32{
	5, // 0: business_offer.InsertOfferRequest.offer:type_name -> business_offer.BusinessOffer
	6, // 1: business_offer.InsertSkillsRequest.skill:type_name -> business_offer.Skill
	0, // 2: business_offer.Skill.proficiency:type_name -> business_offer.Skill.SkillProficiency
	1, // 3: business_offer.BusinessOffersService.InsertBusinessOffer:input_type -> business_offer.InsertOfferRequest
	3, // 4: business_offer.BusinessOffersService.AddBusinessOfferSkill:input_type -> business_offer.InsertSkillsRequest
	2, // 5: business_offer.BusinessOffersService.InsertBusinessOffer:output_type -> business_offer.InsertOfferResponse
	4, // 6: business_offer.BusinessOffersService.AddBusinessOfferSkill:output_type -> business_offer.InsertSkillsResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_common_proto_business_offer_service_business_offer_service_proto_init() }
func file_common_proto_business_offer_service_business_offer_service_proto_init() {
	if File_common_proto_business_offer_service_business_offer_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertOfferRequest); i {
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
		file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertOfferResponse); i {
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
		file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertSkillsRequest); i {
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
		file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertSkillsResponse); i {
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
		file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusinessOffer); i {
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
		file_common_proto_business_offer_service_business_offer_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Skill); i {
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
			RawDescriptor: file_common_proto_business_offer_service_business_offer_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_proto_business_offer_service_business_offer_service_proto_goTypes,
		DependencyIndexes: file_common_proto_business_offer_service_business_offer_service_proto_depIdxs,
		EnumInfos:         file_common_proto_business_offer_service_business_offer_service_proto_enumTypes,
		MessageInfos:      file_common_proto_business_offer_service_business_offer_service_proto_msgTypes,
	}.Build()
	File_common_proto_business_offer_service_business_offer_service_proto = out.File
	file_common_proto_business_offer_service_business_offer_service_proto_rawDesc = nil
	file_common_proto_business_offer_service_business_offer_service_proto_goTypes = nil
	file_common_proto_business_offer_service_business_offer_service_proto_depIdxs = nil
}
