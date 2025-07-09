package deposit

// Calculate calculates the deposit params
func Calculate(amount int, currency string) (min int, max int) {
	minPercent, maxPercent := PersonFor(currency)

	min = amount * minPercent / 100
	max = amount * maxPercent / 100

	return min, max
}

func PersonFor(currency string) (min int, max int) {
	switch currency {
	case "TJS":
		return 10, 12
	case "USD":
		return 3, 4
	default:
		return 0, 0	
	}
}