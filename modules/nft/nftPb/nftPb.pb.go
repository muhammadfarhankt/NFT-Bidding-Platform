// Version

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: modules/nft/nftPb/nftPb.proto

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
type FindNftsInIdsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *FindNftsInIdsReq) Reset() {
	*x = FindNftsInIdsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_nft_nftPb_nftPb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNftsInIdsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNftsInIdsReq) ProtoMessage() {}

func (x *FindNftsInIdsReq) ProtoReflect() protoreflect.Message {
	mi := &file_modules_nft_nftPb_nftPb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNftsInIdsReq.ProtoReflect.Descriptor instead.
func (*FindNftsInIdsReq) Descriptor() ([]byte, []int) {
	return file_modules_nft_nftPb_nftPb_proto_rawDescGZIP(), []int{0}
}

func (x *FindNftsInIdsReq) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type FindNftsInIdsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nfts []*Nft `protobuf:"bytes,1,rep,name=nfts,proto3" json:"nfts,omitempty"`
}

func (x *FindNftsInIdsRes) Reset() {
	*x = FindNftsInIdsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_nft_nftPb_nftPb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindNftsInIdsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindNftsInIdsRes) ProtoMessage() {}

func (x *FindNftsInIdsRes) ProtoReflect() protoreflect.Message {
	mi := &file_modules_nft_nftPb_nftPb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindNftsInIdsRes.ProtoReflect.Descriptor instead.
func (*FindNftsInIdsRes) Descriptor() ([]byte, []int) {
	return file_modules_nft_nftPb_nftPb_proto_rawDescGZIP(), []int{1}
}

func (x *FindNftsInIdsRes) GetNfts() []*Nft {
	if x != nil {
		return x.Nfts
	}
	return nil
}

type Nft struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Price    float64 `protobuf:"fixed64,3,opt,name=price,proto3" json:"price,omitempty"`
	ImageUrl string  `protobuf:"bytes,4,opt,name=imageUrl,proto3" json:"imageUrl,omitempty"`
}

func (x *Nft) Reset() {
	*x = Nft{}
	if protoimpl.UnsafeEnabled {
		mi := &file_modules_nft_nftPb_nftPb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nft) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nft) ProtoMessage() {}

func (x *Nft) ProtoReflect() protoreflect.Message {
	mi := &file_modules_nft_nftPb_nftPb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nft.ProtoReflect.Descriptor instead.
func (*Nft) Descriptor() ([]byte, []int) {
	return file_modules_nft_nftPb_nftPb_proto_rawDescGZIP(), []int{2}
}

func (x *Nft) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Nft) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Nft) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Nft) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

var File_modules_nft_nftPb_nftPb_proto protoreflect.FileDescriptor

var file_modules_nft_nftPb_nftPb_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x2f, 0x6e, 0x66, 0x74, 0x2f, 0x6e, 0x66,
	0x74, 0x50, 0x62, 0x2f, 0x6e, 0x66, 0x74, 0x50, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x24, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x66, 0x74, 0x73, 0x49, 0x6e, 0x49, 0x64, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0x2c, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x66, 0x74,
	0x73, 0x49, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x66, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x4e, 0x66, 0x74, 0x52, 0x04, 0x6e,
	0x66, 0x74, 0x73, 0x22, 0x5d, 0x0a, 0x03, 0x4e, 0x66, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55,
	0x72, 0x6c, 0x32, 0x47, 0x0a, 0x0e, 0x4e, 0x66, 0x74, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x66, 0x74, 0x73,
	0x49, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x11, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4e, 0x66, 0x74, 0x73,
	0x49, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4e,
	0x66, 0x74, 0x73, 0x49, 0x6e, 0x49, 0x64, 0x73, 0x52, 0x65, 0x73, 0x42, 0x32, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x75, 0x68, 0x61, 0x6d, 0x6d,
	0x61, 0x64, 0x66, 0x61, 0x72, 0x68, 0x61, 0x6e, 0x6b, 0x74, 0x2f, 0x4e, 0x46, 0x54, 0x2d, 0x42,
	0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x2d, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_modules_nft_nftPb_nftPb_proto_rawDescOnce sync.Once
	file_modules_nft_nftPb_nftPb_proto_rawDescData = file_modules_nft_nftPb_nftPb_proto_rawDesc
)

func file_modules_nft_nftPb_nftPb_proto_rawDescGZIP() []byte {
	file_modules_nft_nftPb_nftPb_proto_rawDescOnce.Do(func() {
		file_modules_nft_nftPb_nftPb_proto_rawDescData = protoimpl.X.CompressGZIP(file_modules_nft_nftPb_nftPb_proto_rawDescData)
	})
	return file_modules_nft_nftPb_nftPb_proto_rawDescData
}

var file_modules_nft_nftPb_nftPb_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_modules_nft_nftPb_nftPb_proto_goTypes = []interface{}{
	(*FindNftsInIdsReq)(nil), // 0: FindNftsInIdsReq
	(*FindNftsInIdsRes)(nil), // 1: FindNftsInIdsRes
	(*Nft)(nil),              // 2: Nft
}
var file_modules_nft_nftPb_nftPb_proto_depIdxs = []int32{
	2, // 0: FindNftsInIdsRes.nfts:type_name -> Nft
	0, // 1: NftGrpcService.FindNftsInIds:input_type -> FindNftsInIdsReq
	1, // 2: NftGrpcService.FindNftsInIds:output_type -> FindNftsInIdsRes
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_modules_nft_nftPb_nftPb_proto_init() }
func file_modules_nft_nftPb_nftPb_proto_init() {
	if File_modules_nft_nftPb_nftPb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_modules_nft_nftPb_nftPb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNftsInIdsReq); i {
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
		file_modules_nft_nftPb_nftPb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindNftsInIdsRes); i {
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
		file_modules_nft_nftPb_nftPb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nft); i {
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
			RawDescriptor: file_modules_nft_nftPb_nftPb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_modules_nft_nftPb_nftPb_proto_goTypes,
		DependencyIndexes: file_modules_nft_nftPb_nftPb_proto_depIdxs,
		MessageInfos:      file_modules_nft_nftPb_nftPb_proto_msgTypes,
	}.Build()
	File_modules_nft_nftPb_nftPb_proto = out.File
	file_modules_nft_nftPb_nftPb_proto_rawDesc = nil
	file_modules_nft_nftPb_nftPb_proto_goTypes = nil
	file_modules_nft_nftPb_nftPb_proto_depIdxs = nil
}
