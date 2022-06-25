// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: token_wipe_account.proto

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
// Wipes the provided amount of tokens from the specified Account. Must be signed by the Token's
// Wipe key.
// If the provided account is not found, the transaction will resolve to INVALID_ACCOUNT_ID.
// If the provided account has been deleted, the transaction will resolve to ACCOUNT_DELETED.
// If the provided token is not found, the transaction will resolve to INVALID_TOKEN_ID.
// If the provided token has been deleted, the transaction will resolve to TOKEN_WAS_DELETED.
// If an Association between the provided token and account is not found, the transaction will
// resolve to TOKEN_NOT_ASSOCIATED_TO_ACCOUNT.
// If Wipe Key is not present in the Token, transaction results in TOKEN_HAS_NO_WIPE_KEY.
// If the provided account is the Token's Treasury Account, transaction results in
// CANNOT_WIPE_TOKEN_TREASURY_ACCOUNT
// On success, tokens are removed from the account and the total supply of the token is decreased by
// the wiped amount.
//
// If both amount and serialNumbers get filled, a INVALID_TRANSACTION_BODY response code will be
// returned.
// If neither the amount nor the serialNumbers get filled, a INVALID_WIPING_AMOUNT response code
// will be returned.
// If the serialNumbers list contains a non-positive integer as a serial number, a INVALID_NFT_ID
// response code will be returned.
// If the serialNumbers' list count is greater than the batch size limit global dynamic property, a
// BATCH_SIZE_LIMIT_EXCEEDED response code will be returned.
//
// The amount provided is in the lowest denomination possible. Example:
// Token A has 2 decimals. In order to wipe 100 tokens from account, one must provide amount of
// 10000. In order to wipe 100.55 tokens, one must provide amount of 10055.
type TokenWipeAccountTransactionBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//*
	// The token for which the account will be wiped. If token does not exist, transaction results
	// in INVALID_TOKEN_ID
	Token *TokenID `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	//*
	// The account to be wiped
	Account *AccountID `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	//*
	// Applicable to tokens of type FUNGIBLE_COMMON. The amount of tokens to wipe from the specified
	// account. Amount must be a positive non-zero number in the lowest denomination possible, not
	// bigger than the token balance of the account (0; balance]
	Amount uint64 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
	//*
	// Applicable to tokens of type NON_FUNGIBLE_UNIQUE. The list of serial numbers to be wiped.
	SerialNumbers []int64 `protobuf:"varint,4,rep,packed,name=serialNumbers,proto3" json:"serialNumbers,omitempty"`
}

func (x *TokenWipeAccountTransactionBody) Reset() {
	*x = TokenWipeAccountTransactionBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_wipe_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenWipeAccountTransactionBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenWipeAccountTransactionBody) ProtoMessage() {}

func (x *TokenWipeAccountTransactionBody) ProtoReflect() protoreflect.Message {
	mi := &file_token_wipe_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenWipeAccountTransactionBody.ProtoReflect.Descriptor instead.
func (*TokenWipeAccountTransactionBody) Descriptor() ([]byte, []int) {
	return file_token_wipe_account_proto_rawDescGZIP(), []int{0}
}

func (x *TokenWipeAccountTransactionBody) GetToken() *TokenID {
	if x != nil {
		return x.Token
	}
	return nil
}

func (x *TokenWipeAccountTransactionBody) GetAccount() *AccountID {
	if x != nil {
		return x.Account
	}
	return nil
}

func (x *TokenWipeAccountTransactionBody) GetAmount() uint64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TokenWipeAccountTransactionBody) GetSerialNumbers() []int64 {
	if x != nil {
		return x.SerialNumbers
	}
	return nil
}

var File_token_wipe_account_proto protoreflect.FileDescriptor

var file_token_wipe_account_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x77, 0x69, 0x70, 0x65, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x69, 0x63, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x1f, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x57, 0x69,
	0x70, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x24, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2a,
	0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49,
	0x44, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x72, 0x69, 0x61,
	0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x42, 0x26, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x65, 0x64, 0x65, 0x72, 0x61, 0x68, 0x61, 0x73, 0x68, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x6a, 0x61, 0x76, 0x61, 0x50, 0x01,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_wipe_account_proto_rawDescOnce sync.Once
	file_token_wipe_account_proto_rawDescData = file_token_wipe_account_proto_rawDesc
)

func file_token_wipe_account_proto_rawDescGZIP() []byte {
	file_token_wipe_account_proto_rawDescOnce.Do(func() {
		file_token_wipe_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_wipe_account_proto_rawDescData)
	})
	return file_token_wipe_account_proto_rawDescData
}

var file_token_wipe_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_token_wipe_account_proto_goTypes = []interface{}{
	(*TokenWipeAccountTransactionBody)(nil), // 0: proto.TokenWipeAccountTransactionBody
	(*TokenID)(nil),                         // 1: proto.TokenID
	(*AccountID)(nil),                       // 2: proto.AccountID
}
var file_token_wipe_account_proto_depIdxs = []int32{
	1, // 0: proto.TokenWipeAccountTransactionBody.token:type_name -> proto.TokenID
	2, // 1: proto.TokenWipeAccountTransactionBody.account:type_name -> proto.AccountID
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_token_wipe_account_proto_init() }
func file_token_wipe_account_proto_init() {
	if File_token_wipe_account_proto != nil {
		return
	}
	file_basic_types_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_token_wipe_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenWipeAccountTransactionBody); i {
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
			RawDescriptor: file_token_wipe_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_token_wipe_account_proto_goTypes,
		DependencyIndexes: file_token_wipe_account_proto_depIdxs,
		MessageInfos:      file_token_wipe_account_proto_msgTypes,
	}.Build()
	File_token_wipe_account_proto = out.File
	file_token_wipe_account_proto_rawDesc = nil
	file_token_wipe_account_proto_goTypes = nil
	file_token_wipe_account_proto_depIdxs = nil
}
