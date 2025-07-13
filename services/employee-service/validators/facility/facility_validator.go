package roleValidator

import (
	"fmt"
	"regexp"

	v "github.com/nicodanke/gesty-api/services/employee-service/validators"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	isValidName = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`).MatchString
)

func ValidateName(value string) error {
	err := v.ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidName(value) {
		return fmt.Errorf("name only can contain letters or spaces")
	}
	return nil
}

func ValidateDescription(value string) error {
	if len(value) == 0 {
		return nil
	}
	return v.ValidString(value, 3, 1000)
}

func ValidateOpenTime(value *durationpb.Duration) error {
	if value.Seconds < 0 || value.Nanos < 0 {
		return fmt.Errorf("open time must be greater than 0")
	}
	if value.Seconds > 24*60*60 {
		return fmt.Errorf("open time must be less than 24 hours")
	}
	return nil
}

func ValidateCloseTime(value *durationpb.Duration) error {
	if value.Seconds < 0 || value.Nanos < 0 {
		return fmt.Errorf("close time must be greater than 0")
	}
	if value.Seconds > 24*60*60 {
		return fmt.Errorf("close time must be less than 24 hours")
	}
	return nil
}

func ValidateAddressCountry(value string) error {
	return v.ValidString(value, 3, 100)
}

func ValidateAddressState(value string) error {
	return v.ValidString(value, 3, 100)
}

func ValidateAddressSubState(value string) error {
	if len(value) == 0 {
		return nil
	}
	return v.ValidString(value, 3, 100)
}

func ValidateAddressStreet(value string) error {
	return v.ValidString(value, 3, 100)
}

func ValidateAddressNumber(value string) error {
	return v.ValidString(value, 3, 100)
}

func ValidateAddressUnit(value string) error {
	if len(value) == 0 {
		return nil
	}
	return v.ValidString(value, 3, 100)
}

func ValidateAddressPostalcode(value string) error {
	return v.ValidString(value, 3, 100)
}

func ValidateAddressLat(value float64) error {
	if value < -90 || value > 90 {
		return fmt.Errorf("address lat must be between -90 and 90")
	}
	return nil
}

func ValidateAddressLng(value float64) error {
	if value < -180 || value > 180 {
		return fmt.Errorf("address lng must be between -180 and 180")
	}
	return nil
}
