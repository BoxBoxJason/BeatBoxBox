package format_utils

import (
	"reflect"
	"strconv"
	"strings"
)

// ConvertStringToIntArray converts a string to an array of integers
func ConvertStringToIntArray(raw_string string, separator string) ([]int, error) {
	raw_string_array := strings.Split(raw_string, separator)
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

// ConvertStringArrayToIntArray converts an array of strings to an array of integers
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

// ConvertRecordsToInterfaceArray converts a pointer to a slice of records to a slice of interfaces
func ConvertRecordsToInterfaceArray(records_ptr any) []interface{} {
	records := reflect.ValueOf(records_ptr).Elem()
	records_interface := make([]interface{}, records.Len())
	for i := 0; i < records.Len(); i++ {
		records_interface[i] = records.Index(i).Interface()
	}

	return records_interface
}

// CreateSliceOfAny creates a slice of any type
func CreateSliceOfAny(to_slice any) any {
	to_slice_type := reflect.TypeOf(to_slice).Elem()
	return reflect.New(reflect.SliceOf(to_slice_type)).Interface()
}
