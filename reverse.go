package string_util

import (
	"strings"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseWords(s string) string {
	if s == "" {
		return ""
	}
	substrings := strings.Split(s, " ")

	if len(substrings) == 1 {
		return s
	}

	reversed := make([]string,len(substrings), len(substrings))
	for index, sub := range substrings {
		reversed[len(reversed) - index - 1] = sub
	}
	return strings.Join(reversed, " ")
}



