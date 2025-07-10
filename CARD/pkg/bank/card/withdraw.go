package card

import "card/pkg/bank/types"

const WithdrawLimit = 20_000_00

func Withdraw(card *types.Card, amount types.Money) {
	if !card.Active {
		return
	}
	if amount > card.Balance {
		return
	}
	if amount < 0 {
		return
	}
	if amount >= 20_000 {
		return
	}
	if amount > WithdrawLimit {
		return
	}	
	card.Balance = card.Balance - amount
} 