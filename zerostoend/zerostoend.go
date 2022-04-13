package zerostoend

func MoveZeros(arr []int) []int {
	var nums = []int{}
	var zeros = []int{}

	for _, num := range arr {
		if num != 0 {
			nums = append(nums, num)
		} else {
			zeros = append(zeros, num)
		}
	}
	nums = append(nums, zeros...)

	return nums
}
