package card

import "bonus/pkg/bank/types"

func Deposit(card *types.Card, amount types.Money) {
	if card.Active && amount <= 50_000 {
		card.Balance += amount
	}
}