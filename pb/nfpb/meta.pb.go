// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: nfpb/meta.proto

package nfpb

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

type MetaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Overview string `protobuf:"bytes,2,opt,name=overview,proto3" json:"overview,omitempty"`
}

func (x *MetaRequest) Reset() {
	*x = MetaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nfpb_meta_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaRequest) ProtoMessage() {}

func (x *MetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nfpb_meta_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaRequest.ProtoReflect.Descriptor instead.
func (*MetaRequest) Descriptor() ([]byte, []int) {
	return file_nfpb_meta_proto_rawDescGZIP(), []int{0}
}

func (x *MetaRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MetaRequest) GetOverview() string {
	if x != nil {
		return x.Overview
	}
	return ""
}

type MetaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetaID   int64  `protobuf:"varint,1,opt,name=metaID,proto3" json:"metaID,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Overview string `protobuf:"bytes,3,opt,name=overview,proto3" json:"overview,omitempty"`
}

func (x *MetaResponse) Reset() {
	*x = MetaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nfpb_meta_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetaResponse) ProtoMessage() {}

func (x *MetaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nfpb_meta_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetaResponse.ProtoReflect.Descriptor instead.
func (*MetaResponse) Descriptor() ([]byte, []int) {
	return file_nfpb_meta_proto_rawDescGZIP(), []int{1}
}

func (x *MetaResponse) GetMetaID() int64 {
	if x != nil {
		return x.MetaID
	}
	return 0
}

func (x *MetaResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *MetaResponse) GetOverview() string {
	if x != nil {
		return x.Overview
	}
	return ""
}

type AnimeMetaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LanguageID int32        `protobuf:"varint,1,opt,name=languageID,proto3" json:"languageID,omitempty"`
	Meta       *MetaRequest `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *AnimeMetaRequest) Reset() {
	*x = AnimeMetaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nfpb_meta_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeMetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeMetaRequest) ProtoMessage() {}

func (x *AnimeMetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_nfpb_meta_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeMetaRequest.ProtoReflect.Descriptor instead.
func (*AnimeMetaRequest) Descriptor() ([]byte, []int) {
	return file_nfpb_meta_proto_rawDescGZIP(), []int{2}
}

func (x *AnimeMetaRequest) GetLanguageID() int32 {
	if x != nil {
		return x.LanguageID
	}
	return 0
}

func (x *AnimeMetaRequest) GetMeta() *MetaRequest {
	if x != nil {
		return x.Meta
	}
	return nil
}

type AnimeMetaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language  *LanguageResponse      `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	Meta      *MetaResponse          `protobuf:"bytes,2,opt,name=meta,proto3" json:"meta,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *AnimeMetaResponse) Reset() {
	*x = AnimeMetaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nfpb_meta_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeMetaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeMetaResponse) ProtoMessage() {}

func (x *AnimeMetaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_nfpb_meta_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeMetaResponse.ProtoReflect.Descriptor instead.
func (*AnimeMetaResponse) Descriptor() ([]byte, []int) {
	return file_nfpb_meta_proto_rawDescGZIP(), []int{3}
}

func (x *AnimeMetaResponse) GetLanguage() *LanguageResponse {
	if x != nil {
		return x.Language
	}
	return nil
}

func (x *AnimeMetaResponse) GetMeta() *MetaResponse {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *AnimeMetaResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_nfpb_meta_proto protoreflect.FileDescriptor

var file_nfpb_meta_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x6e, 0x66, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a,
	0x0b, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x22, 0x58,
	0x0a, 0x0c, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x65, 0x74, 0x61, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6d, 0x65, 0x74, 0x61, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6f, 0x76, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x22, 0x59, 0x0a, 0x10, 0x41, 0x6e, 0x69, 0x6d,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x25, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6e, 0x66, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x22, 0xa9, 0x01, 0x0a, 0x11, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6e, 0x66,
	0x70, 0x62, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x66,
	0x70, 0x62, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a,
	0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f,
	0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nfpb_meta_proto_rawDescOnce sync.Once
	file_nfpb_meta_proto_rawDescData = file_nfpb_meta_proto_rawDesc
)

func file_nfpb_meta_proto_rawDescGZIP() []byte {
	file_nfpb_meta_proto_rawDescOnce.Do(func() {
		file_nfpb_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_nfpb_meta_proto_rawDescData)
	})
	return file_nfpb_meta_proto_rawDescData
}

var file_nfpb_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_nfpb_meta_proto_goTypes = []interface{}{
	(*MetaRequest)(nil),           // 0: nfpb.MetaRequest
	(*MetaResponse)(nil),          // 1: nfpb.MetaResponse
	(*AnimeMetaRequest)(nil),      // 2: nfpb.AnimeMetaRequest
	(*AnimeMetaResponse)(nil),     // 3: nfpb.AnimeMetaResponse
	(*LanguageResponse)(nil),      // 4: nfpb.LanguageResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_nfpb_meta_proto_depIdxs = []int32{
	0, // 0: nfpb.AnimeMetaRequest.meta:type_name -> nfpb.MetaRequest
	4, // 1: nfpb.AnimeMetaResponse.language:type_name -> nfpb.LanguageResponse
	1, // 2: nfpb.AnimeMetaResponse.meta:type_name -> nfpb.MetaResponse
	5, // 3: nfpb.AnimeMetaResponse.createdAt:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_nfpb_meta_proto_init() }
func file_nfpb_meta_proto_init() {
	if File_nfpb_meta_proto != nil {
		return
	}
	file_nfpb_language_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_nfpb_meta_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaRequest); i {
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
		file_nfpb_meta_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetaResponse); i {
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
		file_nfpb_meta_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeMetaRequest); i {
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
		file_nfpb_meta_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeMetaResponse); i {
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
			RawDescriptor: file_nfpb_meta_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nfpb_meta_proto_goTypes,
		DependencyIndexes: file_nfpb_meta_proto_depIdxs,
		MessageInfos:      file_nfpb_meta_proto_msgTypes,
	}.Build()
	File_nfpb_meta_proto = out.File
	file_nfpb_meta_proto_rawDesc = nil
	file_nfpb_meta_proto_goTypes = nil
	file_nfpb_meta_proto_depIdxs = nil
}