// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: employee-service/requests/employee/rpc_update_employee.proto

package employee

import (
	models "github.com/nicodanke/gesty-api/shared/proto/employee-service/models"
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

type UpdateEmployeeRequest struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	Id                  int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                *string                `protobuf:"bytes,2,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Lastname            *string                `protobuf:"bytes,3,opt,name=lastname,proto3,oneof" json:"lastname,omitempty"`
	Email               *string                `protobuf:"bytes,4,opt,name=email,proto3,oneof" json:"email,omitempty"`
	Phone               *string                `protobuf:"bytes,5,opt,name=phone,proto3,oneof" json:"phone,omitempty"`
	Gender              *string                `protobuf:"bytes,6,opt,name=gender,proto3,oneof" json:"gender,omitempty"`
	RealId              *string                `protobuf:"bytes,7,opt,name=realId,proto3,oneof" json:"realId,omitempty"`
	FiscalId            *string                `protobuf:"bytes,8,opt,name=fiscalId,proto3,oneof" json:"fiscalId,omitempty"`
	AddressCountry      *string                `protobuf:"bytes,9,opt,name=addressCountry,proto3,oneof" json:"addressCountry,omitempty"`
	AddressState        *string                `protobuf:"bytes,10,opt,name=addressState,proto3,oneof" json:"addressState,omitempty"`
	AddressSubState     *string                `protobuf:"bytes,11,opt,name=addressSubState,proto3,oneof" json:"addressSubState,omitempty"`
	AddressStreet       *string                `protobuf:"bytes,12,opt,name=addressStreet,proto3,oneof" json:"addressStreet,omitempty"`
	AddressNumber       *string                `protobuf:"bytes,13,opt,name=addressNumber,proto3,oneof" json:"addressNumber,omitempty"`
	AddressZipCode      *string                `protobuf:"bytes,14,opt,name=addressZipCode,proto3,oneof" json:"addressZipCode,omitempty"`
	AddressUnit         *string                `protobuf:"bytes,15,opt,name=addressUnit,proto3,oneof" json:"addressUnit,omitempty"`
	AddressLat          *float64               `protobuf:"fixed64,16,opt,name=addressLat,proto3,oneof" json:"addressLat,omitempty"`
	AddressLng          *float64               `protobuf:"fixed64,17,opt,name=addressLng,proto3,oneof" json:"addressLng,omitempty"`
	FacilityIds         []int64                `protobuf:"varint,18,rep,packed,name=facilityIds,proto3" json:"facilityIds,omitempty"`
	RemoveAllFacilities *bool                  `protobuf:"varint,19,opt,name=removeAllFacilities,proto3,oneof" json:"removeAllFacilities,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *UpdateEmployeeRequest) Reset() {
	*x = UpdateEmployeeRequest{}
	mi := &file_employee_service_requests_employee_rpc_update_employee_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmployeeRequest) ProtoMessage() {}

