// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: account-service/requests/account/rpc_update_account.proto

package account

import (
	models "github.com/nicodanke/gesty-api/shared/proto/account-service/models"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UpdateAccountRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CompanyName   *string                `protobuf:"bytes,2,opt,name=companyName,proto3,oneof" json:"companyName,omitempty"`
	Email         *string                `protobuf:"bytes,3,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Active        *bool                  `protobuf:"varint,4,opt,name=active,proto3,oneof" json:"active,omitempty"`
	Phone         *string                `protobuf:"bytes,5,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	WebUrl        *string                `protobuf:"bytes,6,opt,name=webUrl,proto3,oneof" json:"webUrl,omitempty"`
	Country       *string                `protobuf:"bytes,7,opt,name=country,proto3,oneof" json:"country,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAccountRequest) Reset() {
	*x = UpdateAccountRequest{}
	mi := &file_account_service_requests_account_rpc_update_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountRequest) ProtoMessage() {}

func (x *UpdateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_requests_account_rpc_update_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountRequest.ProtoReflect.Descriptor instead.
func (*UpdateAccountRequest) Descriptor() ([]byte, []int) {
	return file_account_service_requests_account_rpc_update_account_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateAccountRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAccountRequest) GetCompanyName() string {
	if x != nil && x.CompanyName != nil {
		return *x.CompanyName
	}
	return ""
}

func (x *UpdateAccountRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UpdateAccountRequest) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

func (x *UpdateAccountRequest) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *UpdateAccountRequest) GetWebUrl() string {
	if x != nil && x.WebUrl != nil {
		return *x.WebUrl
	}
	return ""
}

func (x *UpdateAccountRequest) GetCountry() string {
	if x != nil && x.Country != nil {
		return *x.Country
	}
	return ""
}

type UpdateAccountResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Account       *models.Account        `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateAccountResponse) Reset() {
	*x = UpdateAccountResponse{}
	mi := &file_account_service_requests_account_rpc_update_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAccountResponse) ProtoMessage() {}

func (x *UpdateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_requests_account_rpc_update_account_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAccountResponse.ProtoReflect.Descriptor instead.
func (*UpdateAccountResponse) Descriptor() ([]byte, []int) {
	return file_account_service_requests_account_rpc_update_account_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateAccountResponse) GetAccount() *models.Account {
	if x != nil {
		return x.Account
	}
	return nil
}

var File_account_service_requests_account_rpc_update_account_proto protoreflect.FileDescriptor

const file_account_service_requests_account_rpc_update_account_proto_rawDesc = "" +
	"\n" +
	"9account-service/requests/account/rpc_update_account.proto\x12/account_service.requests.account.update_account\x1a$account-service/models/account.proto\"\xa2\x02\n" +
	"\x14UpdateAccountRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12%\n" +
	"\vcompanyName\x18\x02 \x01(\tH\x00R\vcompanyName\x88\x01\x01\x12\x19\n" +
	"\x05email\x18\x03 \x01(\tH\x01R\x05email\x88\x01\x01\x12\x1b\n" +
	"\x06active\x18\x04 \x01(\bH\x02R\x06active\x88\x01\x01\x12\x19\n" +
	"\x05phone\x18\x05 \x01(\tH\x03R\x05phone\x88\x01\x01\x12\x1b\n" +
	"\x06webUrl\x18\x06 \x01(\tH\x04R\x06webUrl\x88\x01\x01\x12\x1d\n" +
	"\acountry\x18\a \x01(\tH\x05R\acountry\x88\x01\x01B\x0e\n" +
	"\f_companyNameB\b\n" +
	"\x06_emailB\t\n" +
	"\a_activeB\b\n" +
	"\x06_phoneB\t\n" +
	"\a_webUrlB\n" +
	"\n" +
	"\b_country\"Z\n" +
	"\x15UpdateAccountResponse\x12A\n" +
	"\aaccount\x18\x01 \x01(\v2'.account_service.models.account.AccountR\aaccountBNZLgithub.com/nicodanke/gesty-api/shared/proto/account-service/requests/accountb\x06proto3"

var (
	file_account_service_requests_account_rpc_update_account_proto_rawDescOnce sync.Once
	file_account_service_requests_account_rpc_update_account_proto_rawDescData []byte
)

func file_account_service_requests_account_rpc_update_account_proto_rawDescGZIP() []byte {
	file_account_service_requests_account_rpc_update_account_proto_rawDescOnce.Do(func() {
		file_account_service_requests_account_rpc_update_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_account_service_requests_account_rpc_update_account_proto_rawDesc), len(file_account_service_requests_account_rpc_update_account_proto_rawDesc)))
	})
	return file_account_service_requests_account_rpc_update_account_proto_rawDescData
}

var file_account_service_requests_account_rpc_update_account_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_account_service_requests_account_rpc_update_account_proto_goTypes = []any{
	(*UpdateAccountRequest)(nil),  // 0: account_service.requests.account.update_account.UpdateAccountRequest
	(*UpdateAccountResponse)(nil), // 1: account_service.requests.account.update_account.UpdateAccountResponse
	(*models.Account)(nil),        // 2: account_service.models.account.Account
}
var file_account_service_requests_account_rpc_update_account_proto_depIdxs = []int32{
	2, // 0: account_service.requests.account.update_account.UpdateAccountResponse.account:type_name -> account_service.models.account.Account
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_account_service_requests_account_rpc_update_account_proto_init() }
func file_account_service_requests_account_rpc_update_account_proto_init() {
	if File_account_service_requests_account_rpc_update_account_proto != nil {
		return
	}
	file_account_service_requests_account_rpc_update_account_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_account_service_requests_account_rpc_update_account_proto_rawDesc), len(file_account_service_requests_account_rpc_update_account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_service_requests_account_rpc_update_account_proto_goTypes,
		DependencyIndexes: file_account_service_requests_account_rpc_update_account_proto_depIdxs,
		MessageInfos:      file_account_service_requests_account_rpc_update_account_proto_msgTypes,
	}.Build()
	File_account_service_requests_account_rpc_update_account_proto = out.File
	file_account_service_requests_account_rpc_update_account_proto_goTypes = nil
	file_account_service_requests_account_rpc_update_account_proto_depIdxs = nil
}
