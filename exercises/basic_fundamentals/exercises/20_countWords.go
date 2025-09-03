package exercises

import (
	"strings"
	"unicode"
)

func CountWords(text string) map[string]int {

	c := make(map[string]int)
	text = strings.ToUpper(text)

	// Remove punctuation
	text = strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, text)

	sliced_text := strings.Fields(text)

	for _, v := range sliced_text {
		c[v]++
	}

	return c
}
