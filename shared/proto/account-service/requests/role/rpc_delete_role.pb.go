// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: account-service/requests/role/rpc_delete_role.proto

package role

import (
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

type DeleteRoleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRoleRequest) Reset() {
	*x = DeleteRoleRequest{}
	mi := &file_account_service_requests_role_rpc_delete_role_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRoleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRoleRequest) ProtoMessage() {}

func (x *DeleteRoleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_requests_role_rpc_delete_role_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRoleRequest.ProtoReflect.Descriptor instead.
func (*DeleteRoleRequest) Descriptor() ([]byte, []int) {
	return file_account_service_requests_role_rpc_delete_role_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteRoleRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_account_service_requests_role_rpc_delete_role_proto protoreflect.FileDescriptor

const file_account_service_requests_role_rpc_delete_role_proto_rawDesc = "" +
	"\n" +
	"3account-service/requests/role/rpc_delete_role.proto\x12)account_service.requests.role.delete_role\"#\n" +
	"\x11DeleteRoleRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02idBKZIgithub.com/nicodanke/gesty-api/shared/proto/account-service/requests/roleb\x06proto3"

var (
	file_account_service_requests_role_rpc_delete_role_proto_rawDescOnce sync.Once
	file_account_service_requests_role_rpc_delete_role_proto_rawDescData []byte
)

func file_account_service_requests_role_rpc_delete_role_proto_rawDescGZIP() []byte {
	file_account_service_requests_role_rpc_delete_role_proto_rawDescOnce.Do(func() {
		file_account_service_requests_role_rpc_delete_role_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_account_service_requests_role_rpc_delete_role_proto_rawDesc), len(file_account_service_requests_role_rpc_delete_role_proto_rawDesc)))
	})
	return file_account_service_requests_role_rpc_delete_role_proto_rawDescData
}

var file_account_service_requests_role_rpc_delete_role_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_account_service_requests_role_rpc_delete_role_proto_goTypes = []any{
	(*DeleteRoleRequest)(nil), // 0: account_service.requests.role.delete_role.DeleteRoleRequest
}
var file_account_service_requests_role_rpc_delete_role_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_account_service_requests_role_rpc_delete_role_proto_init() }
func file_account_service_requests_role_rpc_delete_role_proto_init() {
	if File_account_service_requests_role_rpc_delete_role_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_account_service_requests_role_rpc_delete_role_proto_rawDesc), len(file_account_service_requests_role_rpc_delete_role_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_service_requests_role_rpc_delete_role_proto_goTypes,
		DependencyIndexes: file_account_service_requests_role_rpc_delete_role_proto_depIdxs,
		MessageInfos:      file_account_service_requests_role_rpc_delete_role_proto_msgTypes,
	}.Build()
	File_account_service_requests_role_rpc_delete_role_proto = out.File
	file_account_service_requests_role_rpc_delete_role_proto_goTypes = nil
	file_account_service_requests_role_rpc_delete_role_proto_depIdxs = nil
}
