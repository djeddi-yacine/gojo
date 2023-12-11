// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: aspb/torrent.proto

package aspb

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

type AnimeSerieTorrentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LanguageID  int32  `protobuf:"varint,1,opt,name=languageID,proto3" json:"languageID,omitempty"`
	FileName    string `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	TorrentHash string `protobuf:"bytes,3,opt,name=torrentHash,proto3" json:"torrentHash,omitempty"`
	TorrentFile string `protobuf:"bytes,4,opt,name=torrentFile,proto3" json:"torrentFile,omitempty"`
	Seeds       int32  `protobuf:"varint,5,opt,name=seeds,proto3" json:"seeds,omitempty"`
	Peers       int32  `protobuf:"varint,6,opt,name=peers,proto3" json:"peers,omitempty"`
	Leechers    int32  `protobuf:"varint,7,opt,name=leechers,proto3" json:"leechers,omitempty"`
	SizeBytes   int64  `protobuf:"varint,8,opt,name=sizeBytes,proto3" json:"sizeBytes,omitempty"`
	Quality     string `protobuf:"bytes,9,opt,name=quality,proto3" json:"quality,omitempty"`
}

func (x *AnimeSerieTorrentRequest) Reset() {
	*x = AnimeSerieTorrentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_torrent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeSerieTorrentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeSerieTorrentRequest) ProtoMessage() {}

func (x *AnimeSerieTorrentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_torrent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeSerieTorrentRequest.ProtoReflect.Descriptor instead.
func (*AnimeSerieTorrentRequest) Descriptor() ([]byte, []int) {
	return file_aspb_torrent_proto_rawDescGZIP(), []int{0}
}

func (x *AnimeSerieTorrentRequest) GetLanguageID() int32 {
	if x != nil {
		return x.LanguageID
	}
	return 0
}

func (x *AnimeSerieTorrentRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *AnimeSerieTorrentRequest) GetTorrentHash() string {
	if x != nil {
		return x.TorrentHash
	}
	return ""
}

func (x *AnimeSerieTorrentRequest) GetTorrentFile() string {
	if x != nil {
		return x.TorrentFile
	}
	return ""
}

func (x *AnimeSerieTorrentRequest) GetSeeds() int32 {
	if x != nil {
		return x.Seeds
	}
	return 0
}

func (x *AnimeSerieTorrentRequest) GetPeers() int32 {
	if x != nil {
		return x.Peers
	}
	return 0
}

func (x *AnimeSerieTorrentRequest) GetLeechers() int32 {
	if x != nil {
		return x.Leechers
	}
	return 0
}

func (x *AnimeSerieTorrentRequest) GetSizeBytes() int64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

func (x *AnimeSerieTorrentRequest) GetQuality() string {
	if x != nil {
		return x.Quality
	}
	return ""
}

type AnimeSerieTorrentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          int64                  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	LanguageID  int32                  `protobuf:"varint,2,opt,name=languageID,proto3" json:"languageID,omitempty"`
	FileName    string                 `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
	TorrentHash string                 `protobuf:"bytes,4,opt,name=torrentHash,proto3" json:"torrentHash,omitempty"`
	TorrentFile string                 `protobuf:"bytes,5,opt,name=torrentFile,proto3" json:"torrentFile,omitempty"`
	Seeds       int32                  `protobuf:"varint,6,opt,name=seeds,proto3" json:"seeds,omitempty"`
	Peers       int32                  `protobuf:"varint,7,opt,name=peers,proto3" json:"peers,omitempty"`
	Leechers    int32                  `protobuf:"varint,8,opt,name=leechers,proto3" json:"leechers,omitempty"`
	SizeBytes   int64                  `protobuf:"varint,9,opt,name=sizeBytes,proto3" json:"sizeBytes,omitempty"`
	Quality     string                 `protobuf:"bytes,10,opt,name=quality,proto3" json:"quality,omitempty"`
	CreatedAt   *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *AnimeSerieTorrentResponse) Reset() {
	*x = AnimeSerieTorrentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aspb_torrent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnimeSerieTorrentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnimeSerieTorrentResponse) ProtoMessage() {}

