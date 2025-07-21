package eventdata

import "google.golang.org/protobuf/types/known/durationpb"

type Facility struct {
	Id              string               `json:"id"`
	Name            string               `json:"name"`
	Description     string               `json:"description"`
	OpenTime        *durationpb.Duration `json:"openTime"`
	CloseTime       *durationpb.Duration `json:"closeTime"`
	AddressCountry  string               `json:"addressCountry"`
	AddressState    string               `json:"addressState"`
	AddressSubState string               `json:"addressSubState"`
	AddressStreet   string               `json:"addressStreet"`
	AddressNumber   string               `json:"addressNumber"`
	AddressUnit     string               `json:"addressUnit"`
	AddressZipCode  string               `json:"addressZipCode"`
	AddressLat      float64              `json:"addressLat"`
	AddressLng      float64              `json:"addressLng"`
}
