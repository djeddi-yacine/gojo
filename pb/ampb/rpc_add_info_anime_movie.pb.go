// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: ampb/rpc_add_info_anime_movie.proto

package ampb

import (
	nfpb "github.com/dj-yacine-flutter/gojo/pb/nfpb"
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

type AddInfoAnimeMovieRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeID int64                     `protobuf:"varint,1,opt,name=animeID,proto3" json:"animeID,omitempty"`
	Genres  *nfpb.AnimeGenresRequest  `protobuf:"bytes,2,opt,name=genres,proto3" json:"genres,omitempty"`
	Studios *nfpb.AnimeStudiosRequest `protobuf:"bytes,3,opt,name=studios,proto3" json:"studios,omitempty"`
}

func (x *AddInfoAnimeMovieRequest) Reset() {
	*x = AddInfoAnimeMovieRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ampb_rpc_add_info_anime_movie_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInfoAnimeMovieRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInfoAnimeMovieRequest) ProtoMessage() {}

func (x *AddInfoAnimeMovieRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ampb_rpc_add_info_anime_movie_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInfoAnimeMovieRequest.ProtoReflect.Descriptor instead.
func (*AddInfoAnimeMovieRequest) Descriptor() ([]byte, []int) {
	return file_ampb_rpc_add_info_anime_movie_proto_rawDescGZIP(), []int{0}
}

func (x *AddInfoAnimeMovieRequest) GetAnimeID() int64 {
	if x != nil {
		return x.AnimeID
	}
	return 0
}

func (x *AddInfoAnimeMovieRequest) GetGenres() *nfpb.AnimeGenresRequest {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *AddInfoAnimeMovieRequest) GetStudios() *nfpb.AnimeStudiosRequest {
	if x != nil {
		return x.Studios
	}
	return nil
}

type AddInfoAnimeMovieResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnimeMovie *AnimeMovieResponse        `protobuf:"bytes,1,opt,name=animeMovie,proto3" json:"animeMovie,omitempty"`
	Genres     *nfpb.AnimeGenresResponse  `protobuf:"bytes,2,opt,name=genres,proto3" json:"genres,omitempty"`
	Studios    *nfpb.AnimeStudiosResponse `protobuf:"bytes,3,opt,name=studios,proto3" json:"studios,omitempty"`
}

func (x *AddInfoAnimeMovieResponse) Reset() {
	*x = AddInfoAnimeMovieResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ampb_rpc_add_info_anime_movie_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInfoAnimeMovieResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInfoAnimeMovieResponse) ProtoMessage() {}

func (x *AddInfoAnimeMovieResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ampb_rpc_add_info_anime_movie_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInfoAnimeMovieResponse.ProtoReflect.Descriptor instead.
func (*AddInfoAnimeMovieResponse) Descriptor() ([]byte, []int) {
	return file_ampb_rpc_add_info_anime_movie_proto_rawDescGZIP(), []int{1}
}

func (x *AddInfoAnimeMovieResponse) GetAnimeMovie() *AnimeMovieResponse {
	if x != nil {
		return x.AnimeMovie
	}
	return nil
}

func (x *AddInfoAnimeMovieResponse) GetGenres() *nfpb.AnimeGenresResponse {
	if x != nil {
		return x.Genres
	}
	return nil
}

func (x *AddInfoAnimeMovieResponse) GetStudios() *nfpb.AnimeStudiosResponse {
	if x != nil {
		return x.Studios
	}
	return nil
}

var File_ampb_rpc_add_info_anime_movie_proto protoreflect.FileDescriptor

var file_ampb_rpc_add_info_anime_movie_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x6d, 0x70, 0x62, 0x2f, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x5f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x6d, 0x70, 0x62, 0x1a, 0x16, 0x61, 0x6d, 0x70,
	0x62, 0x2f, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x6e, 0x66, 0x70, 0x62, 0x2f, 0x73, 0x74, 0x75, 0x64,
	0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x01, 0x0a, 0x18, 0x41, 0x64, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x49, 0x44, 0x12,
	0x30, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x72,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65,
	0x73, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53,
	0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x73,
	0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x22, 0xbe, 0x01, 0x0a, 0x19, 0x41, 0x64, 0x64, 0x49, 0x6e,
	0x66, 0x6f, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76,
	0x69, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x6d, 0x70, 0x62, 0x2e,
	0x41, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0a, 0x61, 0x6e, 0x69, 0x6d, 0x65, 0x4d, 0x6f, 0x76, 0x69, 0x65, 0x12, 0x31,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6e, 0x66, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x47, 0x65, 0x6e, 0x72, 0x65,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x72, 0x65,
	0x73, 0x12, 0x34, 0x0a, 0x07, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x66, 0x70, 0x62, 0x2e, 0x41, 0x6e, 0x69, 0x6d, 0x65, 0x53,
	0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07,
	0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x73, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x6a, 0x2d, 0x79, 0x61, 0x63, 0x69, 0x6e, 0x65, 0x2d,
	0x66, 0x6c, 0x75, 0x74, 0x74, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6a, 0x6f, 0x2f, 0x70, 0x62, 0x2f,
	0x61, 0x6d, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ampb_rpc_add_info_anime_movie_proto_rawDescOnce sync.Once
	file_ampb_rpc_add_info_anime_movie_proto_rawDescData = file_ampb_rpc_add_info_anime_movie_proto_rawDesc
)

