// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: network_get_version_info.proto

package services

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

//*
// Get the deployed versions of Hedera Services and the HAPI proto in semantic version format
type NetworkGetVersionInfoQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard info sent from client to node, including the signed payment, and what kind of
	// response is requested (cost, state proof, both, or neither).
	Header *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *NetworkGetVersionInfoQuery) Reset() {
	*x = NetworkGetVersionInfoQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_get_version_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkGetVersionInfoQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkGetVersionInfoQuery) ProtoMessage() {}

func (x *NetworkGetVersionInfoQuery) ProtoReflect() protoreflect.Message {
	mi := &file_network_get_version_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkGetVersionInfoQuery.ProtoReflect.Descriptor instead.
func (*NetworkGetVersionInfoQuery) Descriptor() ([]byte, []int) {
	return file_network_get_version_info_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkGetVersionInfoQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

//*
// Response when the client sends the node NetworkGetVersionInfoQuery
type NetworkGetVersionInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// Standard response from node to client, including the requested fields: cost, or state proof,
	// or both, or neither
	Header *ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	//*
	// The Hedera API (HAPI) protobuf version recognized by the responding node.
	HapiProtoVersion *SemanticVersion `protobuf:"bytes,2,opt,name=hapiProtoVersion,proto3" json:"hapiProtoVersion,omitempty"`
	//*
	// The version of the Hedera Services software deployed on the responding node.
	HederaServicesVersion *SemanticVersion `protobuf:"bytes,3,opt,name=hederaServicesVersion,proto3" json:"hederaServicesVersion,omitempty"`
}

func (x *NetworkGetVersionInfoResponse) Reset() {
	*x = NetworkGetVersionInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_network_get_version_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkGetVersionInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkGetVersionInfoResponse) ProtoMessage() {}

func (x *NetworkGetVersionInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_network_get_version_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkGetVersionInfoResponse.ProtoReflect.Descriptor instead.
func (*NetworkGetVersionInfoResponse) Descriptor() ([]byte, []int) {
	return file_network_get_version_info_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkGetVersionInfoResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *NetworkGetVersionInfoResponse) GetHapiProtoVersion() *SemanticVersion {
	if x != nil {
		return x.HapiProtoVersion
	}
	return nil
}

func (x *NetworkGetVersionInfoResponse) GetHederaServicesVersion() *SemanticVersion {
	if x != nil {
		return x.HederaServicesVersion
	}
	return nil
}

var File_network_get_version_info_proto protoreflect.FileDescriptor

var file_network_get_version_info_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x71, 0x75, 0x65, 0x72,
	0x79, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x1a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22,
	0xe0, 0x01, 0x0a, 0x1d, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x47, 0x65, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x42, 0x0a, 0x10, 0x68, 0x61, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x10, 0x68, 0x61, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x15, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x65, 0x6d, 0x61,
	0x6e, 0x74, 0x69, 0x63, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x15, 0x68, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_network_get_version_info_proto_rawDescOnce sync.Once
	file_network_get_version_info_proto_rawDescData = file_network_get_version_info_proto_rawDesc
)

func file_network_get_version_info_proto_rawDescGZIP() []byte {
	file_network_get_version_info_proto_rawDescOnce.Do(func() {
		file_network_get_version_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_network_get_version_info_proto_rawDescData)
	})
	return file_network_get_version_info_proto_rawDescData
}

var file_network_get_version_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_network_get_version_info_proto_goTypes = []interface{}{
	(*NetworkGetVersionInfoQuery)(nil),    // 0: proto.NetworkGetVersionInfoQuery
	(*NetworkGetVersionInfoResponse)(nil), // 1: proto.NetworkGetVersionInfoResponse
	(*QueryHeader)(nil),                   // 2: proto.QueryHeader
	(*ResponseHeader)(nil),                // 3: proto.ResponseHeader
	(*SemanticVersion)(nil),               // 4: proto.SemanticVersion
}
var file_network_get_version_info_proto_depIdxs = []int32{
	2, // 0: proto.NetworkGetVersionInfoQuery.header:type_name -> proto.QueryHeader
	3, // 1: proto.NetworkGetVersionInfoResponse.header:type_name -> proto.ResponseHeader
	4, // 2: proto.NetworkGetVersionInfoResponse.hapiProtoVersion:type_name -> proto.SemanticVersion
	4, // 3: proto.NetworkGetVersionInfoResponse.hederaServicesVersion:type_name -> proto.SemanticVersion
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_network_get_version_info_proto_init() }
func file_network_get_version_info_proto_init() {
	if File_network_get_version_info_proto != nil {
		return
	}
	file_basic_types_proto_init()
	file_query_header_proto_init()
	file_response_header_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_network_get_version_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkGetVersionInfoQuery); i {
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
		file_network_get_version_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkGetVersionInfoResponse); i {
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
			RawDescriptor: file_network_get_version_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_network_get_version_info_proto_goTypes,
		DependencyIndexes: file_network_get_version_info_proto_depIdxs,
		MessageInfos:      file_network_get_version_info_proto_msgTypes,
	}.Build()
	File_network_get_version_info_proto = out.File
	file_network_get_version_info_proto_rawDesc = nil
	file_network_get_version_info_proto_goTypes = nil
	file_network_get_version_info_proto_depIdxs = nil
}
