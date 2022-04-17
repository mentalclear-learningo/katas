package finduniq

func FindUniq(arr []float32) float32 {
	var newMap = make(map[float32]int)
	for _, num := range arr {
		newMap[num]++
	}
	for key, val := range newMap {
		if val == 1 {
			return key
		}
	}

	return 0
}
