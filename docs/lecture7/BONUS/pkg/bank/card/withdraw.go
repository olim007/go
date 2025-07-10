package card

import "bonus/pkg/bank/types"

func Withdraw(card types.Card, amount types.Money) types.Card {
	if !card.Active {
		return card
	} else if amount > card.Balance {
		return card
	} else if amount < 0 {
		return card
	} else if amount >= 20_000 {
		return card
	} else {
		card.Balance = card.Balance - amount
		return card
	}
} 