func (x *AnimeSerieTorrentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_aspb_torrent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnimeSerieTorrentResponse.ProtoReflect.Descriptor instead.
func (*AnimeSerieTorrentResponse) Descriptor() ([]byte, []int) {
	return file_aspb_torrent_proto_rawDescGZIP(), []int{1}
}

func (x *AnimeSerieTorrentResponse) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AnimeSerieTorrentResponse) GetLanguageID() int32 {
	if x != nil {
		return x.LanguageID
	}
	return 0
}

func (x *AnimeSerieTorrentResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *AnimeSerieTorrentResponse) GetTorrentHash() string {
	if x != nil {
		return x.TorrentHash
	}
	return ""
}

func (x *AnimeSerieTorrentResponse) GetTorrentFile() string {
	if x != nil {
		return x.TorrentFile
	}
	return ""
}

func (x *AnimeSerieTorrentResponse) GetSeeds() int32 {
	if x != nil {
		return x.Seeds
	}
	return 0
}

func (x *AnimeSerieTorrentResponse) GetPeers() int32 {
	if x != nil {
		return x.Peers
	}
	return 0
}

func (x *AnimeSerieTorrentResponse) GetLeechers() int32 {
	if x != nil {
		return x.Leechers
	}
	return 0
}

func (x *AnimeSerieTorrentResponse) GetSizeBytes() int64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

func (x *AnimeSerieTorrentResponse) GetQuality() string {
	if x != nil {
		return x.Quality
	}
	return ""
}

func (x *AnimeSerieTorrentResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_aspb_torrent_proto protoreflect.FileDescriptor

var file_aspb_torrent_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x73, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x02, 0x0a, 0x18,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48,
	0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x65, 0x64,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x65, 0x65, 0x64, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x65, 0x63, 0x68, 0x65, 0x72, 0x73,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x65, 0x65, 0x63, 0x68, 0x65, 0x72, 0x73,
	0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x22, 0xe5, 0x02, 0x0a, 0x19, 0x41, 0x6e, 0x69,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x69, 0x65, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x6f, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x65, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x65, 0x65, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x65, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x65, 0x65,
	0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x65, 0x65, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x65, 0x65, 0x63, 0x68, 0x65, 0x72, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x71, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64,
	0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72,
	0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aspb_torrent_proto_rawDescOnce sync.Once
	file_aspb_torrent_proto_rawDescData = file_aspb_torrent_proto_rawDesc
)

func file_aspb_torrent_proto_rawDescGZIP() []byte {
	file_aspb_torrent_proto_rawDescOnce.Do(func() {
		file_aspb_torrent_proto_rawDescData = protoimpl.X.CompressGZIP(file_aspb_torrent_proto_rawDescData)
	})
	return file_aspb_torrent_proto_rawDescData
}

var file_aspb_torrent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aspb_torrent_proto_goTypes = []interface{}{
	(*AnimeSerieTorrentRequest)(nil),  // 0: aspb.AnimeSerieTorrentRequest
	(*AnimeSerieTorrentResponse)(nil), // 1: aspb.AnimeSerieTorrentResponse
	(*timestamppb.Timestamp)(nil),     // 2: google.protobuf.Timestamp
}
var file_aspb_torrent_proto_depIdxs = []int32{
	2, // 0: aspb.AnimeSerieTorrentResponse.createdAt:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aspb_torrent_proto_init() }
func file_aspb_torrent_proto_init() {
	if File_aspb_torrent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aspb_torrent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeSerieTorrentRequest); i {
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
		file_aspb_torrent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnimeSerieTorrentResponse); i {
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
			RawDescriptor: file_aspb_torrent_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aspb_torrent_proto_goTypes,
		DependencyIndexes: file_aspb_torrent_proto_depIdxs,
		MessageInfos:      file_aspb_torrent_proto_msgTypes,
	}.Build()
	File_aspb_torrent_proto = out.File
	file_aspb_torrent_proto_rawDesc = nil
	file_aspb_torrent_proto_goTypes = nil
	file_aspb_torrent_proto_depIdxs = nil
}