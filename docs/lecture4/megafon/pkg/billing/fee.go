package billing

func Calculate1000(meg int) float64 {
	if meg < 2000 {
		return float64(meg/1000)*35.00 + float64(meg%1000)*0.06
	} else {
		return Calculate2000(meg)
	}
}
func Calculate2000(meg int) float64 {
	if meg < 3000 {
		return float64(meg/2000)*55.00 + float64(meg%2000)*0.06
	} else {
		return Calculate3000(meg)
	}
}
func Calculate3000(meg int) float64 {
	if meg < 5000 {
		return float64(meg/3000)*70.00 + float64(meg%3000)*0.06
	} else {
		return Calculate5000(meg)
	}
}
func Calculate5000(meg int) float64 {
	if meg < 10000 {
		return float64(meg/5000)*95.00 + float64(meg%5000)*0.06
	} else {
		return Calculate10000(meg)
	}
}
func Calculate10000(meg int) float64 {
	return float64(meg/10000)*170.00 + float64(meg%10000)*0.06
}
