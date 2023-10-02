package helpers

import (
	"errors"
	"regexp"
	"strings"
)

func ValidatePassword(password string) error {
	// Check length
	if len(password) < 6 {
		return errors.New("password should be at least 6 characters long")
	}

	// Check for special characters
	specialCharRegex := regexp.MustCompile(`[^a-zA-Z0-9]`)
	if !specialCharRegex.MatchString(password) {
		return errors.New("password should contain at least one special character")
	}

	// Check for lowercase and uppercase letters
	if strings.ToLower(password) == password || strings.ToUpper(password) == password {
		return errors.New("password should contain both lowercase and uppercase letters")
	}

	// Check for at least one number
	hasNumber := false
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasNumber = true
			break
		}
	}

	if !hasNumber {
		return errors.New("password should contain at least one number")
	}

	return nil
}
