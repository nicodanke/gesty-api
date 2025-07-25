// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: account-service/models/role.proto

package models

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

type Role struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description   string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	PermissionIds []int64                `protobuf:"varint,4,rep,packed,name=permissionIds,proto3" json:"permissionIds,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Role) Reset() {
	*x = Role{}
	mi := &file_account_service_models_role_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_models_role_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_account_service_models_role_proto_rawDescGZIP(), []int{0}
}

func (x *Role) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Role) GetPermissionIds() []int64 {
	if x != nil {
		return x.PermissionIds
	}
	return nil
}

var File_account_service_models_role_proto protoreflect.FileDescriptor

const file_account_service_models_role_proto_rawDesc = "" +
	"\n" +
	"!account-service/models/role.proto\x12\x1baccount_service.models.role\"r\n" +
	"\x04Role\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12$\n" +
	"\rpermissionIds\x18\x04 \x03(\x03R\rpermissionIdsBDZBgithub.com/nicodanke/gesty-api/shared/proto/account-service/modelsb\x06proto3"

var (
	file_account_service_models_role_proto_rawDescOnce sync.Once
	file_account_service_models_role_proto_rawDescData []byte
)

func file_account_service_models_role_proto_rawDescGZIP() []byte {
	file_account_service_models_role_proto_rawDescOnce.Do(func() {
		file_account_service_models_role_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_account_service_models_role_proto_rawDesc), len(file_account_service_models_role_proto_rawDesc)))
	})
	return file_account_service_models_role_proto_rawDescData
}

var file_account_service_models_role_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_account_service_models_role_proto_goTypes = []any{
	(*Role)(nil), // 0: account_service.models.role.Role
}
var file_account_service_models_role_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_account_service_models_role_proto_init() }
func file_account_service_models_role_proto_init() {
	if File_account_service_models_role_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_account_service_models_role_proto_rawDesc), len(file_account_service_models_role_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_service_models_role_proto_goTypes,
		DependencyIndexes: file_account_service_models_role_proto_depIdxs,
		MessageInfos:      file_account_service_models_role_proto_msgTypes,
	}.Build()
	File_account_service_models_role_proto = out.File
	file_account_service_models_role_proto_goTypes = nil
	file_account_service_models_role_proto_depIdxs = nil
}
