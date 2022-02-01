// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: api/wallet_pb/wallet.proto

package wallet_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ErrCode int32

const (
	ErrCode_ALREADY_EXECUTED ErrCode = 0
	ErrCode_BALANCE_TOO_LOW  ErrCode = 1
	ErrCode_UNKNOWN_WALLET   ErrCode = 2
	ErrCode_INTERNAL_ERROR   ErrCode = 3
)

// Enum value maps for ErrCode.
var (
	ErrCode_name = map[int32]string{
		0: "ALREADY_EXECUTED",
		1: "BALANCE_TOO_LOW",
		2: "UNKNOWN_WALLET",
		3: "INTERNAL_ERROR",
	}
	ErrCode_value = map[string]int32{
		"ALREADY_EXECUTED": 0,
		"BALANCE_TOO_LOW":  1,
		"UNKNOWN_WALLET":   2,
		"INTERNAL_ERROR":   3,
	}
)

func (x ErrCode) Enum() *ErrCode {
	p := new(ErrCode)
	*p = x
	return p
}

func (x ErrCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrCode) Descriptor() protoreflect.EnumDescriptor {
	return file_api_wallet_pb_wallet_proto_enumTypes[0].Descriptor()
}

func (ErrCode) Type() protoreflect.EnumType {
	return &file_api_wallet_pb_wallet_proto_enumTypes[0]
}

func (x ErrCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrCode.Descriptor instead.
func (ErrCode) EnumDescriptor() ([]byte, []int) {
	return file_api_wallet_pb_wallet_proto_rawDescGZIP(), []int{0}
}

type AlterBalanceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletId string `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"` // uuid
	Nonce    string `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Amount   int64  `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AlterBalanceReq) Reset() {
	*x = AlterBalanceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wallet_pb_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlterBalanceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterBalanceReq) ProtoMessage() {}

func (x *AlterBalanceReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_pb_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterBalanceReq.ProtoReflect.Descriptor instead.
func (*AlterBalanceReq) Descriptor() ([]byte, []int) {
	return file_api_wallet_pb_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *AlterBalanceReq) GetWalletId() string {
	if x != nil {
		return x.WalletId
	}
	return ""
}

func (x *AlterBalanceReq) GetNonce() string {
	if x != nil {
		return x.Nonce
	}
	return ""
}

func (x *AlterBalanceReq) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AlterBalanceRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletId string `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"` // uuid
	Balance  int64  `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *AlterBalanceRes) Reset() {
	*x = AlterBalanceRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wallet_pb_wallet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlterBalanceRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlterBalanceRes) ProtoMessage() {}

func (x *AlterBalanceRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_pb_wallet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlterBalanceRes.ProtoReflect.Descriptor instead.
func (*AlterBalanceRes) Descriptor() ([]byte, []int) {
	return file_api_wallet_pb_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *AlterBalanceRes) GetWalletId() string {
	if x != nil {
		return x.WalletId
	}
	return ""
}

func (x *AlterBalanceRes) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

type GetBalanceReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletId string `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"` // uuid
}

func (x *GetBalanceReq) Reset() {
	*x = GetBalanceReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wallet_pb_wallet_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceReq) ProtoMessage() {}

func (x *GetBalanceReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_pb_wallet_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceReq.ProtoReflect.Descriptor instead.
func (*GetBalanceReq) Descriptor() ([]byte, []int) {
	return file_api_wallet_pb_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *GetBalanceReq) GetWalletId() string {
	if x != nil {
		return x.WalletId
	}
	return ""
}

type GetBalanceRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletId  string                 `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"` // uuid
	Balance   int64                  `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetBalanceRes) Reset() {
	*x = GetBalanceRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wallet_pb_wallet_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBalanceRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBalanceRes) ProtoMessage() {}

func (x *GetBalanceRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_pb_wallet_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBalanceRes.ProtoReflect.Descriptor instead.
func (*GetBalanceRes) Descriptor() ([]byte, []int) {
	return file_api_wallet_pb_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *GetBalanceRes) GetWalletId() string {
	if x != nil {
		return x.WalletId
	}
	return ""
}

func (x *GetBalanceRes) GetBalance() int64 {
	if x != nil {
		return x.Balance
	}
	return 0
}

func (x *GetBalanceRes) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CreateWalletReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletId string `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"` // uuid
}

func (x *CreateWalletReq) Reset() {
	*x = CreateWalletReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wallet_pb_wallet_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletReq) ProtoMessage() {}

func (x *CreateWalletReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_pb_wallet_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletReq.ProtoReflect.Descriptor instead.
func (*CreateWalletReq) Descriptor() ([]byte, []int) {
	return file_api_wallet_pb_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *CreateWalletReq) GetWalletId() string {
	if x != nil {
		return x.WalletId
	}
	return ""
}

type CreateWalletRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletId string `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"` // uuid
}

func (x *CreateWalletRes) Reset() {
	*x = CreateWalletRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_wallet_pb_wallet_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateWalletRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateWalletRes) ProtoMessage() {}

