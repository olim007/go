package payment

import "payment/pkg/types"

func Payment(cards []types.Card) []types.PaymentSource {
	pS := []types.PaymentSource{}
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			pS = append(pS, types.PaymentSource{Type: "card", Number: string(card.Number), Balance: card.Balance})
		}
	}
	return pS
}