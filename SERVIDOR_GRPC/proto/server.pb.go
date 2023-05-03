// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.3
// source: proto/server.proto

package proto

import (
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

type JsonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params string `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (x *JsonRequest) Reset() {
	*x = JsonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonRequest) ProtoMessage() {}

func (x *JsonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonRequest.ProtoReflect.Descriptor instead.
func (*JsonRequest) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{0}
}

func (x *JsonRequest) GetParams() string {
	if x != nil {
		return x.Params
	}
	return ""
}

type JsonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Json string `protobuf:"bytes,1,opt,name=json,proto3" json:"json,omitempty"`
}

func (x *JsonResponse) Reset() {
	*x = JsonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JsonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JsonResponse) ProtoMessage() {}

func (x *JsonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JsonResponse.ProtoReflect.Descriptor instead.
func (*JsonResponse) Descriptor() ([]byte, []int) {
	return file_proto_server_proto_rawDescGZIP(), []int{1}
}

func (x *JsonResponse) GetJson() string {
	if x != nil {
		return x.Json
	}
	return ""
}

var File_proto_server_proto protoreflect.FileDescriptor

var file_proto_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x25, 0x0a, 0x0b, 0x4a, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x22, 0x22, 0x0a, 0x0c, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6a, 0x73, 0x6f, 0x6e, 0x32, 0x58, 0x0a, 0x04, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x50, 0x0a, 0x0a,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4a, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x6a, 0x73,
	0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_server_proto_rawDescOnce sync.Once
	file_proto_server_proto_rawDescData = file_proto_server_proto_rawDesc
)

func file_proto_server_proto_rawDescGZIP() []byte {
	file_proto_server_proto_rawDescOnce.Do(func() {
		file_proto_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_server_proto_rawDescData)
	})
	return file_proto_server_proto_rawDescData
}

var file_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_server_proto_goTypes = []interface{}{
	(*JsonRequest)(nil),  // 0: proto.JsonRequest
	(*JsonResponse)(nil), // 1: proto.JsonResponse
}
var file_proto_server_proto_depIdxs = []int32{
	0, // 0: proto.Json.Connection:input_type -> proto.JsonRequest
	1, // 1: proto.Json.Connection:output_type -> proto.JsonResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_server_proto_init() }
func file_proto_server_proto_init() {
	if File_proto_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonRequest); i {
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
		file_proto_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JsonResponse); i {
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
			RawDescriptor: file_proto_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_server_proto_goTypes,
		DependencyIndexes: file_proto_server_proto_depIdxs,
		MessageInfos:      file_proto_server_proto_msgTypes,
	}.Build()
	File_proto_server_proto = out.File
	file_proto_server_proto_rawDesc = nil
	file_proto_server_proto_goTypes = nil
	file_proto_server_proto_depIdxs = nil
}
