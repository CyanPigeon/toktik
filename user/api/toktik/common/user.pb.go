// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: api/toktik/common/user.proto

package common

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                       // 用户id
	Name            string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                                    // 用户名称
	FollowCount     *int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3,oneof" json:"follow_count,omitempty"`            // 关注总数
	FollowerCount   *int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3,oneof" json:"follower_count,omitempty"`      // 粉丝总数
	IsFollow        bool    `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`                           // true-已关注，false-未关注
	Avatar          *string `protobuf:"bytes,6,opt,name=avatar,proto3,oneof" json:"avatar,omitempty"`                                          // 用户头像
	BackgroundImage *string `protobuf:"bytes,7,opt,name=background_image,json=backgroundImage,proto3,oneof" json:"background_image,omitempty"` // 用户个人页顶部大图
	Signature       *string `protobuf:"bytes,8,opt,name=signature,proto3,oneof" json:"signature,omitempty"`                                    // 个人简介
	TotalFavorited  *int64  `protobuf:"varint,9,opt,name=total_favorited,json=totalFavorited,proto3,oneof" json:"total_favorited,omitempty"`   // 获赞数量
	WorkCount       *int64  `protobuf:"varint,10,opt,name=work_count,json=workCount,proto3,oneof" json:"work_count,omitempty"`                 // 作品数量
	FavoriteCount   *int64  `protobuf:"varint,11,opt,name=favorite_count,json=favoriteCount,proto3,oneof" json:"favorite_count,omitempty"`     // 点赞数量
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_toktik_common_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_api_toktik_common_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_api_toktik_common_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil && x.FollowCount != nil {
		return *x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil && x.FollowerCount != nil {
		return *x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *User) GetAvatar() string {
	if x != nil && x.Avatar != nil {
		return *x.Avatar
	}
	return ""
}

func (x *User) GetBackgroundImage() string {
	if x != nil && x.BackgroundImage != nil {
		return *x.BackgroundImage
	}
	return ""
}

func (x *User) GetSignature() string {
	if x != nil && x.Signature != nil {
		return *x.Signature
	}
	return ""
}

func (x *User) GetTotalFavorited() int64 {
	if x != nil && x.TotalFavorited != nil {
		return *x.TotalFavorited
	}
	return 0
}

func (x *User) GetWorkCount() int64 {
	if x != nil && x.WorkCount != nil {
		return *x.WorkCount
	}
	return 0
}

func (x *User) GetFavoriteCount() int64 {
	if x != nil && x.FavoriteCount != nil {
		return *x.FavoriteCount
	}
	return 0
}

var File_api_toktik_common_user_proto protoreflect.FileDescriptor

var file_api_toktik_common_user_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6b, 0x74, 0x69, 0x6b, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91,
	0x04, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0c, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x00, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x01, 0x52, 0x0d, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x1b, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x10, 0x62, 0x61, 0x63,
	0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e,
	0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2c, 0x0a, 0x0f,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x03, 0x48, 0x05, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61,
	0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x48, 0x06,
	0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x12, 0x2a,
	0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x48, 0x07, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x11, 0x0a, 0x0f, 0x5f,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x62, 0x61,
	0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x12, 0x0a, 0x10,
	0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x42,
	0x11, 0x0a, 0x0f, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x42, 0x1f, 0x5a, 0x1d, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x6f, 0x6b, 0x74, 0x69, 0x6b, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x3b, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_toktik_common_user_proto_rawDescOnce sync.Once
	file_api_toktik_common_user_proto_rawDescData = file_api_toktik_common_user_proto_rawDesc
)

func file_api_toktik_common_user_proto_rawDescGZIP() []byte {
	file_api_toktik_common_user_proto_rawDescOnce.Do(func() {
		file_api_toktik_common_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_toktik_common_user_proto_rawDescData)
	})
	return file_api_toktik_common_user_proto_rawDescData
}

var file_api_toktik_common_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_api_toktik_common_user_proto_goTypes = []interface{}{
	(*User)(nil), // 0: User
}
var file_api_toktik_common_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_toktik_common_user_proto_init() }
func file_api_toktik_common_user_proto_init() {
	if File_api_toktik_common_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_toktik_common_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
	file_api_toktik_common_user_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_toktik_common_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_toktik_common_user_proto_goTypes,
		DependencyIndexes: file_api_toktik_common_user_proto_depIdxs,
		MessageInfos:      file_api_toktik_common_user_proto_msgTypes,
	}.Build()
	File_api_toktik_common_user_proto = out.File
	file_api_toktik_common_user_proto_rawDesc = nil
	file_api_toktik_common_user_proto_goTypes = nil
	file_api_toktik_common_user_proto_depIdxs = nil
}
