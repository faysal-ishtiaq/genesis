package utils

import (
	"fmt"
	"unicode"
)

// FirstLower converts the first letter of a string to lowercase and returns the lowercase string
func FirstLower(str string) string {
	for _, c := range str {
		return fmt.Sprintf("%c%s", unicode.ToLower(c), str[1:])
	}
	return str
}
