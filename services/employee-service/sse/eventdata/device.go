package eventdata

import "google.golang.org/protobuf/types/known/timestamppb"

type Device struct {
	Id                      string                 `json:"id"`
	Name                    string                 `json:"name"`
	Enabled                 bool                   `json:"enabled"`
	Active                  bool                   `json:"active"`
	ActionIds               []string               `json:"actionIds"`
	FacilityId              string                 `json:"facilityId"`
	Password                string                 `json:"password"`
	DeviceName              string                 `json:"deviceName"`
	DeviceModel             string                 `json:"deviceModel"`
	DeviceBrand             string                 `json:"deviceBrand"`
	DeviceSerialNumber      string                 `json:"deviceSerialNumber"`
	DeviceOs                string                 `json:"deviceOs"`
	DeviceRam               string                 `json:"deviceRam"`
	DeviceStorage           string                 `json:"deviceStorage"`
	DeviceOsVersion         string                 `json:"deviceOsVersion"`
	ActivationCode          string                 `json:"activationCode"`
	ActivationCodeExpiresAt *timestamppb.Timestamp `json:"activationCodeExpiresAt"`
}
