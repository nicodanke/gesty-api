// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v4.25.3
// source: employee-service/requests/employee/rpc_create_employee.proto

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

type CreateEmployeeRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Name            string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lastname        string                 `protobuf:"bytes,2,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email           string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone           string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Gender          string                 `protobuf:"bytes,5,opt,name=gender,proto3" json:"gender,omitempty"`
	RealId          string                 `protobuf:"bytes,6,opt,name=realId,proto3" json:"realId,omitempty"`
	FiscalId        string                 `protobuf:"bytes,7,opt,name=fiscalId,proto3" json:"fiscalId,omitempty"`
	AddressCountry  string                 `protobuf:"bytes,8,opt,name=addressCountry,proto3" json:"addressCountry,omitempty"`
	AddressState    string                 `protobuf:"bytes,9,opt,name=addressState,proto3" json:"addressState,omitempty"`
	AddressSubState *string                `protobuf:"bytes,10,opt,name=addressSubState,proto3,oneof" json:"addressSubState,omitempty"`
	AddressStreet   string                 `protobuf:"bytes,11,opt,name=addressStreet,proto3" json:"addressStreet,omitempty"`
	AddressNumber   string                 `protobuf:"bytes,12,opt,name=addressNumber,proto3" json:"addressNumber,omitempty"`
	AddressZipCode  string                 `protobuf:"bytes,13,opt,name=addressZipCode,proto3" json:"addressZipCode,omitempty"`
	AddressUnit     *string                `protobuf:"bytes,14,opt,name=addressUnit,proto3,oneof" json:"addressUnit,omitempty"`
	AddressLat      *float64               `protobuf:"fixed64,15,opt,name=addressLat,proto3,oneof" json:"addressLat,omitempty"`
	AddressLng      *float64               `protobuf:"fixed64,16,opt,name=addressLng,proto3,oneof" json:"addressLng,omitempty"`
	FacilityIds     []int64                `protobuf:"varint,17,rep,packed,name=facilityIds,proto3" json:"facilityIds,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *CreateEmployeeRequest) Reset() {
	*x = CreateEmployeeRequest{}
	mi := &file_employee_service_requests_employee_rpc_create_employee_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateEmployeeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmployeeRequest) ProtoMessage() {}

func (x *CreateEmployeeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_employee_service_requests_employee_rpc_create_employee_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmployeeRequest.ProtoReflect.Descriptor instead.
func (*CreateEmployeeRequest) Descriptor() ([]byte, []int) {
	return file_employee_service_requests_employee_rpc_create_employee_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEmployeeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateEmployeeRequest) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *CreateEmployeeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateEmployeeRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *CreateEmployeeRequest) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *CreateEmployeeRequest) GetRealId() string {
	if x != nil {
		return x.RealId
	}
	return ""
}

func (x *CreateEmployeeRequest) GetFiscalId() string {
	if x != nil {
		return x.FiscalId
	}
	return ""
}

func (x *CreateEmployeeRequest) GetAddressCountry() string {
	if x != nil {
		return x.AddressCountry
	}
	return ""
}

func (x *CreateEmployeeRequest) GetAddressState() string {
	if x != nil {
		return x.AddressState
	}
	return ""
}

func (x *CreateEmployeeRequest) GetAddressSubState() string {
	if x != nil && x.AddressSubState != nil {
		return *x.AddressSubState
	}
	return ""
}

func (x *CreateEmployeeRequest) GetAddressStreet() string {
	if x != nil {
		return x.AddressStreet
	}
	return ""
}

func (x *CreateEmployeeRequest) GetAddressNumber() string {
	if x != nil {
		return x.AddressNumber
	}
	return ""
}

func (x *CreateEmployeeRequest) GetAddressZipCode() string {
	if x != nil {
		return x.AddressZipCode
	}
	return ""
}

func (x *CreateEmployeeRequest) GetAddressUnit() string {
	if x != nil && x.AddressUnit != nil {
		return *x.AddressUnit
	}
	return ""
}

func (x *CreateEmployeeRequest) GetAddressLat() float64 {
	if x != nil && x.AddressLat != nil {
		return *x.AddressLat
	}
	return 0
}

func (x *CreateEmployeeRequest) GetAddressLng() float64 {
	if x != nil && x.AddressLng != nil {
		return *x.AddressLng
	}
	return 0
}

func (x *CreateEmployeeRequest) GetFacilityIds() []int64 {
	if x != nil {
		return x.FacilityIds
	}
	return nil
}

type CreateEmployeeResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Employee      *models.Employee       `protobuf:"bytes,1,opt,name=employee,proto3" json:"employee,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateEmployeeResponse) Reset() {
	*x = CreateEmployeeResponse{}
	mi := &file_employee_service_requests_employee_rpc_create_employee_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateEmployeeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEmployeeResponse) ProtoMessage() {}

func (x *CreateEmployeeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_employee_service_requests_employee_rpc_create_employee_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEmployeeResponse.ProtoReflect.Descriptor instead.
func (*CreateEmployeeResponse) Descriptor() ([]byte, []int) {
	return file_employee_service_requests_employee_rpc_create_employee_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEmployeeResponse) GetEmployee() *models.Employee {
	if x != nil {
		return x.Employee
	}
	return nil
}

var File_employee_service_requests_employee_rpc_create_employee_proto protoreflect.FileDescriptor

const file_employee_service_requests_employee_rpc_create_employee_proto_rawDesc = "" +
	"\n" +
	"<employee-service/requests/employee/rpc_create_employee.proto\x122employee_service.requests.employee.create_employee\x1a&employee-service/models/employee.proto\"\x83\x05\n" +
	"\x15CreateEmployeeRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1a\n" +
	"\blastname\x18\x02 \x01(\tR\blastname\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x14\n" +
	"\x05phone\x18\x04 \x01(\tR\x05phone\x12\x16\n" +
	"\x06gender\x18\x05 \x01(\tR\x06gender\x12\x16\n" +
	"\x06realId\x18\x06 \x01(\tR\x06realId\x12\x1a\n" +
	"\bfiscalId\x18\a \x01(\tR\bfiscalId\x12&\n" +
	"\x0eaddressCountry\x18\b \x01(\tR\x0eaddressCountry\x12\"\n" +
	"\faddressState\x18\t \x01(\tR\faddressState\x12-\n" +
	"\x0faddressSubState\x18\n" +
	" \x01(\tH\x00R\x0faddressSubState\x88\x01\x01\x12$\n" +
	"\raddressStreet\x18\v \x01(\tR\raddressStreet\x12$\n" +
	"\raddressNumber\x18\f \x01(\tR\raddressNumber\x12&\n" +
	"\x0eaddressZipCode\x18\r \x01(\tR\x0eaddressZipCode\x12%\n" +
	"\vaddressUnit\x18\x0e \x01(\tH\x01R\vaddressUnit\x88\x01\x01\x12#\n" +
	"\n" +
	"addressLat\x18\x0f \x01(\x01H\x02R\n" +
	"addressLat\x88\x01\x01\x12#\n" +
	"\n" +
	"addressLng\x18\x10 \x01(\x01H\x03R\n" +
	"addressLng\x88\x01\x01\x12 \n" +
	"\vfacilityIds\x18\x11 \x03(\x03R\vfacilityIdsB\x12\n" +
	"\x10_addressSubStateB\x0e\n" +
	"\f_addressUnitB\r\n" +
	"\v_addressLatB\r\n" +
	"\v_addressLng\"`\n" +
	"\x16CreateEmployeeResponse\x12F\n" +
	"\bemployee\x18\x01 \x01(\v2*.employee_service.models.employee.EmployeeR\bemployeeBPZNgithub.com/nicodanke/gesty-api/shared/proto/employee-service/requests/employeeb\x06proto3"

var (
	file_employee_service_requests_employee_rpc_create_employee_proto_rawDescOnce sync.Once
	file_employee_service_requests_employee_rpc_create_employee_proto_rawDescData []byte
)

func file_employee_service_requests_employee_rpc_create_employee_proto_rawDescGZIP() []byte {
	file_employee_service_requests_employee_rpc_create_employee_proto_rawDescOnce.Do(func() {
		file_employee_service_requests_employee_rpc_create_employee_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_employee_service_requests_employee_rpc_create_employee_proto_rawDesc), len(file_employee_service_requests_employee_rpc_create_employee_proto_rawDesc)))
	})
	return file_employee_service_requests_employee_rpc_create_employee_proto_rawDescData
}

var file_employee_service_requests_employee_rpc_create_employee_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_employee_service_requests_employee_rpc_create_employee_proto_goTypes = []any{
	(*CreateEmployeeRequest)(nil),  // 0: employee_service.requests.employee.create_employee.CreateEmployeeRequest
	(*CreateEmployeeResponse)(nil), // 1: employee_service.requests.employee.create_employee.CreateEmployeeResponse
	(*models.Employee)(nil),        // 2: employee_service.models.employee.Employee
}
var file_employee_service_requests_employee_rpc_create_employee_proto_depIdxs = []int32{
	2, // 0: employee_service.requests.employee.create_employee.CreateEmployeeResponse.employee:type_name -> employee_service.models.employee.Employee
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_employee_service_requests_employee_rpc_create_employee_proto_init() }
func file_employee_service_requests_employee_rpc_create_employee_proto_init() {
	if File_employee_service_requests_employee_rpc_create_employee_proto != nil {
		return
	}
	file_employee_service_requests_employee_rpc_create_employee_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_employee_service_requests_employee_rpc_create_employee_proto_rawDesc), len(file_employee_service_requests_employee_rpc_create_employee_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_employee_service_requests_employee_rpc_create_employee_proto_goTypes,
		DependencyIndexes: file_employee_service_requests_employee_rpc_create_employee_proto_depIdxs,
		MessageInfos:      file_employee_service_requests_employee_rpc_create_employee_proto_msgTypes,
	}.Build()
	File_employee_service_requests_employee_rpc_create_employee_proto = out.File
	file_employee_service_requests_employee_rpc_create_employee_proto_goTypes = nil
	file_employee_service_requests_employee_rpc_create_employee_proto_depIdxs = nil
}
