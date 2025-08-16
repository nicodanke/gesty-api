package eventdata

import (
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Attendance struct {
	Id           string                 `json:"id"`
	TimeIn       *timestamppb.Timestamp `json:"timeIn"`
	EmployeeId   int64                  `json:"employeeId"`
	EmployeeName string                 `json:"employeeName"`
	ActionId     int64                  `json:"actionId"`
	ActionName   string                 `json:"actionName"`
	DeviceId     int64                  `json:"deviceId"`
	DeviceName   string                 `json:"deviceName"`
	Precision    float64                `json:"precision"`
}
