package roleValidator

import (
	"fmt"
	"regexp"

	v "github.com/nicodanke/gesty-api/services/employee-service/validators"
	"google.golang.org/protobuf/types/known/durationpb"
)

var (
	isValidName         = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidLastname     = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidEmail        = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString
	isValidPhone        = regexp.MustCompile(`^[+\-0-9\s]+$`).MatchString
	isValidGender       = regexp.MustCompile(`^[MFX]$`).MatchString
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

func ValidateLastname(value string) error {
	err := v.ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidLastname(value) {
		return fmt.Errorf("lastname only can contain letters or spaces")
	}
	return nil
}

func ValidateEmail(value string) error {
	if !isValidEmail(value) {
		return fmt.Errorf("email is not valid")
	}
	return nil
}

func ValidatePhone(value string) error {
	if !isValidPhone(value) {
		return fmt.Errorf("phone is not valid")
	}
	return nil
}

func ValidateGender(value string) error {
	if !isValidGender(value) {
		return fmt.Errorf("gender is not valid")
	}
	return nil
}

func ValidateRealId(value string) error {
	return v.ValidString(value, 3, 100)
}

func ValidateFiscalId(value string) error {
	return v.ValidString(value, 3, 100)
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
