// Version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: modules/user/userPb/userPb.proto

package NFT_Bidding_Platform

import (
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

// Structures
type UserProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email        string  `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Username     string  `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	WalletAmount float64 `protobuf:"fixed64,4,opt,name=walletAmount,proto3" json:"walletAmount,omitempty"`
	RoleCode     int32   `protobuf:"varint,5,opt,name=roleCode,proto3" json:"roleCode,omitempty"`
	CreatedAt    string  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt    string  `protobuf:"bytes,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UserProfile) Reset() {
	*x = UserProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_user_userPb_userPb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserProfile) ProtoMessage() {}

func (x *UserProfile) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserProfile.ProtoReflect.Descriptor instead.
func (*UserProfile) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{0}
}

func (x *UserProfile) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserProfile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserProfile) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserProfile) GetWalletAmount() float64 {
	if x != nil {
		return x.WalletAmount
	}
	return 0
}

func (x *UserProfile) GetRoleCode() int32 {
	if x != nil {
		return x.RoleCode
	}
	return 0
}

func (x *UserProfile) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserProfile) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type CredentialSearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CredentialSearchReq) Reset() {
	*x = CredentialSearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_user_userPb_userPb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CredentialSearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CredentialSearchReq) ProtoMessage() {}

func (x *CredentialSearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CredentialSearchReq.ProtoReflect.Descriptor instead.
func (*CredentialSearchReq) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{1}
}

func (x *CredentialSearchReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CredentialSearchReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type EmailSearchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *EmailSearchReq) Reset() {
	*x = EmailSearchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_user_userPb_userPb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmailSearchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmailSearchReq) ProtoMessage() {}

func (x *EmailSearchReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmailSearchReq.ProtoReflect.Descriptor instead.
func (*EmailSearchReq) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{2}
}

func (x *EmailSearchReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type FindOneUserProfileToRefreshReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *FindOneUserProfileToRefreshReq) Reset() {
	*x = FindOneUserProfileToRefreshReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_user_userPb_userPb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneUserProfileToRefreshReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneUserProfileToRefreshReq) ProtoMessage() {}

func (x *FindOneUserProfileToRefreshReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneUserProfileToRefreshReq.ProtoReflect.Descriptor instead.
func (*FindOneUserProfileToRefreshReq) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{3}
}

func (x *FindOneUserProfileToRefreshReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserWalletAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetUserWalletAccountReq) Reset() {
	*x = GetUserWalletAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_user_userPb_userPb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserWalletAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserWalletAccountReq) ProtoMessage() {}

func (x *GetUserWalletAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserWalletAccountReq.ProtoReflect.Descriptor instead.
func (*GetUserWalletAccountReq) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserWalletAccountReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserWalletAccountRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Balance float64 `protobuf:"fixed64,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *GetUserWalletAccountRes) Reset() {
	*x = GetUserWalletAccountRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_user_userPb_userPb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserWalletAccountRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserWalletAccountRes) ProtoMessage() {}

func (x *GetUserWalletAccountRes) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserWalletAccountRes.ProtoReflect.Descriptor instead.
func (*GetUserWalletAccountRes) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserWalletAccountRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserWalletAccountRes) GetBalance() float64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type DeductWalletAmountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *DeductWalletAmountReq) Reset() {
	*x = DeductWalletAmountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_user_userPb_userPb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeductWalletAmountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeductWalletAmountReq) ProtoMessage() {}

func (x *DeductWalletAmountReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeductWalletAmountReq.ProtoReflect.Descriptor instead.
func (*DeductWalletAmountReq) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{6}
}

func (x *DeductWalletAmountReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeductWalletAmountReq) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AddWalletAmountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AddWalletAmountReq) Reset() {
	*x = AddWalletAmountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_user_userPb_userPb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddWalletAmountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddWalletAmountReq) ProtoMessage() {}

