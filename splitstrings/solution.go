package splitstings

func Solution(str string) (result []string) {
	if len(str)%2 != 0 {
		str += "_"
	}
	for i := 0; i < len(str); i += 2 {
		result = append(result, str[i:i+2])
	}
	return
}
