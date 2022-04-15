package spinwords

import (
	"strings"
)

func SpinWords(str string) string {
	var result string
	words := strings.Fields(str)
	for _, word := range words {
		if len(word) > 4 {
			result += rotateWord(word) + " "
		} else {
			result += word + " "
		}
	}
	return strings.TrimSpace(result)
}

func rotateWord(word string) string {
	rune_arr := []rune(word)
	var rev []rune
	for i := len(rune_arr) - 1; i >= 0; i-- {
		rev = append(rev, rune_arr[i])
	}
	return string(rev)
}
