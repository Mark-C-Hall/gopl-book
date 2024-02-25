package main

import (
	"strings"
)

func inefficient(words []string) string {
	var s, sep string
	for i := 1; i < len(words); i++ {
		s += sep + words[i]
		sep = " "
	}
	return s
}

func efficient(words []string) string {
	return strings.Join(words, " ")
}
