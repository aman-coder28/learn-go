package main

import (
	"strings"
	"unicode"
)

func capitalize(word string) string {
	if word == "" {
		return word
	}
	
	first := unicode.ToUpper(rune(word[0]))
    others := word[1:]
	
	return string(first) + others
}

func capper(word string) string {
	words := strings.Fields(word)
	
	for i, w := range words {
		words[i] = capitalize(w)
	}
	
	return strings.Join(words, " ")
}
