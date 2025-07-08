package bonus

func Calculate(amount int) int {
	minAmount  := 10
	if amount < minAmount {
		return 0
	}
	bonus := 0
	welcomeBonus := 5
	bonusRate := 5
	bonus = welcomeBonus + amount/bonusRate
	return bonus
}