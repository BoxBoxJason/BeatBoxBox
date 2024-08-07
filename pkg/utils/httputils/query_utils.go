package httputils

import (
	custom_errors "BeatBoxBox/pkg/errors"
	format_utils "BeatBoxBox/pkg/utils/formatutils"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

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
				return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid query parameter: %s. Expected a single string, got %d", param, len(val)))
			}
		}
	}
	// Parse integer parameters and add them to the params map
	for _, param := range integers {
		if val, ok := query_params[param]; ok {
			if len(val) == 1 {
				integer_val, err := strconv.Atoi(strings.TrimSpace(val[0]))
				if err != nil {
					return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid query parameter: %s. Expected an integer, got %s", param, val[0]))
				} else if integer_val < 0 {
					return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid query parameter: %s. Expected a positive integer, got %d", param, integer_val))
				}
				params[param] = integer_val
			} else if len(val) > 1 {
				return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid query parameter: %s. Expected a single integer, got %d", param, len(val)))
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
					return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid query parameter: %s. Expected an integer, got %s", param, v))
				} else if integer_val < 0 {
					return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid query parameter: %s. Expected a positive integer, got %d", param, integer_val))
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
		return nil, custom_errors.NewBadRequestError("Error parsing multipart form: " + err.Error())
	}
	params := make(map[string]interface{})
	// Parse string parameters and add them to the params map
	for _, param := range strings_params {
		if val, ok := r.MultipartForm.Value[param]; ok {
			if len(val) == 1 {
				params[param] = strings.TrimSpace(val[0])
			} else if len(val) > 1 {
				return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected a single string, got %d", param, len(val)))
			}
		}
	}
	// Parse integer parameters and add them to the params map
	for _, param := range integers {
		if val, ok := r.MultipartForm.Value[param]; ok {
			if len(val) == 1 {
				integer_val, err := strconv.Atoi(strings.TrimSpace(val[0]))
				if err != nil {
					return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected an integer, got %s", param, val[0]))
				} else if integer_val < 0 {
					return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected a positive integer, got %d", param, integer_val))
				}
				params[param] = integer_val
			} else if len(val) > 1 {
				return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected a single integer, got %d", param, len(val)))
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
					return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected an integer, got %s", param, v))
				} else if integer_val < 0 {
					return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected a positive integer, got %d", param, integer_val))
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
				err = format_utils.ValidateRequestFile(file[0], format)
				if err != nil {
					return nil, err
				}
				params[param] = file[0]
			} else if len(file) > 1 {
				return nil, custom_errors.NewBadRequestError(fmt.Sprintf("Invalid multipart form parameter: %s. Expected a single file, got %d", param, len(file)))
			}
		}
	}
	return params, nil
}
