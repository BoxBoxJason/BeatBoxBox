package bool_utils

func CheckIntInArray(int_array []int, int_value int) bool {
	for _, value := range int_array {
		if value == int_value {
			return true
		}
	}
	return false
}

// CheckArrayMatch checks if two arrays are equal (contain the same values in a different order)
func CheckArrayMatch[T comparable](array1 []T, array2 []T) bool {
	if len(array1) != len(array2) {
		return false
	}

	counter1 := arrayValuesCounter(array1)
	for _, value := range array2 {
		remaining, ok := counter1[value]
		if !ok || remaining == 0 {
			return false
		}
		counter1[value]--
	}
	return true
}

// arrayValuesCounter counts the occurrences of each element in the slice.
func arrayValuesCounter[T comparable](array []T) map[T]int {
	counter := make(map[T]int)
	for _, value := range array {
		counter[value]++
	}
	return counter
}

func CheckStringInArray(string_array []string, string_value string) bool {
	for _, value := range string_array {
		if value == string_value {
			return true
		}
	}
	return false
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
