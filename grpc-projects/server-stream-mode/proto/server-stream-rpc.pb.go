// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: server-stream-rpc.proto

package proto

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

type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestString string `protobuf:"bytes,1,opt,name=RequestString,proto3" json:"RequestString,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_stream_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_stream_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_server_stream_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetRequestString() string {
	if x != nil {
		return x.RequestString
	}
	return ""
}

type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseString string `protobuf:"bytes,1,opt,name=ResponseString,proto3" json:"ResponseString,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_stream_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_stream_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_server_stream_rpc_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetResponseString() string {
	if x != nil {
		return x.ResponseString
	}
	return ""
}

var File_server_stream_rpc_proto protoreflect.FileDescriptor

var file_server_stream_rpc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2d,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22,
	0x34, 0x0a, 0x0c, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x37, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x32, 0x4e,
	0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3f, 0x0a,
	0x10, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x12, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x1b,
	0x5a, 0x19, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2d,
	0x6d, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_server_stream_rpc_proto_rawDescOnce sync.Once
	file_server_stream_rpc_proto_rawDescData = file_server_stream_rpc_proto_rawDesc
)

func file_server_stream_rpc_proto_rawDescGZIP() []byte {
	file_server_stream_rpc_proto_rawDescOnce.Do(func() {
		file_server_stream_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_stream_rpc_proto_rawDescData)
	})
	return file_server_stream_rpc_proto_rawDescData
}

var file_server_stream_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_server_stream_rpc_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),  // 0: chat.HelloRequest
	(*HelloResponse)(nil), // 1: chat.HelloResponse
}
var file_server_stream_rpc_proto_depIdxs = []int32{
	0, // 0: chat.ChatService.SendServerStream:input_type -> chat.HelloRequest
	1, // 1: chat.ChatService.SendServerStream:output_type -> chat.HelloResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_server_stream_rpc_proto_init() }
func file_server_stream_rpc_proto_init() {
	if File_server_stream_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_stream_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_server_stream_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
			RawDescriptor: file_server_stream_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_stream_rpc_proto_goTypes,
		DependencyIndexes: file_server_stream_rpc_proto_depIdxs,
		MessageInfos:      file_server_stream_rpc_proto_msgTypes,
	}.Build()
	File_server_stream_rpc_proto = out.File
	file_server_stream_rpc_proto_rawDesc = nil
	file_server_stream_rpc_proto_goTypes = nil
	file_server_stream_rpc_proto_depIdxs = nil
}
