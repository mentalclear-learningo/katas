package multsthreefive

// func Multiple3And5(number int) (sum int) {
// 	for i := 1; i < number; i++ {
// 		if i%3 == 0 {
// 			sum += i
// 		} else if i%5 == 0 {
// 			sum += i
// 		} else if i%3 == 0 && i%5 == 0 {
// 			continue
// 		}
// 	}

// 	return
// }

// After all the kata wasn't checking for negative numbers.

func Multiple3And5(number int) (sum int) {
	for i := 1; i < number; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return
}
