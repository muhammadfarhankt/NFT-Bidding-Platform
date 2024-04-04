// Version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: modules/inventory/inventoryPb/inventoryPb.proto

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
type IsAvailableToSellReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	NftId  string `protobuf:"bytes,2,opt,name=nftId,proto3" json:"nftId,omitempty"`
}

func (x *IsAvailableToSellReq) Reset() {
	*x = IsAvailableToSellReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_inventory_inventoryPb_inventoryPb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAvailableToSellReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAvailableToSellReq) ProtoMessage() {}

func (x *IsAvailableToSellReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_inventory_inventoryPb_inventoryPb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAvailableToSellReq.ProtoReflect.Descriptor instead.
func (*IsAvailableToSellReq) Descriptor() ([]byte, []int) {
	return file_modules_inventory_inventoryPb_inventoryPb_proto_rawDescGZIP(), []int{0}
}

func (x *IsAvailableToSellReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *IsAvailableToSellReq) GetNftId() string {
	if x != nil {
		return x.NftId
	}
	return ""
}

type IsAvailableToSellRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAvailable bool `protobuf:"varint,1,opt,name=isAvailable,proto3" json:"isAvailable,omitempty"`
}

func (x *IsAvailableToSellRes) Reset() {
	*x = IsAvailableToSellRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_inventory_inventoryPb_inventoryPb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsAvailableToSellRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsAvailableToSellRes) ProtoMessage() {}

func (x *IsAvailableToSellRes) ProtoReflect() protoreflect.Message {
	mi := &file_modules_inventory_inventoryPb_inventoryPb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsAvailableToSellRes.ProtoReflect.Descriptor instead.
func (*IsAvailableToSellRes) Descriptor() ([]byte, []int) {
	return file_modules_inventory_inventoryPb_inventoryPb_proto_rawDescGZIP(), []int{1}
}

func (x *IsAvailableToSellRes) GetIsAvailable() bool {
	if x != nil {
		return x.IsAvailable
	}
	return false
}

var File_modules_inventory_inventoryPb_inventoryPb_proto protoreflect.FileDescriptor

var file_modules_inventory_inventoryPb_inventoryPb_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x62, 0x2f,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x50, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x44, 0x0a, 0x14, 0x49, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65,
	0x54, 0x6f, 0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x66, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6e, 0x66, 0x74, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x14, 0x49, 0x73, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x69, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x32, 0x5b, 0x0a, 0x14, 0x49, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x47, 0x72,
	0x70, 0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x11, 0x49, 0x73, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x53, 0x65, 0x6c, 0x6c, 0x12, 0x15,
	0x2e, 0x49, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x6f, 0x53, 0x65,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x49, 0x73, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x54, 0x6f, 0x53, 0x65, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x68,
	0x61, 0x6d, 0x6d, 0x61, 0x64, 0x66, 0x61, 0x72, 0x68, 0x61, 0x6e, 0x6b, 0x74, 0x2f, 0x4e, 0x46,
	0x54, 0x2d, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_inventory_inventoryPb_inventoryPb_proto_rawDescOnce sync.Once
	file_modules_inventory_inventoryPb_inventoryPb_proto_rawDescData = file_modules_inventory_inventoryPb_inventoryPb_proto_rawDesc
)

func file_modules_inventory_inventoryPb_inventoryPb_proto_rawDescGZIP() []byte {
	file_modules_inventory_inventoryPb_inventoryPb_proto_rawDescOnce.Do(func() {
		file_modules_inventory_inventoryPb_inventoryPb_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_inventory_inventoryPb_inventoryPb_proto_rawDescData)
	})
	return file_modules_inventory_inventoryPb_inventoryPb_proto_rawDescData
}

var file_modules_inventory_inventoryPb_inventoryPb_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_modules_inventory_inventoryPb_inventoryPb_proto_goTypes = []interface{}{
	(*IsAvailableToSellReq)(nil), // 0: IsAvailableToSellReq
	(*IsAvailableToSellRes)(nil), // 1: IsAvailableToSellRes
}
var file_modules_inventory_inventoryPb_inventoryPb_proto_depIdxs = []int32{
	0, // 0: InventoryGrpcService.IsAvailableToSell:input_type -> IsAvailableToSellReq
	1, // 1: InventoryGrpcService.IsAvailableToSell:output_type -> IsAvailableToSellRes
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_modules_inventory_inventoryPb_inventoryPb_proto_init() }
func file_modules_inventory_inventoryPb_inventoryPb_proto_init() {
	if File_modules_inventory_inventoryPb_inventoryPb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_inventory_inventoryPb_inventoryPb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAvailableToSellReq); i {
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
		file_modules_inventory_inventoryPb_inventoryPb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsAvailableToSellRes); i {
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
			RawDescriptor: file_modules_inventory_inventoryPb_inventoryPb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_inventory_inventoryPb_inventoryPb_proto_goTypes,
		DependencyIndexes: file_modules_inventory_inventoryPb_inventoryPb_proto_depIdxs,
		MessageInfos:      file_modules_inventory_inventoryPb_inventoryPb_proto_msgTypes,
	}.Build()
	File_modules_inventory_inventoryPb_inventoryPb_proto = out.File
	file_modules_inventory_inventoryPb_inventoryPb_proto_rawDesc = nil
	file_modules_inventory_inventoryPb_inventoryPb_proto_goTypes = nil
	file_modules_inventory_inventoryPb_inventoryPb_proto_depIdxs = nil
}