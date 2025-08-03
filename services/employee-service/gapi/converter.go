package gapi

import (
	"strconv"
	"time"

	db "github.com/nicodanke/gesty-api/services/employee-service/db/sqlc"
	"github.com/nicodanke/gesty-api/services/employee-service/sse/eventdata"
	"github.com/nicodanke/gesty-api/shared/proto/employee-service/models"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertAction(action db.Action) *models.Action {
	return &models.Action{
		Id:           action.ID,
		Name:         action.Name,
		Description:  action.Description.String,
		Enabled:      action.Enabled,
		CanBeDeleted: action.CanBeDeleted,
	}
}

func convertActionEvent(action db.Action) *eventdata.Action {
	return &eventdata.Action{
		Id:           strconv.FormatInt(action.ID, 10),
		Name:         action.Name,
		Description:  action.Description.String,
		Enabled:      action.Enabled,
		CanBeDeleted: action.CanBeDeleted,
	}
}

func convertActions(actions []db.Action) []*models.Action {
	result := make([]*models.Action, len(actions))

	for i, v := range actions {
		result[i] = convertAction(v)
	}

	return result
}

func convertFacilitiesGetRows(facilities []db.GetFacilitiesRow) []*models.Facility {
	result := make([]*models.Facility, len(facilities))

	for i, v := range facilities {
		result[i] = convertFacilitiesGetRow(v)
	}

	return result
}

func convertFacilitiesGetRow(facility db.GetFacilitiesRow) *models.Facility {
	return &models.Facility{
		Id:              facility.ID,
		Name:            facility.Name,
		Description:     facility.Description.String,
		OpenTime:        durationpb.New(time.Duration(facility.OpenTime.Microseconds)),
		CloseTime:       durationpb.New(time.Duration(facility.CloseTime.Microseconds)),
		AddressCountry:  facility.Country.String,
		AddressState:    facility.State.String,
		AddressSubState: facility.SubState.String,
		AddressStreet:   facility.Street.String,
		AddressNumber:   facility.Number.String,
		AddressUnit:     facility.Unit.String,
		AddressZipCode:  facility.ZipCode.String,
		AddressLat:      facility.Lat.Float64,
		AddressLng:      facility.Lng.Float64,
	}
}

func convertFacilityGetRow(facility db.GetFacilityRow) *models.Facility {
	return &models.Facility{
		Id:              facility.ID,
		Name:            facility.Name,
		Description:     facility.Description.String,
		OpenTime:        durationpb.New(time.Duration(facility.OpenTime.Microseconds)),
		CloseTime:       durationpb.New(time.Duration(facility.CloseTime.Microseconds)),
		AddressCountry:  facility.Country.String,
		AddressState:    facility.State.String,
		AddressSubState: facility.SubState.String,
		AddressStreet:   facility.Street.String,
		AddressNumber:   facility.Number.String,
		AddressUnit:     facility.Unit.String,
		AddressZipCode:  facility.ZipCode.String,
		AddressLat:      facility.Lat.Float64,
		AddressLng:      facility.Lng.Float64,
	}
}

func convertFacilityCreateTxResult(facility db.CreateFacilityTxResult) *models.Facility {
	return &models.Facility{
		Id:              facility.Facility.ID,
		Name:            facility.Facility.Name,
		Description:     facility.Facility.Description.String,
		OpenTime:        durationpb.New(time.Duration(facility.Facility.OpenTime.Microseconds)),
		CloseTime:       durationpb.New(time.Duration(facility.Facility.CloseTime.Microseconds)),
		AddressCountry:  facility.FacilityAddress.Country,
		AddressState:    facility.FacilityAddress.State,
		AddressSubState: facility.FacilityAddress.SubState.String,
		AddressStreet:   facility.FacilityAddress.Street,
		AddressNumber:   facility.FacilityAddress.Number,
		AddressUnit:     facility.FacilityAddress.Unit.String,
		AddressZipCode:  facility.FacilityAddress.ZipCode,
		AddressLat:      facility.FacilityAddress.Lat.Float64,
		AddressLng:      facility.FacilityAddress.Lng.Float64,
	}
}

func convertFacilityUpdateTxResult(facility db.UpdateFacilityTxResult) *models.Facility {
	return &models.Facility{
		Id:              facility.Facility.ID,
		Name:            facility.Facility.Name,
		Description:     facility.Facility.Description.String,
		OpenTime:        durationpb.New(time.Duration(facility.Facility.OpenTime.Microseconds)),
		CloseTime:       durationpb.New(time.Duration(facility.Facility.CloseTime.Microseconds)),
		AddressCountry:  facility.FacilityAddress.Country,
		AddressState:    facility.FacilityAddress.State,
		AddressSubState: facility.FacilityAddress.SubState.String,
		AddressStreet:   facility.FacilityAddress.Street,
		AddressNumber:   facility.FacilityAddress.Number,
		AddressUnit:     facility.FacilityAddress.Unit.String,
		AddressZipCode:  facility.FacilityAddress.ZipCode,
		AddressLat:      facility.FacilityAddress.Lat.Float64,
		AddressLng:      facility.FacilityAddress.Lng.Float64,
	}
}

func convertCreateFacilityTxResultEvent(facility db.CreateFacilityTxResult) *eventdata.Facility {
	return &eventdata.Facility{
		Id:              strconv.FormatInt(facility.Facility.ID, 10),
		Name:            facility.Facility.Name,
		Description:     facility.Facility.Description.String,
		OpenTime:        durationpb.New(time.Duration(facility.Facility.OpenTime.Microseconds)),
		CloseTime:       durationpb.New(time.Duration(facility.Facility.CloseTime.Microseconds)),
		AddressCountry:  facility.FacilityAddress.Country,
		AddressState:    facility.FacilityAddress.State,
		AddressSubState: facility.FacilityAddress.SubState.String,
		AddressStreet:   facility.FacilityAddress.Street,
		AddressNumber:   facility.FacilityAddress.Number,
		AddressUnit:     facility.FacilityAddress.Unit.String,
		AddressZipCode:  facility.FacilityAddress.ZipCode,
		AddressLat:      facility.FacilityAddress.Lat.Float64,
		AddressLng:      facility.FacilityAddress.Lng.Float64,
	}
}

func convertEmployeeCreateTxResult(employee db.CreateEmployeeTxResult) *models.Employee {
	return &models.Employee{
		Id:              employee.Employee.ID,
		Name:            employee.Employee.Name,
		Lastname:        employee.Employee.Lastname,
		Email:           employee.Employee.Email,
		Phone:           employee.Employee.Phone,
		Gender:          employee.Employee.Gender,
		RealId:          employee.Employee.RealID,
		FiscalId:        employee.Employee.FiscalID,
		AddressCountry:  employee.EmployeeAddress.Country,
		AddressState:    employee.EmployeeAddress.State,
		AddressSubState: employee.EmployeeAddress.SubState.String,
		AddressStreet:   employee.EmployeeAddress.Street,
		AddressNumber:   employee.EmployeeAddress.Number,
		AddressUnit:     employee.EmployeeAddress.Unit.String,
		AddressZipCode:  employee.EmployeeAddress.ZipCode,
		AddressLat:      employee.EmployeeAddress.Lat.Float64,
		AddressLng:      employee.EmployeeAddress.Lng.Float64,
		FacilityIds:     employee.FacilityIDs,
	}
}

func convertEmployeeCreateTxResultEvent(employee db.CreateEmployeeTxResult) *eventdata.Employee {
	facilityIds := make([]string, len(employee.FacilityIDs))
	for i, id := range employee.FacilityIDs {
		facilityIds[i] = strconv.FormatInt(id, 10)
	}

	return &eventdata.Employee{
		Id:              strconv.FormatInt(employee.Employee.ID, 10),
		Name:            employee.Employee.Name,
		Lastname:        employee.Employee.Lastname,
		Email:           employee.Employee.Email,
		Phone:           employee.Employee.Phone,
		Gender:          employee.Employee.Gender,
		RealId:          employee.Employee.RealID,
		FiscalId:        employee.Employee.FiscalID,
		AddressCountry:  employee.EmployeeAddress.Country,
		AddressState:    employee.EmployeeAddress.State,
		AddressSubState: employee.EmployeeAddress.SubState.String,
		AddressStreet:   employee.EmployeeAddress.Street,
		AddressNumber:   employee.EmployeeAddress.Number,
		AddressUnit:     employee.EmployeeAddress.Unit.String,
		AddressZipCode:  employee.EmployeeAddress.ZipCode,
		AddressLat:      employee.EmployeeAddress.Lat.Float64,
		AddressLng:      employee.EmployeeAddress.Lng.Float64,
		FacilityIds:     facilityIds,
	}
}

func convertEmployeesGetRows(employees []db.GetEmployeesRow) []*models.Employee {
	result := make([]*models.Employee, len(employees))

	for i, v := range employees {
		result[i] = convertEmployeesGetRow(v)
	}

	return result
}

func convertEmployeesGetRow(employee db.GetEmployeesRow) *models.Employee {
	facilityIds := make([]int64, 0)
	for _, v := range employee.FacilityIds.([]interface{}) {
		facilityIds = append(facilityIds, v.(int64))
	}

	return &models.Employee{
		Id:              employee.ID,
		Name:            employee.Name,
		Lastname:        employee.Lastname,
		Email:           employee.Email,
		Phone:           employee.Phone,
		Gender:          employee.Gender,
		RealId:          employee.RealID,
		FiscalId:        employee.FiscalID,
		AddressCountry:  employee.Country.String,
		AddressState:    employee.State.String,
		AddressSubState: employee.SubState.String,
		AddressStreet:   employee.Street.String,
		AddressNumber:   employee.Number.String,
		AddressUnit:     employee.Unit.String,
		AddressZipCode:  employee.ZipCode.String,
		AddressLat:      employee.Lat.Float64,
		AddressLng:      employee.Lng.Float64,
		FacilityIds:     facilityIds,
	}
}

func convertEmployeeGetRow(employee db.GetEmployeeRow) *models.Employee {
	facilityIds := make([]int64, 0)
	for _, v := range employee.FacilityIds.([]interface{}) {
		facilityIds = append(facilityIds, v.(int64))
	}

	return &models.Employee{
		Id:              employee.ID,
		Name:            employee.Name,
		Lastname:        employee.Lastname,
		Email:           employee.Email,
		Phone:           employee.Phone,
		Gender:          employee.Gender,
		RealId:          employee.RealID,
		FiscalId:        employee.FiscalID,
		AddressCountry:  employee.Country.String,
		AddressState:    employee.State.String,
		AddressSubState: employee.SubState.String,
		AddressStreet:   employee.Street.String,
		AddressNumber:   employee.Number.String,
		AddressUnit:     employee.Unit.String,
		AddressZipCode:  employee.ZipCode.String,
		AddressLat:      employee.Lat.Float64,
		AddressLng:      employee.Lng.Float64,
		FacilityIds:     facilityIds,
	}
}

func convertEmployeeUpdateTxResult(employee db.UpdateEmployeeTxResult) *models.Employee {
	return &models.Employee{
		Id:              employee.Employee.ID,
		Name:            employee.Employee.Name,
		Lastname:        employee.Employee.Lastname,
		Email:           employee.Employee.Email,
		Phone:           employee.Employee.Phone,
		Gender:          employee.Employee.Gender,
		RealId:          employee.Employee.RealID,
		FiscalId:        employee.Employee.FiscalID,
		AddressCountry:  employee.EmployeeAddress.Country,
		AddressState:    employee.EmployeeAddress.State,
		AddressSubState: employee.EmployeeAddress.SubState.String,
		AddressStreet:   employee.EmployeeAddress.Street,
		AddressNumber:   employee.EmployeeAddress.Number,
		AddressUnit:     employee.EmployeeAddress.Unit.String,
		AddressZipCode:  employee.EmployeeAddress.ZipCode,
		AddressLat:      employee.EmployeeAddress.Lat.Float64,
		AddressLng:      employee.EmployeeAddress.Lng.Float64,
		FacilityIds:     employee.FacilityIds,
	}
}

func convertEmployeeUpdateTxResultEvent(employee db.UpdateEmployeeTxResult) *eventdata.Employee {
	facilityIds := make([]string, len(employee.FacilityIds))
	for i, id := range employee.FacilityIds {
		facilityIds[i] = strconv.FormatInt(id, 10)
	}

	return &eventdata.Employee{
		Id:              strconv.FormatInt(employee.Employee.ID, 10),
		Name:            employee.Employee.Name,
		Lastname:        employee.Employee.Lastname,
		Email:           employee.Employee.Email,
		Phone:           employee.Employee.Phone,
		Gender:          employee.Employee.Gender,
		RealId:          employee.Employee.RealID,
		FiscalId:        employee.Employee.FiscalID,
		AddressCountry:  employee.EmployeeAddress.Country,
		AddressState:    employee.EmployeeAddress.State,
		AddressSubState: employee.EmployeeAddress.SubState.String,
		AddressStreet:   employee.EmployeeAddress.Street,
		AddressNumber:   employee.EmployeeAddress.Number,
		AddressUnit:     employee.EmployeeAddress.Unit.String,
		AddressZipCode:  employee.EmployeeAddress.ZipCode,
		AddressLat:      employee.EmployeeAddress.Lat.Float64,
		AddressLng:      employee.EmployeeAddress.Lng.Float64,
		FacilityIds:     facilityIds,
	}
}

func convertDeviceCreateTxResult(device db.CreateDeviceTxResult) *models.Device {
	return &models.Device{
		Id:         device.Device.ID,
		Name:       device.Device.Name,
		Enabled:    device.Device.Enabled,
		Active:     device.Device.Active,
		Password:   device.Device.Password,
		ActionIds:  device.ActionIDs,
		FacilityId: device.Device.FacilityID,
	}
}

func convertDeviceCreateTxResultEvent(device db.CreateDeviceTxResult) *eventdata.Device {
	actionIds := make([]string, len(device.ActionIDs))
	for i, id := range device.ActionIDs {
		actionIds[i] = strconv.FormatInt(id, 10)
	}

	return &eventdata.Device{
		Id:         strconv.FormatInt(device.Device.ID, 10),
		Name:       device.Device.Name,
		Enabled:    device.Device.Enabled,
		Active:     device.Device.Active,
		Password:   device.Device.Password,
		ActionIds:  actionIds,
		FacilityId: strconv.FormatInt(device.Device.FacilityID, 10),
	}
}

func convertGetDeviceRow(device db.GetDeviceRow) *models.Device {
	actionIds := make([]int64, 0)
	for _, v := range device.ActionIds.([]interface{}) {
		actionIds = append(actionIds, v.(int64))
	}

	return &models.Device{
		Id:                      device.ID,
		Name:                    device.Name,
		Enabled:                 device.Enabled,
		Active:                  device.Active,
		FacilityId:              device.FacilityID,
		Password:                device.Password,
		ActionIds:               actionIds,
		DeviceName:              device.DeviceName.String,
		DeviceModel:             device.DeviceModel.String,
		DeviceBrand:             device.DeviceBrand.String,
		DeviceSerialNumber:      device.DeviceSerialNumber.String,
		DeviceOs:                device.DeviceOs.String,
		DeviceRam:               strconv.FormatFloat(device.DeviceRam.Float64, 'f', -1, 64),
		DeviceStorage:           strconv.FormatFloat(device.DeviceStorage.Float64, 'f', -1, 64),
		DeviceOsVersion:         device.DeviceOsVersion.String,
		ActivationCode:          device.ActivationCode.String,
		ActivationCodeExpiresAt: timestamppb.New(device.ActivationCodeExpiresAt),
	}
}

func convertGetDeviceRowEvent(device db.GetDeviceRow) *eventdata.Device {
	actionIds := make([]string, 0)
	for _, v := range device.ActionIds.([]interface{}) {
		actionIds = append(actionIds, strconv.FormatInt(v.(int64), 10))
	}

	return &eventdata.Device{
		Id:                      strconv.FormatInt(device.ID, 10),
		Name:                    device.Name,
		Enabled:                 device.Enabled,
		Active:                  device.Active,
		FacilityId:              strconv.FormatInt(device.FacilityID, 10),
		Password:                device.Password,
		ActionIds:               actionIds,
		DeviceName:              device.DeviceName.String,
		DeviceModel:             device.DeviceModel.String,
		DeviceBrand:             device.DeviceBrand.String,
		DeviceSerialNumber:      device.DeviceSerialNumber.String,
		DeviceOs:                device.DeviceOs.String,
		DeviceRam:               strconv.FormatFloat(device.DeviceRam.Float64, 'f', -1, 64),
		DeviceStorage:           strconv.FormatFloat(device.DeviceStorage.Float64, 'f', -1, 64),
		DeviceOsVersion:         device.DeviceOsVersion.String,
		ActivationCode:          device.ActivationCode.String,
		ActivationCodeExpiresAt: timestamppb.New(device.ActivationCodeExpiresAt),
	}
}

func convertEmployeePhoto(employeePhotos []db.EmployeePhoto) []*models.EmployeeImage {
	result := make([]*models.EmployeeImage, len(employeePhotos))

	for i, v := range employeePhotos {
		result[i] = convertEmployeePhotosGetRow(v)
	}

	return result
}

func convertEmployeePhotosGetRow(employeePhoto db.EmployeePhoto) *models.EmployeeImage {
	return &models.EmployeeImage{
		Id:          employeePhoto.ID,
		ImageBase64: employeePhoto.ImageBase64,
	}
}

func convertEmployeePhotoTxResult(employeePhoto db.CreateEmployeePhotoTxResult) *models.EmployeeImage {
	return &models.EmployeeImage{
		Id:          employeePhoto.EmployeePhoto.ID,
		ImageBase64: employeePhoto.EmployeePhoto.ImageBase64,
	}
}

func convertGetDevicesRows(devices []db.GetDevicesRow) []*models.Device {
	result := make([]*models.Device, len(devices))

	for i, v := range devices {
		result[i] = convertGetDevicesRow(v)
	}

	return result
}

func convertGetDevicesRow(device db.GetDevicesRow) *models.Device {
	actionIds := make([]int64, 0)
	for _, v := range device.ActionIds.([]interface{}) {
		actionIds = append(actionIds, v.(int64))
	}

	return &models.Device{
		Id:                      device.ID,
		Name:                    device.Name,
		Enabled:                 device.Enabled,
		Active:                  device.Active,
		FacilityId:              device.FacilityID,
		Password:                device.Password,
		ActionIds:               actionIds,
		DeviceName:              device.DeviceName.String,
		DeviceModel:             device.DeviceModel.String,
		DeviceBrand:             device.DeviceBrand.String,
		DeviceSerialNumber:      device.DeviceSerialNumber.String,
		DeviceOs:                device.DeviceOs.String,
		DeviceRam:               strconv.FormatFloat(device.DeviceRam.Float64, 'f', -1, 64),
		DeviceStorage:           strconv.FormatFloat(device.DeviceStorage.Float64, 'f', -1, 64),
		DeviceOsVersion:         device.DeviceOsVersion.String,
		ActivationCode:          device.ActivationCode.String,
		ActivationCodeExpiresAt: timestamppb.New(device.ActivationCodeExpiresAt),
	}
}

func convertDeviceUpdateTxResult(device db.UpdateDeviceTxResult) *models.Device {
	return &models.Device{
		Id:                      device.Device.ID,
		Name:                    device.Device.Name,
		Enabled:                 device.Device.Enabled,
		Active:                  device.Device.Active,
		ActionIds:               device.ActionIDs,
		FacilityId:              device.Device.FacilityID,
		Password:                device.Device.Password,
		DeviceName:              device.Device.DeviceName.String,
		DeviceModel:             device.Device.DeviceModel.String,
		DeviceBrand:             device.Device.DeviceBrand.String,
		DeviceSerialNumber:      device.Device.DeviceSerialNumber.String,
		DeviceOs:                device.Device.DeviceOs.String,
		DeviceRam:               strconv.FormatFloat(device.Device.DeviceRam.Float64, 'f', -1, 64),
		DeviceStorage:           strconv.FormatFloat(device.Device.DeviceStorage.Float64, 'f', -1, 64),
		DeviceOsVersion:         device.Device.DeviceOsVersion.String,
		ActivationCode:          device.Device.ActivationCode.String,
		ActivationCodeExpiresAt: timestamppb.New(device.Device.ActivationCodeExpiresAt),
	}
}

func convertDeviceUpdateTxResultEvent(device db.UpdateDeviceTxResult) *eventdata.Device {
	actionIds := make([]string, len(device.ActionIDs))
	for i, id := range device.ActionIDs {
		actionIds[i] = strconv.FormatInt(id, 10)
	}

	return &eventdata.Device{
		Id:                      strconv.FormatInt(device.Device.ID, 10),
		Name:                    device.Device.Name,
		Enabled:                 device.Device.Enabled,
		Active:                  device.Device.Active,
		Password:                device.Device.Password,
		ActionIds:               actionIds,
		FacilityId:              strconv.FormatInt(device.Device.FacilityID, 10),
		DeviceName:              device.Device.DeviceName.String,
		DeviceModel:             device.Device.DeviceModel.String,
		DeviceBrand:             device.Device.DeviceBrand.String,
		DeviceSerialNumber:      device.Device.DeviceSerialNumber.String,
		DeviceOs:                device.Device.DeviceOs.String,
		DeviceRam:               strconv.FormatFloat(device.Device.DeviceRam.Float64, 'f', -1, 64),
		DeviceStorage:           strconv.FormatFloat(device.Device.DeviceStorage.Float64, 'f', -1, 64),
		DeviceOsVersion:         device.Device.DeviceOsVersion.String,
		ActivationCode:          device.Device.ActivationCode.String,
		ActivationCodeExpiresAt: timestamppb.New(device.Device.ActivationCodeExpiresAt),
	}
}
