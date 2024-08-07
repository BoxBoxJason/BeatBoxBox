package format_utils

import (
	custom_errors "BeatBoxBox/pkg/errors"
	"BeatBoxBox/pkg/logger"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"
)

const MAX_MUSIC_FILE_SIZE = 25 * 1024 * 1024
const MAX_IMAGE_FILE_SIZE = 5 * 1024 * 1024
const MAX_REQUEST_SIZE = MAX_IMAGE_FILE_SIZE + MAX_MUSIC_FILE_SIZE + 2048

var FILE_SIZE_BY_TYPE = map[string]int64{
	"image": MAX_IMAGE_FILE_SIZE,
	"audio": MAX_MUSIC_FILE_SIZE,
}
var FILE_FORMATS = map[string]map[string]bool{
	"image": {
		"image/jpeg":    true,
		"image/png":     true,
		"image/gif":     true,
		"image/webp":    true,
		"image/svg+xml": true,
		"image/x-icon":  true,
	},
	"audio": {
		"audio/mpeg":  true,
		"audio/ogg":   true,
		"audio/wav":   true,
		"audio/flac":  true,
		"audio/x-m4a": true,
		"audio/x-wav": true,
		"audio/aiff":  true,
	},
}

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

// ValidateRequestFile validates the file by checking the file size, file name, and file format.
func ValidateRequestFile(file_header *multipart.FileHeader, format string) error {
	if file_header.Size > FILE_SIZE_BY_TYPE[format] {
		return custom_errors.NewBadRequestError(fmt.Sprintf("File size too large, max size is %d bytes, got %d", FILE_SIZE_BY_TYPE[format], file_header.Size))
	} else if file_header.Filename != filepath.Base(file_header.Filename) {
		return custom_errors.NewBadRequestError("Invalid file name, detected path traversal")
	} else if err := CheckFileFormat(file_header, format); err != nil {
		return err
	}
	return nil
}

// CheckFileFormat checks if the file format is valid by comparing the file header and content.
func CheckFileFormat(file_header *multipart.FileHeader, format string) error {

	if FILE_FORMATS[format] == nil {
		return custom_errors.NewBadRequestError(fmt.Sprintf("Invalid file format: %s", format))
	} else if !FILE_FORMATS[format][file_header.Header.Get("Content-Type")] {
		return custom_errors.NewBadRequestError(fmt.Sprintf("Invalid file format detected from header, expected %s, got %s", format, file_header.Header.Get("Content-Type")))
	}
	src, err := file_header.Open()
	if err != nil {
		return custom_errors.NewBadRequestError("Error opening file: " + err.Error())
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			logger.Error("Error closing file: " + err.Error())
		}
	}(src)
	buffer := make([]byte, 512)
	n, err := src.Read(buffer)
	if err != nil {
		return custom_errors.NewBadRequestError("Error reading file: " + err.Error())
	}
	mime_type := http.DetectContentType(buffer[:n])
	if !FILE_FORMATS[format][mime_type] {
		return custom_errors.NewBadRequestError(fmt.Sprintf("Invalid file format detected from content, expected %s, got %s", format, mime_type))
	}
	return nil
}

func IsValidDate(date string) bool {
	return regexp.MustCompile(`\d{4}-\d{2}-\d{2}`).MatchString(date)
}
