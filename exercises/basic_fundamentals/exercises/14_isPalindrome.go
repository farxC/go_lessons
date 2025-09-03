package exercises

import (
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "")
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes) == str
}
