package card

import "card/pkg/bank/types"

func Deposit(card *types.Card, amount types.Money) {
	if card.Active && amount <= 50_000 {
		card.Balance += amount
	}
}