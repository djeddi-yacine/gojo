// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.1
// source: v1/aspb/rpc_get_optional_full_anime_episode.proto

package aspbv1

import (
	nfpb "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
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

type GetOptionalFullAnimeEpisodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EpisodeID  int64 `protobuf:"varint,1,opt,name=episodeID,proto3" json:"episodeID,omitempty"`
	LanguageID int32 `protobuf:"varint,2,opt,name=languageID,proto3" json:"languageID,omitempty"`
	WithServer *bool `protobuf:"varint,3,opt,name=withServer,proto3,oneof" json:"withServer,omitempty"`
	WithSub    *bool `protobuf:"varint,4,opt,name=withSub,proto3,oneof" json:"withSub,omitempty"`
	WithDub    *bool `protobuf:"varint,5,opt,name=withDub,proto3,oneof" json:"withDub,omitempty"`
}

func (x *GetOptionalFullAnimeEpisodeRequest) Reset() {
	*x = GetOptionalFullAnimeEpisodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_get_optional_full_anime_episode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOptionalFullAnimeEpisodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOptionalFullAnimeEpisodeRequest) ProtoMessage() {}

func (x *GetOptionalFullAnimeEpisodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_get_optional_full_anime_episode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOptionalFullAnimeEpisodeRequest.ProtoReflect.Descriptor instead.
func (*GetOptionalFullAnimeEpisodeRequest) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDescGZIP(), []int{0}
}

func (x *GetOptionalFullAnimeEpisodeRequest) GetEpisodeID() int64 {
	if x != nil {
		return x.EpisodeID
	}
	return 0
}

func (x *GetOptionalFullAnimeEpisodeRequest) GetLanguageID() int32 {
	if x != nil {
		return x.LanguageID
	}
	return 0
}

func (x *GetOptionalFullAnimeEpisodeRequest) GetWithServer() bool {
	if x != nil && x.WithServer != nil {
		return *x.WithServer
	}
	return false
}

func (x *GetOptionalFullAnimeEpisodeRequest) GetWithSub() bool {
	if x != nil && x.WithSub != nil {
		return *x.WithSub
	}
	return false
}

func (x *GetOptionalFullAnimeEpisodeRequest) GetWithDub() bool {
	if x != nil && x.WithDub != nil {
		return *x.WithDub
	}
	return false
}

type GetOptionalFullAnimeEpisodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeEpisode *AnimeEpisodeResponse        `protobuf:"bytes,1,opt,name=animeEpisode,proto3" json:"animeEpisode,omitempty"`
	EpisodeMeta  *nfpb.AnimeMetaResponse      `protobuf:"bytes,2,opt,name=episodeMeta,proto3" json:"episodeMeta,omitempty"`
	ServerID     *int64                       `protobuf:"varint,3,opt,name=serverID,proto3,oneof" json:"serverID,omitempty"`
	Sub          *AnimeEpisodeSubDataResponse `protobuf:"bytes,4,opt,name=sub,proto3,oneof" json:"sub,omitempty"`
	Dub          *AnimeEpisodeDubDataResponse `protobuf:"bytes,5,opt,name=dub,proto3,oneof" json:"dub,omitempty"`
}

func (x *GetOptionalFullAnimeEpisodeResponse) Reset() {
	*x = GetOptionalFullAnimeEpisodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_aspb_rpc_get_optional_full_anime_episode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOptionalFullAnimeEpisodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOptionalFullAnimeEpisodeResponse) ProtoMessage() {}

func (x *GetOptionalFullAnimeEpisodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_aspb_rpc_get_optional_full_anime_episode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOptionalFullAnimeEpisodeResponse.ProtoReflect.Descriptor instead.
func (*GetOptionalFullAnimeEpisodeResponse) Descriptor() ([]byte, []int) {
	return file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDescGZIP(), []int{1}
}

func (x *GetOptionalFullAnimeEpisodeResponse) GetAnimeEpisode() *AnimeEpisodeResponse {
	if x != nil {
		return x.AnimeEpisode
	}
	return nil
}

func (x *GetOptionalFullAnimeEpisodeResponse) GetEpisodeMeta() *nfpb.AnimeMetaResponse {
	if x != nil {
		return x.EpisodeMeta
	}
	return nil
}

func (x *GetOptionalFullAnimeEpisodeResponse) GetServerID() int64 {
	if x != nil && x.ServerID != nil {
		return *x.ServerID
	}
	return 0
}

func (x *GetOptionalFullAnimeEpisodeResponse) GetSub() *AnimeEpisodeSubDataResponse {
	if x != nil {
		return x.Sub
	}
	return nil
}

func (x *GetOptionalFullAnimeEpisodeResponse) GetDub() *AnimeEpisodeDubDataResponse {
	if x != nil {
		return x.Dub
	}
	return nil
}

var File_v1_aspb_rpc_get_optional_full_anime_episode_proto protoreflect.FileDescriptor

var file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDesc = []byte{
	0x0a, 0x31, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x5f,
	0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x1a, 0x1f,
	0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x5f, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x16, 0x76, 0x31, 0x2f, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62,
	0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xec, 0x01, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61,
	0x6c, 0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x70, 0x69, 0x73,
	0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x70, 0x69,
	0x73, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x0a, 0x77, 0x69, 0x74, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0a, 0x77, 0x69,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x77,
	0x69, 0x74, 0x68, 0x53, 0x75, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x07,
	0x77, 0x69, 0x74, 0x68, 0x53, 0x75, 0x62, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07, 0x77, 0x69,
	0x74, 0x68, 0x44, 0x75, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x02, 0x52, 0x07, 0x77,
	0x69, 0x74, 0x68, 0x44, 0x75, 0x62, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x77, 0x69,
	0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x77, 0x69, 0x74,
	0x68, 0x53, 0x75, 0x62, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x44, 0x75, 0x62,
	0x22, 0xe6, 0x02, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x46, 0x75, 0x6c, 0x6c, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x0c, 0x61, 0x6e, 0x69, 0x6d,
	0x65, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x76, 0x31, 0x2e, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65,
	0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52,
	0x0c, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x3e, 0x0a,
	0x0b, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x31, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41,
	0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0b, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x1f, 0x0a,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x00, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x88, 0x01, 0x01, 0x12, 0x3d,
	0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x31,
	0x2e, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x45, 0x70, 0x69,
	0x73, 0x6f, 0x64, 0x65, 0x53, 0x75, 0x62, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x01, 0x52, 0x03, 0x73, 0x75, 0x62, 0x88, 0x01, 0x01, 0x12, 0x3d, 0x0a,
	0x03, 0x64, 0x75, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x73, 0x70, 0x62, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x45, 0x70, 0x69, 0x73,
	0x6f, 0x64, 0x65, 0x44, 0x75, 0x62, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x02, 0x52, 0x03, 0x64, 0x75, 0x62, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x44, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x73, 0x75,
	0x62, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x64, 0x75, 0x62, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e,
	0x65, 0x2d, 0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70,
	0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x70, 0x62, 0x3b, 0x61, 0x73, 0x70, 0x62, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDescOnce sync.Once
	file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDescData = file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDesc
)

func file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDescGZIP() []byte {
	file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDescOnce.Do(func() {
		file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDescData)
	})
	return file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDescData
}

var file_v1_aspb_rpc_get_optional_full_anime_episode_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_aspb_rpc_get_optional_full_anime_episode_proto_goTypes = []interface{}{
	(*GetOptionalFullAnimeEpisodeRequest)(nil),  // 0: v1.aspbv1.GetOptionalFullAnimeEpisodeRequest
	(*GetOptionalFullAnimeEpisodeResponse)(nil), // 1: v1.aspbv1.GetOptionalFullAnimeEpisodeResponse
	(*AnimeEpisodeResponse)(nil),                // 2: v1.aspbv1.AnimeEpisodeResponse
	(*nfpb.AnimeMetaResponse)(nil),              // 3: v1.nfpbv1.AnimeMetaResponse
	(*AnimeEpisodeSubDataResponse)(nil),         // 4: v1.aspbv1.AnimeEpisodeSubDataResponse
	(*AnimeEpisodeDubDataResponse)(nil),         // 5: v1.aspbv1.AnimeEpisodeDubDataResponse
}
var file_v1_aspb_rpc_get_optional_full_anime_episode_proto_depIdxs = []int32{
	2, // 0: v1.aspbv1.GetOptionalFullAnimeEpisodeResponse.animeEpisode:type_name -> v1.aspbv1.AnimeEpisodeResponse
	3, // 1: v1.aspbv1.GetOptionalFullAnimeEpisodeResponse.episodeMeta:type_name -> v1.nfpbv1.AnimeMetaResponse
	4, // 2: v1.aspbv1.GetOptionalFullAnimeEpisodeResponse.sub:type_name -> v1.aspbv1.AnimeEpisodeSubDataResponse
	5, // 3: v1.aspbv1.GetOptionalFullAnimeEpisodeResponse.dub:type_name -> v1.aspbv1.AnimeEpisodeDubDataResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_v1_aspb_rpc_get_optional_full_anime_episode_proto_init() }
func file_v1_aspb_rpc_get_optional_full_anime_episode_proto_init() {
	if File_v1_aspb_rpc_get_optional_full_anime_episode_proto != nil {
		return
	}
	file_v1_aspb_msg_anime_episode_proto_init()
	file_v1_aspb_msg_server_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_v1_aspb_rpc_get_optional_full_anime_episode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOptionalFullAnimeEpisodeRequest); i {
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
		file_v1_aspb_rpc_get_optional_full_anime_episode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOptionalFullAnimeEpisodeResponse); i {
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
	file_v1_aspb_rpc_get_optional_full_anime_episode_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_v1_aspb_rpc_get_optional_full_anime_episode_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1_aspb_rpc_get_optional_full_anime_episode_proto_goTypes,
		DependencyIndexes: file_v1_aspb_rpc_get_optional_full_anime_episode_proto_depIdxs,
		MessageInfos:      file_v1_aspb_rpc_get_optional_full_anime_episode_proto_msgTypes,
	}.Build()
	File_v1_aspb_rpc_get_optional_full_anime_episode_proto = out.File
	file_v1_aspb_rpc_get_optional_full_anime_episode_proto_rawDesc = nil
	file_v1_aspb_rpc_get_optional_full_anime_episode_proto_goTypes = nil
	file_v1_aspb_rpc_get_optional_full_anime_episode_proto_depIdxs = nil
}