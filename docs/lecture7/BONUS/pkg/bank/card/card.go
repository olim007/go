package card

import "bonus/pkg/bank/types"

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		ID:       1000,
		PAN:      "5058 xxxx xxxx 9999",
		Name:     name,
		Balance:  0,
		Currency: currency,
		Color:    color,
		Active:   true,
	}

	return card
}