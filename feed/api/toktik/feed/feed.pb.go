// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.19.1
// source: api/toktik/feed/feed.proto

package feed

import (
	common "feed/api/toktik/common"
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

type FeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestTime *int64  `protobuf:"varint,1,opt,name=latest_time,json=latestTime,proto3,oneof" json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `protobuf:"bytes,2,opt,name=token,proto3,oneof" json:"token,omitempty"`                              // 可选参数，登录用户设置
}

func (x *FeedRequest) Reset() {
	*x = FeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_toktik_feed_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedRequest) ProtoMessage() {}

func (x *FeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_toktik_feed_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedRequest.ProtoReflect.Descriptor instead.
func (*FeedRequest) Descriptor() ([]byte, []int) {
	return file_api_toktik_feed_feed_proto_rawDescGZIP(), []int{0}
}

func (x *FeedRequest) GetLatestTime() int64 {
	if x != nil && x.LatestTime != nil {
		return *x.LatestTime
	}
	return 0
}

func (x *FeedRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type FeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32           `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code"`   // 状态码，0-成功，其他值-失败
	StatusMsg  *string         `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg"` // 返回状态描述
	VideoList  []*common.Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list"`       // 视频列表
	NextTime   *int64          `protobuf:"varint,4,opt,name=next_time,json=nextTime,proto3,oneof" json:"next_time"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

func (x *FeedResponse) Reset() {
	*x = FeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_toktik_feed_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeedResponse) ProtoMessage() {}

func (x *FeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_toktik_feed_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeedResponse.ProtoReflect.Descriptor instead.
func (*FeedResponse) Descriptor() ([]byte, []int) {
	return file_api_toktik_feed_feed_proto_rawDescGZIP(), []int{1}
}

func (x *FeedResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *FeedResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *FeedResponse) GetVideoList() []*common.Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *FeedResponse) GetNextTime() int64 {
	if x != nil && x.NextTime != nil {
		return *x.NextTime
	}
	return 0
}

var File_api_toktik_feed_feed_proto protoreflect.FileDescriptor

var file_api_toktik_feed_feed_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6b, 0x74, 0x69, 0x6b, 0x2f, 0x66, 0x65, 0x65,
	0x64, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x61, 0x70, 0x69, 0x2f,
	0x74, 0x6f, 0x6b, 0x74, 0x69, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x0b, 0x46, 0x65, 0x65,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52,
	0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xb9, 0x01, 0x0a, 0x0c, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01, 0x12, 0x25, 0x0a, 0x0a, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x06, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73,
	0x67, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x32,
	0x44, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x3c, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x64, 0x53,
	0x72, 0x76, 0x12, 0x0c, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e,
	0x2f, 0x66, 0x65, 0x65, 0x64, 0x42, 0x1b, 0x5a, 0x19, 0x66, 0x65, 0x65, 0x64, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x74, 0x6f, 0x6b, 0x74, 0x69, 0x6b, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x3b, 0x66, 0x65,
	0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_toktik_feed_feed_proto_rawDescOnce sync.Once
	file_api_toktik_feed_feed_proto_rawDescData = file_api_toktik_feed_feed_proto_rawDesc
)

func file_api_toktik_feed_feed_proto_rawDescGZIP() []byte {
	file_api_toktik_feed_feed_proto_rawDescOnce.Do(func() {
		file_api_toktik_feed_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_toktik_feed_feed_proto_rawDescData)
	})
	return file_api_toktik_feed_feed_proto_rawDescData
}

var file_api_toktik_feed_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_toktik_feed_feed_proto_goTypes = []interface{}{
	(*FeedRequest)(nil),  // 0: FeedRequest
	(*FeedResponse)(nil), // 1: FeedResponse
	(*common.Video)(nil), // 2: Video
}
var file_api_toktik_feed_feed_proto_depIdxs = []int32{
	2, // 0: FeedResponse.video_list:type_name -> Video
	0, // 1: Feed.FeedSrv:input_type -> FeedRequest
	1, // 2: Feed.FeedSrv:output_type -> FeedResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_toktik_feed_feed_proto_init() }
func file_api_toktik_feed_feed_proto_init() {
	if File_api_toktik_feed_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_toktik_feed_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedRequest); i {
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
		file_api_toktik_feed_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeedResponse); i {
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
	file_api_toktik_feed_feed_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_api_toktik_feed_feed_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_toktik_feed_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_toktik_feed_feed_proto_goTypes,
		DependencyIndexes: file_api_toktik_feed_feed_proto_depIdxs,
		MessageInfos:      file_api_toktik_feed_feed_proto_msgTypes,
	}.Build()
	File_api_toktik_feed_feed_proto = out.File
	file_api_toktik_feed_feed_proto_rawDesc = nil
	file_api_toktik_feed_feed_proto_goTypes = nil
	file_api_toktik_feed_feed_proto_depIdxs = nil
}