func file_ampb_rpc_add_info_anime_movie_proto_rawDescGZIP() []byte {
	file_ampb_rpc_add_info_anime_movie_proto_rawDescOnce.Do(func() {
		file_ampb_rpc_add_info_anime_movie_proto_rawDescData = protoimpl.X.CompressGZIP(file_ampb_rpc_add_info_anime_movie_proto_rawDescData)
	})
	return file_ampb_rpc_add_info_anime_movie_proto_rawDescData
}

var file_ampb_rpc_add_info_anime_movie_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ampb_rpc_add_info_anime_movie_proto_goTypes = []interface{}{
	(*AddInfoAnimeMovieRequest)(nil),  // 0: ampb.AddInfoAnimeMovieRequest
	(*AddInfoAnimeMovieResponse)(nil), // 1: ampb.AddInfoAnimeMovieResponse
	(*nfpb.AnimeGenresRequest)(nil),   // 2: nfpb.AnimeGenresRequest
	(*nfpb.AnimeStudiosRequest)(nil),  // 3: nfpb.AnimeStudiosRequest
	(*AnimeMovieResponse)(nil),        // 4: ampb.AnimeMovieResponse
	(*nfpb.AnimeGenresResponse)(nil),  // 5: nfpb.AnimeGenresResponse
	(*nfpb.AnimeStudiosResponse)(nil), // 6: nfpb.AnimeStudiosResponse
}
var file_ampb_rpc_add_info_anime_movie_proto_depIdxs = []int32{
	2, // 0: ampb.AddInfoAnimeMovieRequest.genres:type_name -> nfpb.AnimeGenresRequest
	3, // 1: ampb.AddInfoAnimeMovieRequest.studios:type_name -> nfpb.AnimeStudiosRequest
	4, // 2: ampb.AddInfoAnimeMovieResponse.animeMovie:type_name -> ampb.AnimeMovieResponse
	5, // 3: ampb.AddInfoAnimeMovieResponse.genres:type_name -> nfpb.AnimeGenresResponse
	6, // 4: ampb.AddInfoAnimeMovieResponse.studios:type_name -> nfpb.AnimeStudiosResponse
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ampb_rpc_add_info_anime_movie_proto_init() }
func file_ampb_rpc_add_info_anime_movie_proto_init() {
	if File_ampb_rpc_add_info_anime_movie_proto != nil {
		return
	}
	file_ampb_anime_movie_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ampb_rpc_add_info_anime_movie_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInfoAnimeMovieRequest); i {
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
		file_ampb_rpc_add_info_anime_movie_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInfoAnimeMovieResponse); i {
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
			RawDescriptor: file_ampb_rpc_add_info_anime_movie_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ampb_rpc_add_info_anime_movie_proto_goTypes,
		DependencyIndexes: file_ampb_rpc_add_info_anime_movie_proto_depIdxs,
		MessageInfos:      file_ampb_rpc_add_info_anime_movie_proto_msgTypes,
	}.Build()
	File_ampb_rpc_add_info_anime_movie_proto = out.File
	file_ampb_rpc_add_info_anime_movie_proto_rawDesc = nil
	file_ampb_rpc_add_info_anime_movie_proto_goTypes = nil
	file_ampb_rpc_add_info_anime_movie_proto_depIdxs = nil
}