func (x *AddWalletAmountReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_user_userPb_userPb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddWalletAmountReq.ProtoReflect.Descriptor instead.
func (*AddWalletAmountReq) Descriptor() ([]byte, []int) {
	return file_modules_user_userPb_userPb_proto_rawDescGZIP(), []int{7}
}

func (x *AddWalletAmountReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddWalletAmountReq) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

var File_modules_user_userPb_userPb_proto protoreflect.FileDescriptor

var file_modules_user_userPb_userPb_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x50, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x50, 0x62, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcd, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x77, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x47, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x26, 0x0a, 0x0e, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x38, 0x0a, 0x1e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x31, 0x0a,
	0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x4b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x47, 0x0a,
	0x15, 0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xa2, 0x03, 0x0a,
	0x0f, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x36, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x4c, 0x0a, 0x1b, 0x46, 0x69, 0x6e, 0x64,
	0x4f, 0x6e, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x6f,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x1f, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x6f, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x4a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18,
	0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x12, 0x33, 0x0a, 0x12, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x12, 0x44, 0x65, 0x64, 0x75, 0x63,
	0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e,
	0x44, 0x65, 0x64, 0x75, 0x63, 0x74, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x12,
	0x40, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x13, 0x2e, 0x41, 0x64, 0x64, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x75, 0x68, 0x61, 0x6d, 0x6d, 0x61, 0x64, 0x66, 0x61, 0x72, 0x68, 0x61, 0x6e, 0x6b, 0x74,
	0x2f, 0x4e, 0x46, 0x54, 0x2d, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x50, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_user_userPb_userPb_proto_rawDescOnce sync.Once
	file_modules_user_userPb_userPb_proto_rawDescData = file_modules_user_userPb_userPb_proto_rawDesc
)

func file_modules_user_userPb_userPb_proto_rawDescGZIP() []byte {
	file_modules_user_userPb_userPb_proto_rawDescOnce.Do(func() {
		file_modules_user_userPb_userPb_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_user_userPb_userPb_proto_rawDescData)
	})
	return file_modules_user_userPb_userPb_proto_rawDescData
}

var file_modules_user_userPb_userPb_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_modules_user_userPb_userPb_proto_goTypes = []interface{}{
	(*UserProfile)(nil),                    // 0: UserProfile
	(*CredentialSearchReq)(nil),            // 1: CredentialSearchReq
	(*EmailSearchReq)(nil),                 // 2: EmailSearchReq
	(*FindOneUserProfileToRefreshReq)(nil), // 3: FindOneUserProfileToRefreshReq
	(*GetUserWalletAccountReq)(nil),        // 4: GetUserWalletAccountReq
	(*GetUserWalletAccountRes)(nil),        // 5: GetUserWalletAccountRes
	(*DeductWalletAmountReq)(nil),          // 6: DeductWalletAmountReq
	(*AddWalletAmountReq)(nil),             // 7: AddWalletAmountReq
}
var file_modules_user_userPb_userPb_proto_depIdxs = []int32{
	1, // 0: UserGrpcService.CredentialSearch:input_type -> CredentialSearchReq
	3, // 1: UserGrpcService.FindOneUserProfileToRefresh:input_type -> FindOneUserProfileToRefreshReq
	4, // 2: UserGrpcService.GetUserWalletAccount:input_type -> GetUserWalletAccountReq
	2, // 3: UserGrpcService.FindOneUserProfile:input_type -> EmailSearchReq
	6, // 4: UserGrpcService.DeductWalletAmount:input_type -> DeductWalletAmountReq
	7, // 5: UserGrpcService.AddWalletAmount:input_type -> AddWalletAmountReq
	0, // 6: UserGrpcService.CredentialSearch:output_type -> UserProfile
	0, // 7: UserGrpcService.FindOneUserProfileToRefresh:output_type -> UserProfile
	5, // 8: UserGrpcService.GetUserWalletAccount:output_type -> GetUserWalletAccountRes
	0, // 9: UserGrpcService.FindOneUserProfile:output_type -> UserProfile
	5, // 10: UserGrpcService.DeductWalletAmount:output_type -> GetUserWalletAccountRes
	5, // 11: UserGrpcService.AddWalletAmount:output_type -> GetUserWalletAccountRes
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modules_user_userPb_userPb_proto_init() }
func file_modules_user_userPb_userPb_proto_init() {
	if File_modules_user_userPb_userPb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_user_userPb_userPb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserProfile); i {
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
		file_modules_user_userPb_userPb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CredentialSearchReq); i {
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
		file_modules_user_userPb_userPb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmailSearchReq); i {
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
		file_modules_user_userPb_userPb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneUserProfileToRefreshReq); i {
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
		file_modules_user_userPb_userPb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserWalletAccountReq); i {
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
		file_modules_user_userPb_userPb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserWalletAccountRes); i {
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
		file_modules_user_userPb_userPb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeductWalletAmountReq); i {
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
		file_modules_user_userPb_userPb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddWalletAmountReq); i {
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
			RawDescriptor: file_modules_user_userPb_userPb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_user_userPb_userPb_proto_goTypes,
		DependencyIndexes: file_modules_user_userPb_userPb_proto_depIdxs,
		MessageInfos:      file_modules_user_userPb_userPb_proto_msgTypes,
	}.Build()
	File_modules_user_userPb_userPb_proto = out.File
	file_modules_user_userPb_userPb_proto_rawDesc = nil
	file_modules_user_userPb_userPb_proto_goTypes = nil
	file_modules_user_userPb_userPb_proto_depIdxs = nil
}
