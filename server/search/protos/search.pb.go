// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: protos/search.proto

package protos

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

// Aggregate io description
type AggregateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AggregateRequest) Reset() {
	*x = AggregateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_search_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateRequest) ProtoMessage() {}

func (x *AggregateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_search_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateRequest.ProtoReflect.Descriptor instead.
func (*AggregateRequest) Descriptor() ([]byte, []int) {
	return file_protos_search_proto_rawDescGZIP(), []int{0}
}

type AggregateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AggregateResponse) Reset() {
	*x = AggregateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_search_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregateResponse) ProtoMessage() {}

func (x *AggregateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_search_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregateResponse.ProtoReflect.Descriptor instead.
func (*AggregateResponse) Descriptor() ([]byte, []int) {
	return file_protos_search_proto_rawDescGZIP(), []int{1}
}

var File_protos_search_proto protoreflect.FileDescriptor

var file_protos_search_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x12, 0x0a,
	0x10, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x13, 0x0a, 0x11, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x4a, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x12, 0x40, 0x0a, 0x09, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1b, 0x5a, 0x19, 0x6f, 0x6b, 0x61, 0x70, 0x69, 0x2d, 0x64, 0x61, 0x74, 0x61,
	0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_search_proto_rawDescOnce sync.Once
	file_protos_search_proto_rawDescData = file_protos_search_proto_rawDesc
)

func file_protos_search_proto_rawDescGZIP() []byte {
	file_protos_search_proto_rawDescOnce.Do(func() {
		file_protos_search_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_search_proto_rawDescData)
	})
	return file_protos_search_proto_rawDescData
}

var file_protos_search_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_search_proto_goTypes = []interface{}{
	(*AggregateRequest)(nil),  // 0: search.AggregateRequest
	(*AggregateResponse)(nil), // 1: search.AggregateResponse
}
var file_protos_search_proto_depIdxs = []int32{
	0, // 0: search.Search.Aggregate:input_type -> search.AggregateRequest
	1, // 1: search.Search.Aggregate:output_type -> search.AggregateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_search_proto_init() }
func file_protos_search_proto_init() {
	if File_protos_search_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_search_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateRequest); i {
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
		file_protos_search_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregateResponse); i {
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
			RawDescriptor: file_protos_search_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_search_proto_goTypes,
		DependencyIndexes: file_protos_search_proto_depIdxs,
		MessageInfos:      file_protos_search_proto_msgTypes,
	}.Build()
	File_protos_search_proto = out.File
	file_protos_search_proto_rawDesc = nil
	file_protos_search_proto_goTypes = nil
	file_protos_search_proto_depIdxs = nil
}