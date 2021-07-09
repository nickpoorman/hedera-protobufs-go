// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: ContractGetRecords.proto

package services

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Before v0.9.0, requested records of all transactions against the given contract in the last 25 hours.
//
// Deprecated: Do not use.
type ContractGetRecordsQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`         // Standard info sent from client to node, including the signed payment, and what kind of response is requested (cost, state proof, both, or neither).
	ContractID *ContractID  `protobuf:"bytes,2,opt,name=contractID,proto3" json:"contractID,omitempty"` // The smart contract instance for which the records should be retrieved
}

func (x *ContractGetRecordsQuery) Reset() {
	*x = ContractGetRecordsQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ContractGetRecords_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractGetRecordsQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractGetRecordsQuery) ProtoMessage() {}

func (x *ContractGetRecordsQuery) ProtoReflect() protoreflect.Message {
	mi := &file_ContractGetRecords_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractGetRecordsQuery.ProtoReflect.Descriptor instead.
func (*ContractGetRecordsQuery) Descriptor() ([]byte, []int) {
	return file_ContractGetRecords_proto_rawDescGZIP(), []int{0}
}

func (x *ContractGetRecordsQuery) GetHeader() *QueryHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ContractGetRecordsQuery) GetContractID() *ContractID {
	if x != nil {
		return x.ContractID
	}
	return nil
}

// Before v0.9.0, returned records of all transactions against the given contract in the last 25 hours.
//
// Deprecated: Do not use.
type ContractGetRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header     *ResponseHeader      `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`         // Standard response from node to client, including the requested fields: cost, or state proof, or both, or neither
	ContractID *ContractID          `protobuf:"bytes,2,opt,name=contractID,proto3" json:"contractID,omitempty"` // The smart contract instance that this record is for
	Records    []*TransactionRecord `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`       // List of records, each with contractCreateResult or contractCallResult as its body
}

func (x *ContractGetRecordsResponse) Reset() {
	*x = ContractGetRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ContractGetRecords_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractGetRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractGetRecordsResponse) ProtoMessage() {}

func (x *ContractGetRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ContractGetRecords_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractGetRecordsResponse.ProtoReflect.Descriptor instead.
func (*ContractGetRecordsResponse) Descriptor() ([]byte, []int) {
	return file_ContractGetRecords_proto_rawDescGZIP(), []int{1}
}

func (x *ContractGetRecordsResponse) GetHeader() *ResponseHeader {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *ContractGetRecordsResponse) GetContractID() *ContractID {
	if x != nil {
		return x.ContractID
	}
	return nil
}

func (x *ContractGetRecordsResponse) GetRecords() []*TransactionRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

var File_ContractGetRecords_proto protoreflect.FileDescriptor

var file_ContractGetRecords_proto_rawDesc = []byte{
	0x0a, 0x18, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x10, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x14, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7c, 0x0a, 0x17, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x2a, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x48, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0a,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x49, 0x44, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x3a,
	0x02, 0x18, 0x01, 0x22, 0xb6, 0x01, 0x0a, 0x1a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2d, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x49, 0x44, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x49, 0x44, 0x12, 0x32, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52,
	0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x3a, 0x02, 0x18, 0x01, 0x42, 0x26, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61,
	0x76, 0x61, 0x50, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ContractGetRecords_proto_rawDescOnce sync.Once
	file_ContractGetRecords_proto_rawDescData = file_ContractGetRecords_proto_rawDesc
)

func file_ContractGetRecords_proto_rawDescGZIP() []byte {
	file_ContractGetRecords_proto_rawDescOnce.Do(func() {
		file_ContractGetRecords_proto_rawDescData = protoimpl.X.CompressGZIP(file_ContractGetRecords_proto_rawDescData)
	})
	return file_ContractGetRecords_proto_rawDescData
}

var (
	file_ContractGetRecords_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_ContractGetRecords_proto_goTypes  = []interface{}{
		(*ContractGetRecordsQuery)(nil),    // 0: proto.ContractGetRecordsQuery
		(*ContractGetRecordsResponse)(nil), // 1: proto.ContractGetRecordsResponse
		(*QueryHeader)(nil),                // 2: proto.QueryHeader
		(*ContractID)(nil),                 // 3: proto.ContractID
		(*ResponseHeader)(nil),             // 4: proto.ResponseHeader
		(*TransactionRecord)(nil),          // 5: proto.TransactionRecord
	}
)

var file_ContractGetRecords_proto_depIdxs = []int32{
	2, // 0: proto.ContractGetRecordsQuery.header:type_name -> proto.QueryHeader
	3, // 1: proto.ContractGetRecordsQuery.contractID:type_name -> proto.ContractID
	4, // 2: proto.ContractGetRecordsResponse.header:type_name -> proto.ResponseHeader
	3, // 3: proto.ContractGetRecordsResponse.contractID:type_name -> proto.ContractID
	5, // 4: proto.ContractGetRecordsResponse.records:type_name -> proto.TransactionRecord
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ContractGetRecords_proto_init() }
func file_ContractGetRecords_proto_init() {
	if File_ContractGetRecords_proto != nil {
		return
	}
	file_BasicTypes_proto_init()
	file_TransactionRecord_proto_init()
	file_QueryHeader_proto_init()
	file_ResponseHeader_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ContractGetRecords_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractGetRecordsQuery); i {
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
		file_ContractGetRecords_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractGetRecordsResponse); i {
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
			RawDescriptor: file_ContractGetRecords_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ContractGetRecords_proto_goTypes,
		DependencyIndexes: file_ContractGetRecords_proto_depIdxs,
		MessageInfos:      file_ContractGetRecords_proto_msgTypes,
	}.Build()
	File_ContractGetRecords_proto = out.File
	file_ContractGetRecords_proto_rawDesc = nil
	file_ContractGetRecords_proto_goTypes = nil
	file_ContractGetRecords_proto_depIdxs = nil
}
