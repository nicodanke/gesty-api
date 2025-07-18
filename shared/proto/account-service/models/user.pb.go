// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: account-service/models/user.proto

package models

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type User struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	Id                int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username          string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Name              string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Lastname          string                 `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email             string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone             string                 `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Active            bool                   `protobuf:"varint,7,opt,name=active,proto3" json:"active,omitempty"`
	IsAdmin           bool                   `protobuf:"varint,8,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
	RoleId            int64                  `protobuf:"varint,9,opt,name=roleId,proto3" json:"roleId,omitempty"`
	PasswordChangedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=passwordChangedAt,proto3" json:"passwordChangedAt,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_account_service_models_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_account_service_models_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_account_service_models_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *User) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *User) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *User) GetRoleId() int64 {
	if x != nil {
		return x.RoleId
	}
	return 0
}

func (x *User) GetPasswordChangedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.PasswordChangedAt
	}
	return nil
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_account_service_models_user_proto protoreflect.FileDescriptor

const file_account_service_models_user_proto_rawDesc = "" +
	"\n" +
	"!account-service/models/user.proto\x12\x1baccount_service.models.user\x1a\x1fgoogle/protobuf/timestamp.proto\"\xdc\x02\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12\x1a\n" +
	"\blastname\x18\x04 \x01(\tR\blastname\x12\x14\n" +
	"\x05email\x18\x05 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x06 \x01(\tR\x05phone\x12\x16\n" +
	"\x06active\x18\a \x01(\bR\x06active\x12\x18\n" +
	"\aisAdmin\x18\b \x01(\bR\aisAdmin\x12\x16\n" +
	"\x06roleId\x18\t \x01(\x03R\x06roleId\x12H\n" +
	"\x11passwordChangedAt\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampR\x11passwordChangedAt\x128\n" +
	"\tcreatedAt\x18\v \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAtBDZBgithub.com/nicodanke/gesty-api/shared/proto/account-service/modelsb\x06proto3"

var (
	file_account_service_models_user_proto_rawDescOnce sync.Once
	file_account_service_models_user_proto_rawDescData []byte
)

func file_account_service_models_user_proto_rawDescGZIP() []byte {
	file_account_service_models_user_proto_rawDescOnce.Do(func() {
		file_account_service_models_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_account_service_models_user_proto_rawDesc), len(file_account_service_models_user_proto_rawDesc)))
	})
	return file_account_service_models_user_proto_rawDescData
}

var file_account_service_models_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_account_service_models_user_proto_goTypes = []any{
	(*User)(nil),                  // 0: account_service.models.user.User
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_account_service_models_user_proto_depIdxs = []int32{
	1, // 0: account_service.models.user.User.passwordChangedAt:type_name -> google.protobuf.Timestamp
	1, // 1: account_service.models.user.User.createdAt:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_account_service_models_user_proto_init() }
func file_account_service_models_user_proto_init() {
	if File_account_service_models_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_account_service_models_user_proto_rawDesc), len(file_account_service_models_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_account_service_models_user_proto_goTypes,
		DependencyIndexes: file_account_service_models_user_proto_depIdxs,
		MessageInfos:      file_account_service_models_user_proto_msgTypes,
	}.Build()
	File_account_service_models_user_proto = out.File
	file_account_service_models_user_proto_goTypes = nil
	file_account_service_models_user_proto_depIdxs = nil
}
