package oddint

func FindOdd(seq []int) int {
	var newMap = make(map[int]int)
	for _, num := range seq {
		newMap[num]++
	}
	for key, val := range newMap {
		if val%2 != 0 {
			return key
		}
	}
	return 0
}

func FindOddXor(seq []int) int {
	res := 0
	for _, x := range seq {
		res ^= x
	}
	return res
}
