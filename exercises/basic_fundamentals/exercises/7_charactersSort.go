package exercises

import (
	"slices"
	"strings"
)

func CharactersSort(str string) []string {
	str = strings.ToLower(str)
	sliced_str := strings.Split(str, "")
	slices.Sort(sliced_str)
	return sliced_str
}
