package eventdata

type Employee struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	Lastname        string   `json:"lastname"`
	Email           string   `json:"email"`
	Phone           string   `json:"phone"`
	Gender          string   `json:"gender"`
	RealId          string   `json:"realId"`
	FiscalId        string   `json:"fiscalId"`
	AddressCountry  string   `json:"addressCountry"`
	AddressState    string   `json:"addressState"`
	AddressSubState string   `json:"addressSubState"`
	AddressStreet   string   `json:"addressStreet"`
	AddressNumber   string   `json:"addressNumber"`
	AddressUnit     string   `json:"addressUnit"`
	AddressZipCode  string   `json:"addressZipCode"`
	AddressLat      float64  `json:"addressLat"`
	AddressLng      float64  `json:"addressLng"`
	FacilityIds     []string `json:"facilityIds"`
}
