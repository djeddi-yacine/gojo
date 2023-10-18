package utils

import (
	"fmt"
	"net/mail"
	"regexp"
	"time"

	"github.com/dj-yacine-flutter/gojo/pb"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidToken    = regexp.MustCompile(`^v2\.[^.]*\.[^.]*\.[^.]*$`).MatchString
)

func ValidateString(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("must contain from %d-%d characters", minLength, maxLength)
	}
	return nil
}

func ValidateInt(value int64) error {
	if value <= 0 {
		return fmt.Errorf("%d must be a unsigned number", value)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("must contain only lowercase letters, digits, or underscore")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateString(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("must contain only letters or spaces")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateString(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateString(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("is not a valid email address")
	}
	return nil
}

func ValidateEmailID(value int64) error {
	if value <= 0 {
		return fmt.Errorf("must be a positive integer")
	}
	return nil
}

func ValidateSecretCode(value string) error {
	return ValidateString(value, 32, 128)
}

func ValidateDuration(value string) error {
	_, err := time.ParseDuration(value)
	return err
}

func ValidateYear(value int32) error {
	if value < 1900 || value > 2100 {
		return fmt.Errorf("invalid year number")
	}
	return nil
}

func ValidateDate(input string) error {
	_, err := time.Parse(time.DateOnly, input)
	return err
}

func ValidateToken(value string) error {
	if !isValidToken(value) {
		return fmt.Errorf("must be a valid token")

	}
	return nil
}

func ValidateGenreAndStudio(values []string) error {
	for _, value := range values {
		if err := ValidateString(value, 2, 50); err != nil {
			return err
		}
	}
	return nil
}

func ValidateLanguage(values []*pb.LanguageRequest) error {
	for _, value := range values {
		if err := ValidateString(value.LanguageCode, 2, 3); err != nil {
			return err
		}
		if err := ValidateString(value.LanguageName, 2, 30); err != nil {
			return err
		}
	}
	return nil
}
