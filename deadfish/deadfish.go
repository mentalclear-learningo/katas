package deadfish

func Parse(data string) []int {
	result := []int{}
	var value int

	for _, lttr := range data {
		switch string(lttr) {
		case "i":
			value++
		case "d":
			value--
		case "s":
			value *= value
		case "o":
			result = append(result, value)
		}

	}

	return result
}
