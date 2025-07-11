package payment

import (
	"fmt"
	"payment/pkg/types"
)

func ExamplePayment() {
	cards := []types.Card{
		{
		ID:       1111,
		Number:      "2323 xxxx xxxx 3333",
		Name:     "Olim",
		Balance:  22_000_00,
		Currency: "USD",
		Color:    "black",
		Active:   true,
		},
		{
			ID: 2222,
			Name: "Dfddf",
			Number: "2221 xxxx xxxx 1112",
			Currency: "TJS",
			Color: "yellow",
			Balance: 122343,
			Active: true,
		},
	}
	pS := Payment(cards)
	for _, ps := range pS {
		fmt.Println(ps.Number)
	}

	// Output:
	// 2323 xxxx xxxx 3333
	// 2221 xxxx xxxx 1112
}