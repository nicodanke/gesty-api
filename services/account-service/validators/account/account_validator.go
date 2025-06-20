package accountValidator

import (
	"fmt"
	"regexp"

	v "github.com/nicodanke/gesty-api/services/account-service/validators"
)

var (
	isValidCompanyName = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`).MatchString
	isValidPhone       = regexp.MustCompile(`^[+\-0-9\s]+$`).MatchString
	isValidCountry     = regexp.MustCompile(`^ARG$`).MatchString
)

func ValidateCompanyName(value string) error {
	err := v.ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidCompanyName(value) {
		return fmt.Errorf("company name only can contain letters, diggits and spaces")
	}
	return nil
}

func ValidateWebUrl(value string) error {
	err := v.ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	return nil
}

func ValidatePhone(value string) error {
	err := v.ValidString(value, 6, 100)
	if err != nil {
		return err
	}
	if !isValidPhone(value) {
		return fmt.Errorf("phone can contain only numbers, + or -")
	}
	return nil
}
