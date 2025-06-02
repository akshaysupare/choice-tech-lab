package util

import (
	"fmt"
	"regexp"
)

func ValidateEmail(email string) error {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !re.MatchString(email) {
		return fmt.Errorf("invalid email format")
	}
	return nil
}

func ValidateName(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("name cannot be empty")
	}
	return nil
}

func ValidateAge(age int) error {
	if age < 0 || age > 120 {
		return fmt.Errorf("age must be between 0 and 120")
	}
	return nil
}
