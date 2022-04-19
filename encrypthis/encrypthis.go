package encrypthis

import (
	"strconv"
	"strings"
)

func EncryptThis(text string) string {
	var cypher string
	for _, word := range strings.Fields(text) {
		var neword string
		var last rune
		for i, lttr := range word {
			if i == 0 {
				neword += strconv.Itoa(int(lttr))
			} else if i == 1 {
				last = lttr
				neword += string(word[len(word)-1])
			} else if i == len(word)-1 {
				neword += string(last)
			} else {
				neword += string(lttr)
			}
		}
		cypher += neword + " "
	}

	return strings.TrimSpace(cypher)
}
