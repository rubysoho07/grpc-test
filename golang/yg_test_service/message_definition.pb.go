// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: message_definition.proto

package message_definition

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type YGRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReqMessage string `protobuf:"bytes,1,opt,name=req_message,json=reqMessage,proto3" json:"req_message,omitempty"`
}

func (x *YGRequest) Reset() {
	*x = YGRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_definition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YGRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YGRequest) ProtoMessage() {}

func (x *YGRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_definition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YGRequest.ProtoReflect.Descriptor instead.
func (*YGRequest) Descriptor() ([]byte, []int) {
	return file_message_definition_proto_rawDescGZIP(), []int{0}
}

func (x *YGRequest) GetReqMessage() string {
	if x != nil {
		return x.ReqMessage
	}
	return ""
}

type YGResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResMessage string `protobuf:"bytes,1,opt,name=res_message,json=resMessage,proto3" json:"res_message,omitempty"`
}

func (x *YGResponse) Reset() {
	*x = YGResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_definition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YGResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YGResponse) ProtoMessage() {}

func (x *YGResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_definition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YGResponse.ProtoReflect.Descriptor instead.
func (*YGResponse) Descriptor() ([]byte, []int) {
	return file_message_definition_proto_rawDescGZIP(), []int{1}
}

func (x *YGResponse) GetResMessage() string {
	if x != nil {
		return x.ResMessage
	}
	return ""
}

var File_message_definition_proto protoreflect.FileDescriptor

var file_message_definition_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a, 0x09, 0x59, 0x47,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x71, 0x5f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65,
	0x71, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2d, 0x0a, 0x0a, 0x59, 0x47, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x34, 0x0a, 0x0d, 0x59, 0x47, 0x54, 0x65, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x0a, 0x2e, 0x59, 0x47, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0b, 0x2e, 0x59, 0x47, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_definition_proto_rawDescOnce sync.Once
	file_message_definition_proto_rawDescData = file_message_definition_proto_rawDesc
)

func file_message_definition_proto_rawDescGZIP() []byte {
	file_message_definition_proto_rawDescOnce.Do(func() {
		file_message_definition_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_definition_proto_rawDescData)
	})
	return file_message_definition_proto_rawDescData
}

var file_message_definition_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_message_definition_proto_goTypes = []interface{}{
	(*YGRequest)(nil),  // 0: YGRequest
	(*YGResponse)(nil), // 1: YGResponse
}
var file_message_definition_proto_depIdxs = []int32{
	0, // 0: YGTestService.SayHello:input_type -> YGRequest
	1, // 1: YGTestService.SayHello:output_type -> YGResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_message_definition_proto_init() }
func file_message_definition_proto_init() {
	if File_message_definition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_definition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YGRequest); i {
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
		file_message_definition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*YGResponse); i {
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
			RawDescriptor: file_message_definition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_definition_proto_goTypes,
		DependencyIndexes: file_message_definition_proto_depIdxs,
		MessageInfos:      file_message_definition_proto_msgTypes,
	}.Build()
	File_message_definition_proto = out.File
	file_message_definition_proto_rawDesc = nil
	file_message_definition_proto_goTypes = nil
	file_message_definition_proto_depIdxs = nil
}
