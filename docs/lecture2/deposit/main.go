package main

func main() {
	amount := 999_999_00
	minPercent := 4
	maxPercent := 6
	minIncome := amount * minPercent / 100
	maxIncome := amount * maxPercent / 100
	println(minIncome)
	println(maxIncome)
}