package highestscoring

import (
	"fmt"
	"strings"
)

func High(s string) string {
	alph := "abcdefghijklmnopqrstuvwxyz"
	total := make(map[string]int)
	ns := strings.Fields(s)
	var hgst int

	for _, c := range ns {
		for _, lttr := range c {
			total[c] += strings.Index(alph, string(lttr)) + 1
		}
		if total[c] == hgst {
			total[c] = 0
		}
		if total[c] > hgst {
			hgst = total[c]
		}
	}

	fmt.Println("Total:", total, "Highest:", hgst)
	for key := range total {
		if total[key] == hgst {
			return key
		}
	}

	return ""
}
