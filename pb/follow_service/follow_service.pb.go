// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.14.0
// source: follow_service.proto

package follow_service

import (
	follow_system "./follow_system"
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

type FollowReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Opt       int32 `protobuf:"varint,1,opt,name=opt,proto3" json:"opt,omitempty"`
	Uid       int64 `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Fans      int64 `protobuf:"varint,3,opt,name=fans,proto3" json:"fans,omitempty"`
	EventType int32 `protobuf:"varint,4,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
}

func (x *FollowReq) Reset() {
	*x = FollowReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowReq) ProtoMessage() {}

func (x *FollowReq) ProtoReflect() protoreflect.Message {
	mi := &file_follow_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowReq.ProtoReflect.Descriptor instead.
func (*FollowReq) Descriptor() ([]byte, []int) {
	return file_follow_service_proto_rawDescGZIP(), []int{0}
}

func (x *FollowReq) GetOpt() int32 {
	if x != nil {
		return x.Opt
	}
	return 0
}

func (x *FollowReq) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *FollowReq) GetFans() int64 {
	if x != nil {
		return x.Fans
	}
	return 0
}

func (x *FollowReq) GetEventType() int32 {
	if x != nil {
		return x.EventType
	}
	return 0
}

type FollowRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err       *follow_system.ErrorInfo `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	Opt       int32                    `protobuf:"varint,2,opt,name=opt,proto3" json:"opt,omitempty"`
	Uid       int64                    `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Fans      int64                    `protobuf:"varint,4,opt,name=fans,proto3" json:"fans,omitempty"`
	EventType int32                    `protobuf:"varint,5,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
}

func (x *FollowRsp) Reset() {
	*x = FollowRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_follow_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FollowRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FollowRsp) ProtoMessage() {}

func (x *FollowRsp) ProtoReflect() protoreflect.Message {
	mi := &file_follow_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FollowRsp.ProtoReflect.Descriptor instead.
func (*FollowRsp) Descriptor() ([]byte, []int) {
	return file_follow_service_proto_rawDescGZIP(), []int{1}
}

func (x *FollowRsp) GetErr() *follow_system.ErrorInfo {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *FollowRsp) GetOpt() int32 {
	if x != nil {
		return x.Opt
	}
	return 0
}

func (x *FollowRsp) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *FollowRsp) GetFans() int64 {
	if x != nil {
		return x.Fans
	}
	return 0
}

func (x *FollowRsp) GetEventType() int32 {
	if x != nil {
		return x.EventType
	}
	return 0
}

var File_follow_service_proto protoreflect.FileDescriptor

var file_follow_service_proto_rawDesc = []byte{
	0x0a, 0x14, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x09, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65,
	0x71, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x70, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6f, 0x70, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x61, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x61, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x87, 0x01, 0x0a, 0x09, 0x46, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x52, 0x73, 0x70, 0x12, 0x23, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6f,
	0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6f, 0x70, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x61, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66,
	0x61, 0x6e, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x32, 0x4e, 0x0a, 0x0e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x19, 0x2e, 0x66,
	0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f,
	0x6c, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x52,
	0x73, 0x70, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_follow_service_proto_rawDescOnce sync.Once
	file_follow_service_proto_rawDescData = file_follow_service_proto_rawDesc
)

func file_follow_service_proto_rawDescGZIP() []byte {
	file_follow_service_proto_rawDescOnce.Do(func() {
		file_follow_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_follow_service_proto_rawDescData)
	})
	return file_follow_service_proto_rawDescData
}

var file_follow_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_follow_service_proto_goTypes = []interface{}{
	(*FollowReq)(nil),               // 0: follow_service.FollowReq
	(*FollowRsp)(nil),               // 1: follow_service.FollowRsp
	(*follow_system.ErrorInfo)(nil), // 2: common.ErrorInfo
}
var file_follow_service_proto_depIdxs = []int32{
	2, // 0: follow_service.FollowRsp.err:type_name -> common.ErrorInfo
	0, // 1: follow_service.Follow_service.Ping:input_type -> follow_service.FollowReq
	1, // 2: follow_service.Follow_service.Ping:output_type -> follow_service.FollowRsp
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_follow_service_proto_init() }
func file_follow_service_proto_init() {
	if File_follow_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_follow_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowReq); i {
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
		file_follow_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FollowRsp); i {
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
			RawDescriptor: file_follow_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_follow_service_proto_goTypes,
		DependencyIndexes: file_follow_service_proto_depIdxs,
		MessageInfos:      file_follow_service_proto_msgTypes,
	}.Build()
	File_follow_service_proto = out.File
	file_follow_service_proto_rawDesc = nil
	file_follow_service_proto_goTypes = nil
	file_follow_service_proto_depIdxs = nil
}