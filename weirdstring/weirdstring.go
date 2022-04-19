package weirdstring

import (
	"strings"
)

func ToWeirdCase(str string) string {
	var result string
	for _, word := range strings.Fields(str) {
		var neword string
		for i, lttr := range word {
			if i == 0 || i%2 == 0 {
				neword += strings.ToUpper(string(lttr))
			} else {
				neword += strings.ToLower(string(lttr))
			}
		}
		result += neword + " "
	}

	return strings.TrimSpace(result)
}
