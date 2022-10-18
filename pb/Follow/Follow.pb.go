// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: Follow.proto

package Follow

import (
	Common "follow_system/pb/Common"
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

type FollowRQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Opt       *int32 `protobuf:"varint,1,opt,name=opt" json:"opt,omitempty"` // 0-关注 1-取消
	Uid       *int64 `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	Fans      *int64 `protobuf:"varint,3,opt,name=fans" json:"fans,omitempty"`
	EventType *int32 `protobuf:"varint,4,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
}

func (x *FollowRQ) Reset() {
	*x = FollowRQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Follow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRQ) ProtoMessage() {}

func (x *FollowRQ) ProtoReflect() protoreflect.Message {
	mi := &file_Follow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRQ.ProtoReflect.Descriptor instead.
func (*FollowRQ) Descriptor() ([]byte, []int) {
	return file_Follow_proto_rawDescGZIP(), []int{0}
}

func (x *FollowRQ) GetOpt() int32 {
	if x != nil && x.Opt != nil {
		return *x.Opt
	}
	return 0
}

func (x *FollowRQ) GetUid() int64 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *FollowRQ) GetFans() int64 {
	if x != nil && x.Fans != nil {
		return *x.Fans
	}
	return 0
}

func (x *FollowRQ) GetEventType() int32 {
	if x != nil && x.EventType != nil {
		return *x.EventType
	}
	return 0
}

type FollowRS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err       *Common.ErrorInfo `protobuf:"bytes,1,opt,name=err" json:"err,omitempty"`
	Opt       *int32            `protobuf:"varint,2,opt,name=opt" json:"opt,omitempty"`
	Uid       *int64            `protobuf:"varint,3,opt,name=uid" json:"uid,omitempty"`
	Fans      *int64            `protobuf:"varint,4,opt,name=fans" json:"fans,omitempty"`
	EventType *int32            `protobuf:"varint,5,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
}

func (x *FollowRS) Reset() {
	*x = FollowRS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Follow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRS) ProtoMessage() {}

func (x *FollowRS) ProtoReflect() protoreflect.Message {
	mi := &file_Follow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRS.ProtoReflect.Descriptor instead.
func (*FollowRS) Descriptor() ([]byte, []int) {
	return file_Follow_proto_rawDescGZIP(), []int{1}
}

func (x *FollowRS) GetErr() *Common.ErrorInfo {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *FollowRS) GetOpt() int32 {
	if x != nil && x.Opt != nil {
		return *x.Opt
	}
	return 0
}

func (x *FollowRS) GetUid() int64 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *FollowRS) GetFans() int64 {
	if x != nil && x.Fans != nil {
		return *x.Fans
	}
	return 0
}

func (x *FollowRS) GetEventType() int32 {
	if x != nil && x.EventType != nil {
		return *x.EventType
	}
	return 0
}

type GetFollowListRQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType *int32 `protobuf:"varint,1,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	Uid       *int64 `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
}

func (x *GetFollowListRQ) Reset() {
	*x = GetFollowListRQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Follow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowListRQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowListRQ) ProtoMessage() {}

func (x *GetFollowListRQ) ProtoReflect() protoreflect.Message {
	mi := &file_Follow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowListRQ.ProtoReflect.Descriptor instead.
func (*GetFollowListRQ) Descriptor() ([]byte, []int) {
	return file_Follow_proto_rawDescGZIP(), []int{2}
}

func (x *GetFollowListRQ) GetEventType() int32 {
	if x != nil && x.EventType != nil {
		return *x.EventType
	}
	return 0
}

func (x *GetFollowListRQ) GetUid() int64 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

type GetFollowListRS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType  *int32            `protobuf:"varint,1,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	Uid        *int64            `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	FollowList []int64           `protobuf:"varint,3,rep,name=follow_list,json=followList" json:"follow_list,omitempty"`
	Err        *Common.ErrorInfo `protobuf:"bytes,4,opt,name=err" json:"err,omitempty"`
}

func (x *GetFollowListRS) Reset() {
	*x = GetFollowListRS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Follow_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFollowListRS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFollowListRS) ProtoMessage() {}

func (x *GetFollowListRS) ProtoReflect() protoreflect.Message {
	mi := &file_Follow_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFollowListRS.ProtoReflect.Descriptor instead.
func (*GetFollowListRS) Descriptor() ([]byte, []int) {
	return file_Follow_proto_rawDescGZIP(), []int{3}
}

func (x *GetFollowListRS) GetEventType() int32 {
	if x != nil && x.EventType != nil {
		return *x.EventType
	}
	return 0
}

func (x *GetFollowListRS) GetUid() int64 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *GetFollowListRS) GetFollowList() []int64 {
	if x != nil {
		return x.FollowList
	}
	return nil
}

func (x *GetFollowListRS) GetErr() *Common.ErrorInfo {
	if x != nil {
		return x.Err
	}
	return nil
}

type GetFansListRQ struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType *int32 `protobuf:"varint,1,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	Uid       *int64 `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
}

func (x *GetFansListRQ) Reset() {
	*x = GetFansListRQ{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Follow_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFansListRQ) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFansListRQ) ProtoMessage() {}

