package mexwave

import (
	"unicode"
)

func Wave(words string) []string {
	res := []string{}

	i := 0
	for i < len(words) {
		word := []rune(words)
		word[i] = unicode.ToUpper(rune(word[i]))
		if word[i] != 32 {
			res = append(res, string(word))
		}
		i++
	}

	return res
}
