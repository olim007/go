package card

import (
	"bonus/pkg/bank/types"
	"fmt"
)

func ExampleDeposit() {
	card := types.Card{
		ID:       1111,
		PAN:      "2323 xxxx xxxx 3333",
		Name:     "Olim",
		Balance:  22_000_00,
		Currency: "USD",
		Color:    "black",
		Active:   true,
	}
	Deposit(&card, 1100)
	fmt.Println(card.Balance)

	// Output:
	// 2201100
}