func (x *GetFansListRQ) ProtoReflect() protoreflect.Message {
	mi := &file_Follow_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFansListRQ.ProtoReflect.Descriptor instead.
func (*GetFansListRQ) Descriptor() ([]byte, []int) {
	return file_Follow_proto_rawDescGZIP(), []int{4}
}

func (x *GetFansListRQ) GetEventType() int32 {
	if x != nil && x.EventType != nil {
		return *x.EventType
	}
	return 0
}

func (x *GetFansListRQ) GetUid() int64 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

type GetFansListRS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType *int32            `protobuf:"varint,1,opt,name=event_type,json=eventType" json:"event_type,omitempty"`
	Uid       *int64            `protobuf:"varint,2,opt,name=uid" json:"uid,omitempty"`
	FansList  []int64           `protobuf:"varint,3,rep,name=fans_list,json=fansList" json:"fans_list,omitempty"`
	Err       *Common.ErrorInfo `protobuf:"bytes,4,opt,name=err" json:"err,omitempty"`
}

func (x *GetFansListRS) Reset() {
	*x = GetFansListRS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Follow_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFansListRS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFansListRS) ProtoMessage() {}

func (x *GetFansListRS) ProtoReflect() protoreflect.Message {
	mi := &file_Follow_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFansListRS.ProtoReflect.Descriptor instead.
func (*GetFansListRS) Descriptor() ([]byte, []int) {
	return file_Follow_proto_rawDescGZIP(), []int{5}
}

func (x *GetFansListRS) GetEventType() int32 {
	if x != nil && x.EventType != nil {
		return *x.EventType
	}
	return 0
}

func (x *GetFansListRS) GetUid() int64 {
	if x != nil && x.Uid != nil {
		return *x.Uid
	}
	return 0
}

func (x *GetFansListRS) GetFansList() []int64 {
	if x != nil {
		return x.FansList
	}
	return nil
}

func (x *GetFansListRS) GetErr() *Common.ErrorInfo {
	if x != nil {
		return x.Err
	}
	return nil
}

var File_Follow_proto protoreflect.FileDescriptor

var file_Follow_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x1a, 0x0c, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x61, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x51, 0x12, 0x10,
	0x0a, 0x03, 0x6f, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6f, 0x70, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x66, 0x61, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x08, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x52, 0x53, 0x12, 0x23, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x70, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6f, 0x70, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x61, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x61, 0x6e, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x42,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x51, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x22, 0x88, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x53, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x40, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x51, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x82, 0x01, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x53, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x61, 0x6e, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x08, 0x66, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x23, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x03, 0x65, 0x72, 0x72, 0x32, 0x86, 0x02, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x12, 0x10, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f,
	0x77, 0x52, 0x51, 0x1a, 0x10, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x53, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x22, 0x07, 0x2f,
	0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x3a, 0x01, 0x2a, 0x12, 0x5c, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x17, 0x2e, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x51, 0x1a, 0x17, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x53, 0x22, 0x19, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x67, 0x65, 0x74, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x4c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x54, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x61,
	0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x15, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x51, 0x1a, 0x15, 0x2e,
	0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x53, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x22, 0x0c, 0x2f, 0x67,
	0x65, 0x74, 0x46, 0x61, 0x6e, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x0a, 0x5a,
	0x08, 0x2e, 0x2f, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
}

var (
	file_Follow_proto_rawDescOnce sync.Once
	file_Follow_proto_rawDescData = file_Follow_proto_rawDesc
)

func file_Follow_proto_rawDescGZIP() []byte {
	file_Follow_proto_rawDescOnce.Do(func() {
		file_Follow_proto_rawDescData = protoimpl.X.CompressGZIP(file_Follow_proto_rawDescData)
	})
	return file_Follow_proto_rawDescData
}

var file_Follow_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Follow_proto_goTypes = []interface{}{
	(*FollowRQ)(nil),         // 0: Follow.FollowRQ
	(*FollowRS)(nil),         // 1: Follow.FollowRS
	(*GetFollowListRQ)(nil),  // 2: Follow.GetFollowListRQ
	(*GetFollowListRS)(nil),  // 3: Follow.GetFollowListRS
	(*GetFansListRQ)(nil),    // 4: Follow.GetFansListRQ
	(*GetFansListRS)(nil),    // 5: Follow.GetFansListRS
	(*Common.ErrorInfo)(nil), // 6: Common.ErrorInfo
}
var file_Follow_proto_depIdxs = []int32{
	6, // 0: Follow.FollowRS.err:type_name -> Common.ErrorInfo
	6, // 1: Follow.GetFollowListRS.err:type_name -> Common.ErrorInfo
	6, // 2: Follow.GetFansListRS.err:type_name -> Common.ErrorInfo
	0, // 3: Follow.Follow_service.Follow:input_type -> Follow.FollowRQ
	2, // 4: Follow.Follow_service.GetFollowList:input_type -> Follow.GetFollowListRQ
	4, // 5: Follow.Follow_service.GetFansList:input_type -> Follow.GetFansListRQ
	1, // 6: Follow.Follow_service.Follow:output_type -> Follow.FollowRS
	3, // 7: Follow.Follow_service.GetFollowList:output_type -> Follow.GetFollowListRS
	5, // 8: Follow.Follow_service.GetFansList:output_type -> Follow.GetFansListRS
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_Follow_proto_init() }
func file_Follow_proto_init() {
	if File_Follow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Follow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRQ); i {
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
		file_Follow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRS); i {
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
		file_Follow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowListRQ); i {
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
		file_Follow_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFollowListRS); i {
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
		file_Follow_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFansListRQ); i {
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
		file_Follow_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFansListRS); i {
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
			RawDescriptor: file_Follow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Follow_proto_goTypes,
		DependencyIndexes: file_Follow_proto_depIdxs,
		MessageInfos:      file_Follow_proto_msgTypes,
	}.Build()
	File_Follow_proto = out.File
	file_Follow_proto_rawDesc = nil
	file_Follow_proto_goTypes = nil
	file_Follow_proto_depIdxs = nil
}