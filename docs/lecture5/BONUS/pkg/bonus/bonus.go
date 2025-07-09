package bonus

func Calculate(amount int) (result int) {
	result = int(float64(amount) * 1.005)
	return result
}