func (x *UpdateEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_service_requests_employee_rpc_update_employee_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmployeeRequest.ProtoReflect.Descriptor instead.
func (*UpdateEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_employee_service_requests_employee_rpc_update_employee_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateEmployeeRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateEmployeeRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetLastname() string {
	if x != nil && x.Lastname != nil {
		return *x.Lastname
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetPhone() string {
	if x != nil && x.Phone != nil {
		return *x.Phone
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetGender() string {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetRealId() string {
	if x != nil && x.RealId != nil {
		return *x.RealId
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetFiscalId() string {
	if x != nil && x.FiscalId != nil {
		return *x.FiscalId
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetAddressCountry() string {
	if x != nil && x.AddressCountry != nil {
		return *x.AddressCountry
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetAddressState() string {
	if x != nil && x.AddressState != nil {
		return *x.AddressState
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetAddressSubState() string {
	if x != nil && x.AddressSubState != nil {
		return *x.AddressSubState
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetAddressStreet() string {
	if x != nil && x.AddressStreet != nil {
		return *x.AddressStreet
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetAddressNumber() string {
	if x != nil && x.AddressNumber != nil {
		return *x.AddressNumber
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetAddressZipCode() string {
	if x != nil && x.AddressZipCode != nil {
		return *x.AddressZipCode
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetAddressUnit() string {
	if x != nil && x.AddressUnit != nil {
		return *x.AddressUnit
	}
	return ""
}

func (x *UpdateEmployeeRequest) GetAddressLat() float64 {
	if x != nil && x.AddressLat != nil {
		return *x.AddressLat
	}
	return 0
}

func (x *UpdateEmployeeRequest) GetAddressLng() float64 {
	if x != nil && x.AddressLng != nil {
		return *x.AddressLng
	}
	return 0
}

func (x *UpdateEmployeeRequest) GetFacilityIds() []int64 {
	if x != nil {
		return x.FacilityIds
	}
	return nil
}

func (x *UpdateEmployeeRequest) GetRemoveAllFacilities() bool {
	if x != nil && x.RemoveAllFacilities != nil {
		return *x.RemoveAllFacilities
	}
	return false
}

type UpdateEmployeeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Employee      *models.Employee       `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateEmployeeResponse) Reset() {
	*x = UpdateEmployeeResponse{}
	mi := &file_employee_service_requests_employee_rpc_update_employee_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmployeeResponse) ProtoMessage() {}

func (x *UpdateEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employee_service_requests_employee_rpc_update_employee_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmployeeResponse.ProtoReflect.Descriptor instead.
func (*UpdateEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_employee_service_requests_employee_rpc_update_employee_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateEmployeeResponse) GetEmployee() *models.Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

var File_employee_service_requests_employee_rpc_update_employee_proto protoreflect.FileDescriptor

const file_employee_service_requests_employee_rpc_update_employee_proto_rawDesc = "" +
	"\n" +
	"<employee-service/requests/employee/rpc_update_employee.proto\x122employee_service.requests.employee.update_employee\x1a&employee-service/models/employee.proto\"\xc6\a\n" +
	"\x15UpdateEmployeeRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x17\n" +
	"\x04name\x18\x02 \x01(\tH\x00R\x04name\x88\x01\x01\x12\x1f\n" +
	"\blastname\x18\x03 \x01(\tH\x01R\blastname\x88\x01\x01\x12\x19\n" +
	"\x05email\x18\x04 \x01(\tH\x02R\x05email\x88\x01\x01\x12\x19\n" +
	"\x05phone\x18\x05 \x01(\tH\x03R\x05phone\x88\x01\x01\x12\x1b\n" +
	"\x06gender\x18\x06 \x01(\tH\x04R\x06gender\x88\x01\x01\x12\x1b\n" +
	"\x06realId\x18\a \x01(\tH\x05R\x06realId\x88\x01\x01\x12\x1f\n" +
	"\bfiscalId\x18\b \x01(\tH\x06R\bfiscalId\x88\x01\x01\x12+\n" +
	"\x0eaddressCountry\x18\t \x01(\tH\aR\x0eaddressCountry\x88\x01\x01\x12'\n" +
	"\faddressState\x18\n" +
	" \x01(\tH\bR\faddressState\x88\x01\x01\x12-\n" +
	"\x0faddressSubState\x18\v \x01(\tH\tR\x0faddressSubState\x88\x01\x01\x12)\n" +
	"\raddressStreet\x18\f \x01(\tH\n" +
	"R\raddressStreet\x88\x01\x01\x12)\n" +
	"\raddressNumber\x18\r \x01(\tH\vR\raddressNumber\x88\x01\x01\x12+\n" +
	"\x0eaddressZipCode\x18\x0e \x01(\tH\fR\x0eaddressZipCode\x88\x01\x01\x12%\n" +
	"\vaddressUnit\x18\x0f \x01(\tH\rR\vaddressUnit\x88\x01\x01\x12#\n" +
	"\n" +
	"addressLat\x18\x10 \x01(\x01H\x0eR\n" +
	"addressLat\x88\x01\x01\x12#\n" +
	"\n" +
	"addressLng\x18\x11 \x01(\x01H\x0fR\n" +
	"addressLng\x88\x01\x01\x12 \n" +
	"\vfacilityIds\x18\x12 \x03(\x03R\vfacilityIds\x125\n" +
	"\x13removeAllFacilities\x18\x13 \x01(\bH\x10R\x13removeAllFacilities\x88\x01\x01B\a\n" +
	"\x05_nameB\v\n" +
	"\t_lastnameB\b\n" +
	"\x06_emailB\b\n" +
	"\x06_phoneB\t\n" +
	"\a_genderB\t\n" +
	"\a_realIdB\v\n" +
	"\t_fiscalIdB\x11\n" +
	"\x0f_addressCountryB\x0f\n" +
	"\r_addressStateB\x12\n" +
	"\x10_addressSubStateB\x10\n" +
	"\x0e_addressStreetB\x10\n" +
	"\x0e_addressNumberB\x11\n" +
	"\x0f_addressZipCodeB\x0e\n" +
	"\f_addressUnitB\r\n" +
	"\v_addressLatB\r\n" +
	"\v_addressLngB\x16\n" +
	"\x14_removeAllFacilities\"`\n" +
	"\x16UpdateEmployeeResponse\x12F\n" +
	"\bemployee\x18\x01 \x01(\v2*.employee_service.models.employee.EmployeeR\bemployeeBPZNgithub.com/nicodanke/gesty-api/shared/proto/employee-service/requests/employeeb\x06proto3"

var (
	file_employee_service_requests_employee_rpc_update_employee_proto_rawDescOnce sync.Once
	file_employee_service_requests_employee_rpc_update_employee_proto_rawDescData []byte
)

func file_employee_service_requests_employee_rpc_update_employee_proto_rawDescGZIP() []byte {
	file_employee_service_requests_employee_rpc_update_employee_proto_rawDescOnce.Do(func() {
		file_employee_service_requests_employee_rpc_update_employee_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_employee_service_requests_employee_rpc_update_employee_proto_rawDesc), len(file_employee_service_requests_employee_rpc_update_employee_proto_rawDesc)))
	})
	return file_employee_service_requests_employee_rpc_update_employee_proto_rawDescData
}

var file_employee_service_requests_employee_rpc_update_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_employee_service_requests_employee_rpc_update_employee_proto_goTypes = []any{
	(*UpdateEmployeeRequest)(nil),  // 0: employee_service.requests.employee.update_employee.UpdateEmployeeRequest
	(*UpdateEmployeeResponse)(nil), // 1: employee_service.requests.employee.update_employee.UpdateEmployeeResponse
	(*models.Employee)(nil),        // 2: employee_service.models.employee.Employee
}
var file_employee_service_requests_employee_rpc_update_employee_proto_depIdxs = []int32{
	2, // 0: employee_service.requests.employee.update_employee.UpdateEmployeeResponse.employee:type_name -> employee_service.models.employee.Employee
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_employee_service_requests_employee_rpc_update_employee_proto_init() }
func file_employee_service_requests_employee_rpc_update_employee_proto_init() {
	if File_employee_service_requests_employee_rpc_update_employee_proto != nil {
		return
	}
	file_employee_service_requests_employee_rpc_update_employee_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_employee_service_requests_employee_rpc_update_employee_proto_rawDesc), len(file_employee_service_requests_employee_rpc_update_employee_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_employee_service_requests_employee_rpc_update_employee_proto_goTypes,
		DependencyIndexes: file_employee_service_requests_employee_rpc_update_employee_proto_depIdxs,
		MessageInfos:      file_employee_service_requests_employee_rpc_update_employee_proto_msgTypes,
	}.Build()
	File_employee_service_requests_employee_rpc_update_employee_proto = out.File
	file_employee_service_requests_employee_rpc_update_employee_proto_goTypes = nil
	file_employee_service_requests_employee_rpc_update_employee_proto_depIdxs = nil
}
