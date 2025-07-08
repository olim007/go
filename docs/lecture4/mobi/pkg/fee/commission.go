package fee

func Calculate(amount float64) float64 {
	if amount < 5000 {
		return 0;
	}
	return amount * 0.005
}