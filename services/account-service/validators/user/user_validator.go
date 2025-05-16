package userValidator

import (
	"fmt"
	"regexp"

	v "github.com/nicodanke/gesty-api/services/account-service/validators"
)

var (
	isValidUsername     = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`).MatchString
	isValidName         = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidLastname     = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidFullUsername = regexp.MustCompile(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9]+$`).MatchString
	isValidPhone        = regexp.MustCompile(`^[+\-0-9\s]+$`).MatchString
)

func ValidateUsername(value string) error {
	err := v.ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("username only can contain lower case letters, diggits or underscore")
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

func ValidateFullUsername(value string) error {
	err := v.ValidString(value, 3, 100)
	if err != nil {
		return err
	}
	if !isValidFullUsername(value) {
		return fmt.Errorf("username is invalid")
	}
	return nil
}

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
		return fmt.Errorf("fullname only can contain letters or spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	return v.ValidString(value, 8, 256)
}