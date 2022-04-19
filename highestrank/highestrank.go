package highestrank

func HighestRank(nums []int) int {
	var newMap = make(map[int]int)
	for _, num := range nums {
		newMap[num]++
	}

	var count, most int
	for key, val := range newMap {
		if val > count || (val == count && key > most) {
			count, most = val, key
		}
	}

	return most
}

// Packing all inside one loop:
func HighestRankRef(nums []int) int {
	newMap, maxKey, maxVal := map[int]int{}, 0, 0
	for _, val := range nums {
		newMap[val]++
		if newMap[val] > maxVal || (newMap[val] == maxVal && val > maxKey) {
			maxKey, maxVal = val, newMap[val]
		}
	}
	return maxKey
}
