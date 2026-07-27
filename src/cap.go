package main

import (
	"strings"
	"unicode"
)

func cap(word string) string {
	first := unicode.ToUpper(rune(word[0]))
	others := word[1:]

	return string(first) + others
}

func capper(word string) string {
	splitted := strings.SplitSeq(word, " ")
	w := ""

	for v := range splitted {
		w += " " + cap(v)
	}

	return w[1:]
}
