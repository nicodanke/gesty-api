package eventdata

type Employee struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	Lastname        string   `json:"lastname"`
	Email           string   `json:"email"`
	Phone           string   `json:"phone"`
	Gender          string   `json:"gender"`
	RealId          string   `json:"real_id"`
	FiscalId        string   `json:"fiscal_id"`
	AddressCountry  string   `json:"address_country"`
	AddressState    string   `json:"address_state"`
	AddressSubState string   `json:"address_sub_state"`
	AddressStreet   string   `json:"address_street"`
	AddressNumber   string   `json:"address_number"`
	AddressUnit     string   `json:"address_unit"`
	AddressZipCode  string   `json:"address_ZipCode"`
	AddressLat      float64  `json:"address_lat"`
	AddressLng      float64  `json:"address_lng"`
	FacilityIds     []string `json:"facility_ids"`
}
