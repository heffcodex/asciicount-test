package helper

func CountDigits(i int) (count int) {
	for i != 0 {
		i /= 10
		count++
	}

	return count
}
