package utils

import (
	"fmt"
	"os"
	"reflect"
)

// Contains returns true if a string is in a slice, otherwise false
func Contains(haystack interface{}, needle interface{}) bool {

	haystackValue := reflect.ValueOf(haystack)
	if haystackValue.Kind() != reflect.Slice {
		fmt.Println("Error: expected type slice, got", haystackValue.Type())
		os.Exit(1)
	}

	for idx := 0; idx < haystackValue.Len(); idx++ {
		if haystackValue.Index(idx).Interface() == needle {
			return true
		}
	}

	return false
}
