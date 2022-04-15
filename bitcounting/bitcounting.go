package bitcounting

import (
	"math/bits"
	"strconv"
	"strings"
)

func CountBitsSimplified(n uint) int {
	return bits.OnesCount(n)
}

func CountBits(n uint) (count int) {
	str := strconv.FormatInt(int64(n), 2)
	nums := strings.Split(str, "")
	for _, num := range nums {
		cnum, err := strconv.Atoi(num)
		if err == nil {
			if cnum == 1 {
				count++
			}
		}
	}
	return
}
