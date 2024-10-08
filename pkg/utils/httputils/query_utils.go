package httputils

import (
	"BeatBoxBox/pkg/logger"
	"fmt"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
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

// ParseQueryParams parses the query parameters of an HTTP request and returns a map of the parameters.
// The function takes in four lists of strings: strings_params, integers, strings_lists, and integers_lists.
// The strings_params list contains the names of the query parameters that are expected to be strings.
// The integers list contains the names of the query parameters that are expected to be integers.
// The strings_lists list contains the names of the query parameters that are expected to be lists of strings.
// The integers_lists list contains the names of the query parameters that are expected to be lists of integers.
// The function returns a map of the query parameters.
func ParseQueryParams(r *http.Request, strings_params []string, integers []string, strings_lists []string, integers_lists []string) (map[string]interface{}, error) {
	query_params := r.URL.Query()
	params := make(map[string]interface{})
	// Parse string parameters and add them to the params map
	for _, param := range strings_params {
		if val, ok := query_params[param]; ok {
			if len(val) == 1 {
				params[param] = strings.TrimSpace(val[0])
			} else if len(val) > 1 {
				return nil, NewBadRequestError(fmt.Sprintf("Invalid query parameter: %s. Expected a single string, got %d", param, len(val)))
			}
		}
	}
	// Parse integer parameters and add them to the params map
	for _, param := range integers {
		if val, ok := query_params[param]; ok {
			if len(val) == 1 {
				integer_val, err := strconv.Atoi(strings.TrimSpace(val[0]))
				if err != nil {
					return nil, NewBadRequestError(fmt.Sprintf("Invalid query parameter: %s. Expected an integer, got %s", param, val[0]))
				} else if integer_val < 0 {
					return nil, NewBadRequestError(fmt.Sprintf("Invalid query parameter: %s. Expected a positive integer, got %d", param, integer_val))
				}
				params[param] = integer_val
			} else if len(val) > 1 {
				return nil, NewBadRequestError(fmt.Sprintf("Invalid query parameter: %s. Expected a single integer, got %d", param, len(val)))
			}
		}
	}
	// Parse string list parameters and add them to the params map
	for _, param := range strings_lists {
		if val, ok := query_params[param]; ok {
			for i, v := range val {
				val[i] = strings.TrimSpace(v)
			}
			params[param] = val
		}
	}
	// Parse integer list parameters and add them to the params map
	for _, param := range integers_lists {
		if val, ok := query_params[param]; ok {
			int_list := make([]int, len(val))
			for i, v := range val {
				integer_val, err := strconv.Atoi(strings.TrimSpace(v))
				if err != nil {
					return nil, NewBadRequestError(fmt.Sprintf("Invalid query parameter: %s. Expected an integer, got %s", param, v))
				} else if integer_val < 0 {
					return nil, NewBadRequestError(fmt.Sprintf("Invalid query parameter: %s. Expected a positive integer, got %d", param, integer_val))
				}
				int_list[i] = integer_val
			}
			params[param] = int_list
		}
	}
	return params, nil
}

// ParseMultiPartFormParams parses the multipart form parameters of an HTTP request and returns a map of the parameters.
// The function takes in five lists of strings: strings_params, integers, strings_lists, integers_lists, and files.
// The strings_params list contains the names of the multipart form parameters that are expected to be strings.
// The integers list contains the names of the multipart form parameters that are expected to be integers.
// The strings_lists list contains the names of the multipart form parameters that are expected to be lists of strings.
// The integers_lists list contains the names of the multipart form parameters that are expected to be lists of integers.
// The files map contains the names of the multipart form parameters that are expected to be files and their corresponding formats.
func ParseMultiPartFormParams(r *http.Request, strings_params []string, integers []string, strings_lists []string, integers_lists []string, files map[string]string) (map[string]interface{}, error) {
	err := r.ParseMultipartForm(10 << 35)
	if err != nil {
		return nil, NewBadRequestError("Error parsing multipart form: " + err.Error())
	}
	params := make(map[string]interface{})
	// Parse string parameters and add them to the params map
	for _, param := range strings_params {
		if val, ok := r.MultipartForm.Value[param]; ok {
			if len(val) == 1 {
				params[param] = strings.TrimSpace(val[0])
			} else if len(val) > 1 {
				return nil, NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected a single string, got %d", param, len(val)))
			}
		}
	}
	// Parse integer parameters and add them to the params map
	for _, param := range integers {
		if val, ok := r.MultipartForm.Value[param]; ok {
			if len(val) == 1 {
				integer_val, err := strconv.Atoi(strings.TrimSpace(val[0]))
				if err != nil {
					return nil, NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected an integer, got %s", param, val[0]))
				} else if integer_val < 0 {
					return nil, NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected a positive integer, got %d", param, integer_val))
				}
				params[param] = integer_val
			} else if len(val) > 1 {
				return nil, NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected a single integer, got %d", param, len(val)))
			}
		}
	}
	// Parse string list parameters and add them to the params map
	for _, param := range strings_lists {
		if val, ok := r.MultipartForm.Value[param]; ok {
			for i, v := range val {
				val[i] = strings.TrimSpace(v)
			}
			params[param] = val
		}
	}
	// Parse integer list parameters and add them to the params map
	for _, param := range integers_lists {
		if val, ok := r.MultipartForm.Value[param]; ok {
			int_list := make([]int, len(val))
			for i, v := range val {
				integer_val, err := strconv.Atoi(strings.TrimSpace(v))
				if err != nil {
					return nil, NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected an integer, got %s", param, v))
				} else if integer_val < 0 {
					return nil, NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected a positive integer, got %d", param, integer_val))
				}
				int_list[i] = integer_val
			}
			params[param] = int_list
		}
	}
	// Parse file parameters and add them to the params map
	for param, format := range files {
		if file, ok := r.MultipartForm.File[param]; ok {
			if len(file) == 1 {
				err = ValidateRequestFile(file[0], format)
				if err != nil {
					return nil, err
				}
				params[param] = file[0]
			} else if len(file) > 1 {
				return nil, NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected a single file, got %d", param, len(file)))
			}
		}
	}
	return params, nil
}

func ParseFormBodyParams(r *http.Request, strings_params []string, integers []string, strings_lists []string, integers_lists []string) (map[string]interface{}, error) {
	if r.ParseForm() != nil {
		return nil, NewBadRequestError("Error parsing form")
	}
	params := make(map[string]interface{})
	// Parse string parameters and add them to the params map
	for _, param := range strings_params {
		if val, ok := r.PostForm[param]; ok {
			if len(val) == 1 {
				params[param] = strings.TrimSpace(val[0])
			} else if len(val) > 1 {
				return nil, NewBadRequestError(fmt.Sprintf("Invalid form parameter: %s. Expected a single string, got %d", param, len(val)))
			}
		}
	}
	// Parse integer parameters and add them to the params map
	for _, param := range integers {
		if val, ok := r.PostForm[param]; ok {
			if len(val) == 1 {
				integer_val, err := strconv.Atoi(strings.TrimSpace(val[0]))
				if err != nil {
					return nil, NewBadRequestError(fmt.Sprintf("Invalid form parameter: %s. Expected an integer, got %s", param, val[0]))
				} else if integer_val < 0 {
					return nil, NewBadRequestError(fmt.Sprintf("Invalid form parameter: %s. Expected a positive integer, got %d", param, integer_val))
				}
				params[param] = integer_val
			} else if len(val) > 1 {
				return nil, NewBadRequestError(fmt.Sprintf("Invalid form parameter: %s. Expected a single integer, got %d", param, len(val)))
			}
		}
	}
	// Parse string list parameters and add them to the params map
	for _, param := range strings_lists {
		if val, ok := r.PostForm[param]; ok {
			for i, v := range val {
				val[i] = strings.TrimSpace(v)
			}
			params[param] = val
		}
	}
	// Parse integer list parameters and add them to the params map
	for _, param := range integers_lists {
		if val, ok := r.PostForm[param]; ok {
			int_list := make([]int, len(val))
			for i, v := range val {
				integer_val, err := strconv.Atoi(strings.TrimSpace(v))
				if err != nil {
					return nil, NewBadRequestError(fmt.Sprintf("Invalid form parameter: %s. Expected an integer, got %s", param, v))
				} else if integer_val < 0 {
					return nil, NewBadRequestError(fmt.Sprintf("Invalid form parameter: %s. Expected a positive integer, got %d", param, integer_val))
				}
				int_list[i] = integer_val
			}
			params[param] = int_list
		}
	}
	return params, nil
}

// ValidateRequestFile validates the file by checking the file size, file name, and file format.
func ValidateRequestFile(file_header *multipart.FileHeader, format string) error {
	if file_header.Size > FILE_SIZE_BY_TYPE[format] {
		return NewBadRequestError(fmt.Sprintf("File size too large, max size is %d bytes, got %d", FILE_SIZE_BY_TYPE[format], file_header.Size))
	} else if file_header.Filename != filepath.Base(file_header.Filename) {
		return NewBadRequestError("Invalid file name, detected path traversal")
	} else if err := CheckFileFormat(file_header, format); err != nil {
		return err
	}
	return nil
}

// CheckFileFormat checks if the file format is valid by comparing the file header and content.
func CheckFileFormat(file_header *multipart.FileHeader, format string) error {

	if FILE_FORMATS[format] == nil {
		return NewBadRequestError(fmt.Sprintf("Invalid file format: %s", format))
	} else if !FILE_FORMATS[format][file_header.Header.Get("Content-Type")] {
		return NewBadRequestError(fmt.Sprintf("Invalid file format detected from header, expected %s, got %s", format, file_header.Header.Get("Content-Type")))
	}
	src, err := file_header.Open()
	if err != nil {
		return NewBadRequestError("Error opening file: " + err.Error())
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
		return NewBadRequestError("Error reading file: " + err.Error())
	}
	mime_type := http.DetectContentType(buffer[:n])
	if !FILE_FORMATS[format][mime_type] {
		return NewBadRequestError(fmt.Sprintf("Invalid file format detected from content, expected %s, got %s", format, mime_type))
	}
	return nil
}
