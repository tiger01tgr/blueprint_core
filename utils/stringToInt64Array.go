package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ConvertStringToInt64Array(stringSlice string) []int64 {
	substrings := strings.Split(stringSlice, ",")

	// Initialize a slice to store the converted integers
	var int64Slice []int64

	for _, subStr := range substrings {
		intValue, err := strconv.ParseInt(subStr, 10, 64)
		if err == nil {
			int64Slice = append(int64Slice, intValue)
		} else {
			// Handle the error if the substring cannot be converted to an int64
			fmt.Printf("Error converting string to int64: %v\n", err)
		}
	}
	return int64Slice
}
