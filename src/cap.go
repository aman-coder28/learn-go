package main

import (
	"strings"
	"unicode"
)

func capper(word string) string {
	if len(word) < 1 {
		return word
	}

	first := unicode.ToUpper(rune(word[0]))

	return string(first) + word[1:]
}

func Capitalize(word string) string {
	words := strings.Fields(word)

	for i, w := range words {
		words[i] = capper(w)
	}

	return strings.Join(words, " ")
}
