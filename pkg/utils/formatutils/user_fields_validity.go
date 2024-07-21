package format_utils

import (
	"regexp"
	"strings"
)

// CheckPseudoValidity checks if the pseudo is valid
func CheckPseudoValidity(pseudo string) bool {
	// Example criterion: pseudo must be between 3 and 32 characters
	return len(pseudo) >= 3 && len(pseudo) <= 32
}

// CheckEmailValidity checks if the email is valid using a regex pattern.
func CheckEmailValidity(email string) bool {
	regex := regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4}$`)
	return len(email) <= 256 && regex.MatchString(email)
}

// CheckRawPasswordValidity checks if the password contains at least one number, one special character, and alphanumeric characters.
func CheckRawPasswordValidity(rawPassword string) bool {
	return len(rawPassword) >= 6 && len(rawPassword) <= 64 && strings.ContainsAny(rawPassword, "0123456789") && strings.ContainsAny(rawPassword, "!@#$%^&*()-_+=[]{}|;:,.<>?") && strings.ContainsAny(rawPassword, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
}
