package card

import (
	"bonus/pkg/bank/types"
)

func addBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {
	if (card.Active && card.Balance > 0) {
		card.Balance += types.Money(int64(float64(card.Balance) * 0.03) * int64(daysInMonth / daysInYear))
	}
}