package whichin

import (
	"sort"
	"strings"
)

func InArray(array1 []string, array2 []string) []string {
	var r = []string{}
	var cur = make(map[string]int)
	for _, substr := range array1 {
		cur[substr]++
		if cur[substr] > 1 {
			continue
		}
		var wrdcount int
		for _, word := range array2 {
			if strings.Contains(word, substr) && wrdcount == 0 {
				r = append(r, substr)
				wrdcount++
			}
		}
	}
	sort.Strings(r)

	return r
}
