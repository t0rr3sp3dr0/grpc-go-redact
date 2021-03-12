// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: test/test.proto

package test

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

type TestStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StandardValue string `protobuf:"bytes,1,opt,name=StandardValue,proto3" json:"StandardValue,omitempty"`
	SecretValue   string `protobuf:"bytes,2,opt,name=SecretValue,proto3" json:"SecretValue,omitempty"`
}

func (x *TestStruct) Reset() {
	*x = TestStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hcp_types_agentpool_v1_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestStruct) ProtoMessage() {}

func (x *TestStruct) ProtoReflect() protoreflect.Message {
	mi := &file_hcp_types_agentpool_v1_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestStruct.ProtoReflect.Descriptor instead.
func (*TestStruct) Descriptor() ([]byte, []int) {
	return file_hcp_types_agentpool_v1_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestStruct) GetStandardValue() string {
	if x != nil {
		return x.StandardValue
	}
	return ""
}

func (x *TestStruct) GetSecretValue() string {
	if x != nil {
		return x.SecretValue
	}
	return ""
}

var File_hcp_types_agentpool_v1_test_proto protoreflect.FileDescriptor

var file_hcp_types_agentpool_v1_test_proto_rawDesc = []byte{
	0x0a, 0x21, 0x68, 0x63, 0x70, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x16, 0x68, 0x63, 0x70, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x61,
	0x67, 0x65, 0x6e, 0x74, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x22, 0x54, 0x0a, 0x0a, 0x54,
	0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x74, 0x61,
	0x6e, 0x64, 0x61, 0x72, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x6f, 0x6d, 0x73, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x6b, 0x73,
	0x2f, 0x72, 0x70, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x68, 0x63, 0x70, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hcp_types_agentpool_v1_test_proto_rawDescOnce sync.Once
	file_hcp_types_agentpool_v1_test_proto_rawDescData = file_hcp_types_agentpool_v1_test_proto_rawDesc
)

func file_hcp_types_agentpool_v1_test_proto_rawDescGZIP() []byte {
	file_hcp_types_agentpool_v1_test_proto_rawDescOnce.Do(func() {
		file_hcp_types_agentpool_v1_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_hcp_types_agentpool_v1_test_proto_rawDescData)
	})
	return file_hcp_types_agentpool_v1_test_proto_rawDescData
}

var file_hcp_types_agentpool_v1_test_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_hcp_types_agentpool_v1_test_proto_goTypes = []interface{}{
	(*TestStruct)(nil), // 0: hcp.types.agentpool.v1.TestStruct
}
var file_hcp_types_agentpool_v1_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hcp_types_agentpool_v1_test_proto_init() }
func file_hcp_types_agentpool_v1_test_proto_init() {
	if File_hcp_types_agentpool_v1_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hcp_types_agentpool_v1_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestStruct); i {
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
			RawDescriptor: file_hcp_types_agentpool_v1_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hcp_types_agentpool_v1_test_proto_goTypes,
		DependencyIndexes: file_hcp_types_agentpool_v1_test_proto_depIdxs,
		MessageInfos:      file_hcp_types_agentpool_v1_test_proto_msgTypes,
	}.Build()
	File_hcp_types_agentpool_v1_test_proto = out.File
	file_hcp_types_agentpool_v1_test_proto_rawDesc = nil
	file_hcp_types_agentpool_v1_test_proto_goTypes = nil
	file_hcp_types_agentpool_v1_test_proto_depIdxs = nil
}
