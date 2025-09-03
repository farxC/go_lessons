package exercises

import (
	"slices"
	"strings"
)

type ValuesCounted struct {
	consoants int
	vowels    int
}

func CountConsonantAndVowels(text string) ValuesCounted {
	c, v := 0, 0

	vowels := []string{"a", "e", "i", "o", "u"}
	consonants := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "w", "x", "y", "z"}
	text = strings.ToLower(text)
	text_sliced := strings.Split(text, "")

	for _, value := range text_sliced {
		if slices.Contains(consonants, value) {
			c += 1
		}
		if slices.Contains(vowels, value) {
			v += 1
		}
	}

	return ValuesCounted{
		consoants: c,
		vowels:    v,
	}
}
