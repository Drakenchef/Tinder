// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: likes.proto

package gen

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

type LikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UIDFirstLike  string `protobuf:"bytes,1,opt,name=UIDFirstLike,proto3" json:"UIDFirstLike,omitempty"`
	UIDSecondLike string `protobuf:"bytes,2,opt,name=UIDSecondLike,proto3" json:"UIDSecondLike,omitempty"`
}

func (x *LikeRequest) Reset() {
	*x = LikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_likes_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeRequest) ProtoMessage() {}

func (x *LikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_likes_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeRequest.ProtoReflect.Descriptor instead.
func (*LikeRequest) Descriptor() ([]byte, []int) {
	return file_likes_proto_rawDescGZIP(), []int{0}
}

func (x *LikeRequest) GetUIDFirstLike() string {
	if x != nil {
		return x.UIDFirstLike
	}
	return ""
}

func (x *LikeRequest) GetUIDSecondLike() string {
	if x != nil {
		return x.UIDSecondLike
	}
	return ""
}

type LikeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
}

func (x *LikeResponse) Reset() {
	*x = LikeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_likes_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeResponse) ProtoMessage() {}

func (x *LikeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_likes_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeResponse.ProtoReflect.Descriptor instead.
func (*LikeResponse) Descriptor() ([]byte, []int) {
	return file_likes_proto_rawDescGZIP(), []int{1}
}

func (x *LikeResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type MutualLikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UIDFirstLike  string `protobuf:"bytes,1,opt,name=UIDFirstLike,proto3" json:"UIDFirstLike,omitempty"`
	UIDSecondLike string `protobuf:"bytes,2,opt,name=UIDSecondLike,proto3" json:"UIDSecondLike,omitempty"`
}

func (x *MutualLikeRequest) Reset() {
	*x = MutualLikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_likes_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutualLikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutualLikeRequest) ProtoMessage() {}

func (x *MutualLikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_likes_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutualLikeRequest.ProtoReflect.Descriptor instead.
func (*MutualLikeRequest) Descriptor() ([]byte, []int) {
	return file_likes_proto_rawDescGZIP(), []int{2}
}

func (x *MutualLikeRequest) GetUIDFirstLike() string {
	if x != nil {
		return x.UIDFirstLike
	}
	return ""
}

func (x *MutualLikeRequest) GetUIDSecondLike() string {
	if x != nil {
		return x.UIDSecondLike
	}
	return ""
}

type MutualLikeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MutuallyLiked bool `protobuf:"varint,1,opt,name=MutuallyLiked,proto3" json:"MutuallyLiked,omitempty"`
}

func (x *MutualLikeResponse) Reset() {
	*x = MutualLikeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_likes_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MutualLikeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MutualLikeResponse) ProtoMessage() {}

func (x *MutualLikeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_likes_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MutualLikeResponse.ProtoReflect.Descriptor instead.
func (*MutualLikeResponse) Descriptor() ([]byte, []int) {
	return file_likes_proto_rawDescGZIP(), []int{3}
}

func (x *MutualLikeResponse) GetMutuallyLiked() bool {
	if x != nil {
		return x.MutuallyLiked
	}
	return false
}

var File_likes_proto protoreflect.FileDescriptor

var file_likes_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6c,
	0x69, 0x6b, 0x65, 0x73, 0x22, 0x57, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x55, 0x49, 0x44, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4c,
	0x69, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x55, 0x49, 0x44, 0x46, 0x69,
	0x72, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x55, 0x49, 0x44, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x55, 0x49, 0x44, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x22, 0x28, 0x0a,
	0x0c, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x5d, 0x0a, 0x11, 0x4d, 0x75, 0x74, 0x75, 0x61,
	0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x55, 0x49, 0x44, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x55, 0x49, 0x44, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65,
	0x12, 0x24, 0x0a, 0x0d, 0x55, 0x49, 0x44, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4c, 0x69, 0x6b,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x55, 0x49, 0x44, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x22, 0x3a, 0x0a, 0x12, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c,
	0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x24, 0x0a, 0x0d,
	0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x6c, 0x79, 0x4c, 0x69, 0x6b,
	0x65, 0x64, 0x32, 0x87, 0x01, 0x0a, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x35, 0x0a, 0x08,
	0x4c, 0x69, 0x6b, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x73,
	0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x6c,
	0x69, 0x6b, 0x65, 0x73, 0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0e, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x69, 0x6b,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x2e, 0x4d, 0x75,
	0x74, 0x75, 0x61, 0x6c, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x2e, 0x4d, 0x75, 0x74, 0x75, 0x61, 0x6c, 0x4c, 0x69,
	0x6b, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b,
	0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c,
	0x69, 0x6b, 0x65, 0x73, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x3b, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_likes_proto_rawDescOnce sync.Once
	file_likes_proto_rawDescData = file_likes_proto_rawDesc
)

func file_likes_proto_rawDescGZIP() []byte {
	file_likes_proto_rawDescOnce.Do(func() {
		file_likes_proto_rawDescData = protoimpl.X.CompressGZIP(file_likes_proto_rawDescData)
	})
	return file_likes_proto_rawDescData
}

var file_likes_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_likes_proto_goTypes = []interface{}{
	(*LikeRequest)(nil),        // 0: likes.LikeRequest
	(*LikeResponse)(nil),       // 1: likes.LikeResponse
	(*MutualLikeRequest)(nil),  // 2: likes.MutualLikeRequest
	(*MutualLikeResponse)(nil), // 3: likes.MutualLikeResponse
}
var file_likes_proto_depIdxs = []int32{
	0, // 0: likes.Likes.LikeUser:input_type -> likes.LikeRequest
	2, // 1: likes.Likes.MutualLikeUser:input_type -> likes.MutualLikeRequest
	1, // 2: likes.Likes.LikeUser:output_type -> likes.LikeResponse
	3, // 3: likes.Likes.MutualLikeUser:output_type -> likes.MutualLikeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_likes_proto_init() }
func file_likes_proto_init() {
	if File_likes_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_likes_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeRequest); i {
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
		file_likes_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeResponse); i {
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
		file_likes_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutualLikeRequest); i {
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
		file_likes_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MutualLikeResponse); i {
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
			RawDescriptor: file_likes_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_likes_proto_goTypes,
		DependencyIndexes: file_likes_proto_depIdxs,
		MessageInfos:      file_likes_proto_msgTypes,
	}.Build()
	File_likes_proto = out.File
	file_likes_proto_rawDesc = nil
	file_likes_proto_goTypes = nil
	file_likes_proto_depIdxs = nil
}
