package createphone

import (
	"fmt"
	"strconv"
	"strings"
)

func CreatePhoneNumber(numbers [10]uint) string {
	var f, s, t string
	for i, num := range numbers {
		if i < 3 {
			f += strconv.FormatUint(uint64(num), 10)
		} else if i > 2 && i < 6 {
			s += strconv.FormatUint(uint64(num), 10)
		} else {
			t += strconv.FormatUint(uint64(num), 10)
		}
	}
	return fmt.Sprintf("(%s) %s-%s", f, s, t)
}

func CreatePhoneNumberRefactored(numbers [10]uint) string {
	str := strings.Trim(strings.Replace(fmt.Sprint(numbers), " ", "", -1), "[]")
	return fmt.Sprintf("(%s) %s-%s", str[0:3], str[3:6], str[6:10])
}
