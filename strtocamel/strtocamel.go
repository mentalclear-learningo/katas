package strtocamel

import (
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	// Create slice based on the regex, can be done in one line
	words := regexp.MustCompile(`-|_`).Split(s, -1)

	// slicing words[1:] to use range starting from the 2nd word
	for i, w := range words[1:] {

		// Changing value in the existing slice to start with capitla letter,
		// +1 needed because index is now moved now due to words[1:]
		words[i+1] = strings.Title(w)
	}

	return strings.Join(words, "")
}