func (x *CreateWalletRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_wallet_pb_wallet_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateWalletRes.ProtoReflect.Descriptor instead.
func (*CreateWalletRes) Descriptor() ([]byte, []int) {
	return file_api_wallet_pb_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *CreateWalletRes) GetWalletId() string {
	if x != nil {
		return x.WalletId
	}
	return ""
}

var File_api_wallet_pb_wallet_proto protoreflect.FileDescriptor

var file_api_wallet_pb_wallet_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x70, 0x62, 0x2f,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x0f, 0x41, 0x6c, 0x74, 0x65,
	0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x48, 0x0a, 0x0f, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0x2c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x22, 0x81,
	0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x2e, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x49, 0x64, 0x22, 0x2e, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x49, 0x64, 0x2a, 0x5c, 0x0a, 0x07, 0x45, 0x72, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x41, 0x4c, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x42, 0x41, 0x4c, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x54,
	0x4f, 0x4f, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x5f, 0x57, 0x41, 0x4c, 0x4c, 0x45, 0x54, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03,
	0x32, 0xe1, 0x01, 0x0a, 0x0d, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x46, 0x0a, 0x0c, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x12, 0x1a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x70, 0x62, 0x2e, 0x41,
	0x6c, 0x74, 0x65, 0x72, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a,
	0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x70, 0x62, 0x2e, 0x41, 0x6c, 0x74, 0x65, 0x72,
	0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x5f, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x12, 0x46, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x1a, 0x2e, 0x77,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x5f, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x3b, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_wallet_pb_wallet_proto_rawDescOnce sync.Once
	file_api_wallet_pb_wallet_proto_rawDescData = file_api_wallet_pb_wallet_proto_rawDesc
)

func file_api_wallet_pb_wallet_proto_rawDescGZIP() []byte {
	file_api_wallet_pb_wallet_proto_rawDescOnce.Do(func() {
		file_api_wallet_pb_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_wallet_pb_wallet_proto_rawDescData)
	})
	return file_api_wallet_pb_wallet_proto_rawDescData
}

var file_api_wallet_pb_wallet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_wallet_pb_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_api_wallet_pb_wallet_proto_goTypes = []interface{}{
	(ErrCode)(0),                  // 0: wallet_pb.ErrCode
	(*AlterBalanceReq)(nil),       // 1: wallet_pb.AlterBalanceReq
	(*AlterBalanceRes)(nil),       // 2: wallet_pb.AlterBalanceRes
	(*GetBalanceReq)(nil),         // 3: wallet_pb.GetBalanceReq
	(*GetBalanceRes)(nil),         // 4: wallet_pb.GetBalanceRes
	(*CreateWalletReq)(nil),       // 5: wallet_pb.CreateWalletReq
	(*CreateWalletRes)(nil),       // 6: wallet_pb.CreateWalletRes
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_api_wallet_pb_wallet_proto_depIdxs = []int32{
	7, // 0: wallet_pb.GetBalanceRes.updated_at:type_name -> google.protobuf.Timestamp
	1, // 1: wallet_pb.WalletService.AlterBalance:input_type -> wallet_pb.AlterBalanceReq
	3, // 2: wallet_pb.WalletService.GetBalance:input_type -> wallet_pb.GetBalanceReq
	5, // 3: wallet_pb.WalletService.CreateWallet:input_type -> wallet_pb.CreateWalletReq
	2, // 4: wallet_pb.WalletService.AlterBalance:output_type -> wallet_pb.AlterBalanceRes
	4, // 5: wallet_pb.WalletService.GetBalance:output_type -> wallet_pb.GetBalanceRes
	6, // 6: wallet_pb.WalletService.CreateWallet:output_type -> wallet_pb.CreateWalletRes
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_wallet_pb_wallet_proto_init() }
func file_api_wallet_pb_wallet_proto_init() {
	if File_api_wallet_pb_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_wallet_pb_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlterBalanceReq); i {
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
		file_api_wallet_pb_wallet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlterBalanceRes); i {
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
		file_api_wallet_pb_wallet_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceReq); i {
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
		file_api_wallet_pb_wallet_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBalanceRes); i {
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
		file_api_wallet_pb_wallet_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletReq); i {
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
		file_api_wallet_pb_wallet_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateWalletRes); i {
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
			RawDescriptor: file_api_wallet_pb_wallet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_wallet_pb_wallet_proto_goTypes,
		DependencyIndexes: file_api_wallet_pb_wallet_proto_depIdxs,
		EnumInfos:         file_api_wallet_pb_wallet_proto_enumTypes,
		MessageInfos:      file_api_wallet_pb_wallet_proto_msgTypes,
	}.Build()
	File_api_wallet_pb_wallet_proto = out.File
	file_api_wallet_pb_wallet_proto_rawDesc = nil
	file_api_wallet_pb_wallet_proto_goTypes = nil
	file_api_wallet_pb_wallet_proto_depIdxs = nil
}
