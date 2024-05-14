package format_utils

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// CheckPseudoValidity checks if the pseudo is valid
func CheckPseudoValidity(pseudo string) bool {
	// Example criterion: pseudo must be between 3 and 32 characters
	return len(pseudo) >= 3 && len(pseudo) <= 32
}

// CheckEmailValidity checks if the email is valid using a regex pattern.
func CheckEmailValidity(email string) bool {
	regex := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`)
	return len(email) <= 256 && regex.MatchString(email)
}

// CheckRawPasswordValidity checks if the password contains at least one number, one special character, and alphanumeric characters.
func CheckRawPasswordValidity(rawPassword string) bool {
	var lengthValid, hasUpper, hasLower, hasNumber, hasSpecial bool
	lengthValid = len(rawPassword) >= 6 && len(rawPassword) <= 32

	for _, ch := range rawPassword {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasNumber = true
		case strings.ContainsRune("!@#$%^&*(),.?\":{}|<>", ch):
			hasSpecial = true
		}
	}

	return lengthValid && hasUpper && hasLower && hasNumber && hasSpecial
}

func ConvertStringToIntArray(raw_string string, separator string) ([]int, error) {
	raw_string_array := strings.Split(raw_string, ",")
	int_array := make([]int, len(raw_string_array))
	for i, raw_string := range raw_string_array {
		int_value, err := strconv.Atoi(raw_string)
		if err != nil {
			return nil, err
		}
		int_array[i] = int_value
	}
	return int_array, nil
}

func ConvertStringArrayToIntArray(raw_string_array []string) ([]int, error) {
	int_array := make([]int, len(raw_string_array))
	for i, raw_string := range raw_string_array {
		int_value, err := strconv.Atoi(raw_string)
		if err != nil {
			return nil, err
		}
		int_array[i] = int_value
	}
	return int_array, nil
}

func CheckIntInArray(int_array []int, int_value int) bool {
	for _, value := range int_array {
		if value == int_value {
			return true
		}
	}
	return false
}

func CheckStringInArray(string_array []string, string_value string) bool {
	for _, value := range string_array {
		if value == string_value {
			return true
		}
	}
	return false
}
