package utils

import (
	"strconv"
	"strings"
)

func StrSplit(menus string) []int {
	// Split the string by commas
	parts := strings.Split(menus, ",")
	// Create a slice to hold the integers
	var result []int
	// Convert each part to an integer and append to the result slice
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}
