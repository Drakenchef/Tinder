// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.12.4
// source: gmodels/gmodels.proto

package gmodels

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

	UID         string   `protobuf:"bytes,1,opt,name=UID,proto3" json:"UID,omitempty"`
	Login       string   `protobuf:"bytes,2,opt,name=Login,proto3" json:"Login,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
	Images      []*Image `protobuf:"bytes,4,rep,name=Images,proto3" json:"Images,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gmodels_gmodels_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_gmodels_gmodels_proto_msgTypes[0]
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
	return file_gmodels_gmodels_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetUID() string {
	if x != nil {
		return x.UID
	}
	return ""
}

func (x *User) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *User) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *User) GetImages() []*Image {
	if x != nil {
		return x.Images
	}
	return nil
}

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID     int32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	URL    string `protobuf:"bytes,2,opt,name=URL,proto3" json:"URL,omitempty"`
	UserID string `protobuf:"bytes,3,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gmodels_gmodels_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_gmodels_gmodels_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_gmodels_gmodels_proto_rawDescGZIP(), []int{1}
}

func (x *Image) GetID() int32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Image) GetURL() string {
	if x != nil {
		return x.URL
	}
	return ""
}

func (x *Image) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login       string `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Image       string `protobuf:"bytes,2,opt,name=Image,proto3" json:"Image,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gmodels_gmodels_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_gmodels_gmodels_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_gmodels_gmodels_proto_rawDescGZIP(), []int{2}
}

func (x *Profile) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Profile) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Profile) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type SignInInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=Login,proto3" json:"Login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *SignInInput) Reset() {
	*x = SignInInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gmodels_gmodels_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignInInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignInInput) ProtoMessage() {}

func (x *SignInInput) ProtoReflect() protoreflect.Message {
	mi := &file_gmodels_gmodels_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignInInput.ProtoReflect.Descriptor instead.
func (*SignInInput) Descriptor() ([]byte, []int) {
	return file_gmodels_gmodels_proto_rawDescGZIP(), []int{3}
}

func (x *SignInInput) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *SignInInput) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type ChangePassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Password string `protobuf:"bytes,1,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *ChangePassword) Reset() {
	*x = ChangePassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gmodels_gmodels_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangePassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangePassword) ProtoMessage() {}

func (x *ChangePassword) ProtoReflect() protoreflect.Message {
	mi := &file_gmodels_gmodels_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangePassword.ProtoReflect.Descriptor instead.
func (*ChangePassword) Descriptor() ([]byte, []int) {
	return file_gmodels_gmodels_proto_rawDescGZIP(), []int{4}
}

func (x *ChangePassword) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type Likes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UIDFirstLike  string               `protobuf:"bytes,1,opt,name=UIDFirstLike,proto3" json:"UIDFirstLike,omitempty"`
	UIDSecondLike string               `protobuf:"bytes,2,opt,name=UIDSecondLike,proto3" json:"UIDSecondLike,omitempty"`
	Date          *timestamp.Timestamp `protobuf:"bytes,3,opt,name=Date,proto3" json:"Date,omitempty"`
	Mutual        bool                 `protobuf:"varint,4,opt,name=Mutual,proto3" json:"Mutual,omitempty"`
}

func (x *Likes) Reset() {
	*x = Likes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gmodels_gmodels_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Likes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Likes) ProtoMessage() {}

func (x *Likes) ProtoReflect() protoreflect.Message {
	mi := &file_gmodels_gmodels_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Likes.ProtoReflect.Descriptor instead.
func (*Likes) Descriptor() ([]byte, []int) {
	return file_gmodels_gmodels_proto_rawDescGZIP(), []int{5}
}

func (x *Likes) GetUIDFirstLike() string {
	if x != nil {
		return x.UIDFirstLike
	}
	return ""
}

func (x *Likes) GetUIDSecondLike() string {
	if x != nil {
		return x.UIDSecondLike
	}
	return ""
}

func (x *Likes) GetDate() *timestamp.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Likes) GetMutual() bool {
	if x != nil {
		return x.Mutual
	}
	return false
}

type LikesUID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UIDFirstLike  string `protobuf:"bytes,1,opt,name=UIDFirstLike,proto3" json:"UIDFirstLike,omitempty"`
	UIDSecondLike string `protobuf:"bytes,2,opt,name=UIDSecondLike,proto3" json:"UIDSecondLike,omitempty"`
}

