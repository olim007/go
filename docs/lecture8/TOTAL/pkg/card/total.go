package card

import "total/pkg/types"

func Total(cards []types.Card) types.Money {
	sum := types.Money(0)
	for _, card := range cards {
		sum += card.Balance
	}
	return sum
}