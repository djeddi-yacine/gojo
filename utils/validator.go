package utils

import (
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
	isValidToken    = regexp.MustCompile(`^v2\.[^.]*\.[^.]*\.[^.]*$`).MatchString
	isValidQuality  = regexp.MustCompile(`^(auto|1080p|720p|480p|480p|360p|240p|144p|SD|HD|FHD|4K)$`).MatchString
	isValidImage    = regexp.MustCompile(`(?i)\.(jpg|jpeg|png|gif|bmp|tiff|webp|svg|ico)$`).MatchString
	isValidMAC      = regexp.MustCompile(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`).MatchString
	showTypes       = []string{"movie", "tv", "movie special", "tv special", "ova", "ona", "compilation"}
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

func ValidateIMDbID(value string) error {
	if !strings.Contains(value, "tt") {
		return fmt.Errorf("must be a valid IMDb ID")

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

func ValidateURL(value, domain string) error {
	u, err := url.Parse(value)
	if err != nil {
		return fmt.Errorf("invalid URL: %w", err)
	}

	if domain != "" {
		if u.Scheme == "" || u.Host == "" {
			return fmt.Errorf("URL is missing scheme or host")
		}

		if domain != "" && !strings.Contains(u.Host, domain) {
			return fmt.Errorf("URL domain does not match the expected domain : %s", domain)
		}
	}

	return nil
}

func ValidateQuality(value string) error {
	if !isValidQuality(value) {
		return fmt.Errorf("quality text is wrong. must be one of : 'auto|1080p|720p|480p|480p|360p|240p|144p|SD|HD|FHD|4K'")
	}
	return nil
}

func ValidateImage(value string) error {
	if !isValidImage(value) {
		return fmt.Errorf("must be a valid image url")

	}
	return nil
}

func ValidateMAC(value string) error {
	if !isValidMAC(value) {
		return fmt.Errorf("must be a valid MAC Address")

	}

	_, err := net.ParseMAC(value)
	return err
}

func ValidateShow(value string) error {
	for i := range showTypes {
		if value == showTypes[i] {
			return nil
		}
	}

	return fmt.Errorf("must be a valid show Type : 'movie | tv | movie special | tv special | ova | ona | compilation'")
}