func (x *LikesUID) Reset() {
	*x = LikesUID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gmodels_gmodels_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikesUID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikesUID) ProtoMessage() {}

func (x *LikesUID) ProtoReflect() protoreflect.Message {
	mi := &file_gmodels_gmodels_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikesUID.ProtoReflect.Descriptor instead.
func (*LikesUID) Descriptor() ([]byte, []int) {
	return file_gmodels_gmodels_proto_rawDescGZIP(), []int{6}
}

func (x *LikesUID) GetUIDFirstLike() string {
	if x != nil {
		return x.UIDFirstLike
	}
	return ""
}

func (x *LikesUID) GetUIDSecondLike() string {
	if x != nil {
		return x.UIDSecondLike
	}
	return ""
}

var File_gmodels_gmodels_proto protoreflect.FileDescriptor

var file_gmodels_gmodels_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x67, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x78, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x67, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x52, 0x06, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x22, 0x41, 0x0a, 0x05, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x55, 0x52, 0x4c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x55, 0x52, 0x4c, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x57,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3f, 0x0a, 0x0b, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2c, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x99, 0x01, 0x0a, 0x05, 0x4c, 0x69, 0x6b, 0x65, 0x73,
	0x12, 0x22, 0x0a, 0x0c, 0x55, 0x49, 0x44, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x55, 0x49, 0x44, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x4c, 0x69, 0x6b, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x55, 0x49, 0x44, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x4c, 0x69, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x55, 0x49, 0x44,
	0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x75,
	0x74, 0x75, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x4d, 0x75, 0x74, 0x75,
	0x61, 0x6c, 0x22, 0x54, 0x0a, 0x08, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x55, 0x49, 0x44, 0x12, 0x22,
	0x0a, 0x0c, 0x55, 0x49, 0x44, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x69, 0x6b, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x55, 0x49, 0x44, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4c, 0x69,
	0x6b, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x55, 0x49, 0x44, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4c,
	0x69, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x55, 0x49, 0x44, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x72, 0x61, 0x6b, 0x65, 0x6e, 0x63, 0x68, 0x65,
	0x66, 0x2f, 0x54, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x3b, 0x67, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gmodels_gmodels_proto_rawDescOnce sync.Once
	file_gmodels_gmodels_proto_rawDescData = file_gmodels_gmodels_proto_rawDesc
)

func file_gmodels_gmodels_proto_rawDescGZIP() []byte {
	file_gmodels_gmodels_proto_rawDescOnce.Do(func() {
		file_gmodels_gmodels_proto_rawDescData = protoimpl.X.CompressGZIP(file_gmodels_gmodels_proto_rawDescData)
	})
	return file_gmodels_gmodels_proto_rawDescData
}

var file_gmodels_gmodels_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_gmodels_gmodels_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: gmodels.User
	(*Image)(nil),               // 1: gmodels.Image
	(*Profile)(nil),             // 2: gmodels.Profile
	(*SignInInput)(nil),         // 3: gmodels.SignInInput
	(*ChangePassword)(nil),      // 4: gmodels.ChangePassword
	(*Likes)(nil),               // 5: gmodels.Likes
	(*LikesUID)(nil),            // 6: gmodels.LikesUID
	(*timestamp.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_gmodels_gmodels_proto_depIdxs = []int32{
	1, // 0: gmodels.User.Images:type_name -> gmodels.Image
	7, // 1: gmodels.Likes.Date:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gmodels_gmodels_proto_init() }
func file_gmodels_gmodels_proto_init() {
	if File_gmodels_gmodels_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gmodels_gmodels_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_gmodels_gmodels_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_gmodels_gmodels_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_gmodels_gmodels_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignInInput); i {
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
		file_gmodels_gmodels_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangePassword); i {
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
		file_gmodels_gmodels_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Likes); i {
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
		file_gmodels_gmodels_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikesUID); i {
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
			RawDescriptor: file_gmodels_gmodels_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gmodels_gmodels_proto_goTypes,
		DependencyIndexes: file_gmodels_gmodels_proto_depIdxs,
		MessageInfos:      file_gmodels_gmodels_proto_msgTypes,
	}.Build()
	File_gmodels_gmodels_proto = out.File
	file_gmodels_gmodels_proto_rawDesc = nil
	file_gmodels_gmodels_proto_goTypes = nil
	file_gmodels_gmodels_proto_depIdxs = nil
}
