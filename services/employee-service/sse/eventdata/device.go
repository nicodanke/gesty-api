package eventdata

import "google.golang.org/protobuf/types/known/timestamppb"

type Device struct {
	Id                      string                 `json:"id"`
	Name                    string                 `json:"name"`
	Enabled                 bool                   `json:"enabled"`
	ActionIds               []string               `json:"action_ids"`
	FacilityId              string                 `json:"facility_id"`
	Password                string                 `json:"password"`
	DeviceName              string                 `json:"device_name"`
	DeviceModel             string                 `json:"device_model"`
	DeviceBrand             string                 `json:"device_brand"`
	DeviceSerialNumber      string                 `json:"device_serial_number"`
	DeviceOs                string                 `json:"device_os"`
	DeviceRam               string                 `json:"device_ram"`
	DeviceStorage           string                 `json:"device_storage"`
	DeviceOsVersion         string                 `json:"device_os_version"`
	ActivationCode          string                 `json:"activation_code"`
	ActivationCodeExpiresAt *timestamppb.Timestamp `json:"activation_code_expires_at"`
}
