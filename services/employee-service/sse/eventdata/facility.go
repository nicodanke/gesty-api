package eventdata

import "google.golang.org/protobuf/types/known/durationpb"

type Facility struct {
	Id                string               `json:"id"`
	Name              string               `json:"name"`
	Description       string               `json:"description"`
	OpenTime          *durationpb.Duration `json:"open_time"`
	CloseTime         *durationpb.Duration `json:"close_time"`
	AddressCountry    string               `json:"address_country"`
	AddressState      string               `json:"address_state"`
	AddressSubState   string               `json:"address_sub_state"`
	AddressStreet     string               `json:"address_street"`
	AddressNumber     string               `json:"address_number"`
	AddressUnit       string               `json:"address_unit"`
	AddressPostalcode string               `json:"address_postalcode"`
	AddressLat        float64              `json:"address_lat"`
	AddressLng        float64              `json:"address_lng"`
}
