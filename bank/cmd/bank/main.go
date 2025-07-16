package main

import (
	// "bank/pkg/bank/credit"
	// "bank/pkg/bank/deposit"
	"github.com/olim007/go/bank/pkg/bank/types"
	"fmt"
)


func main() {
	// credit.Calculate(1,1,1,"")
	// min, max := deposit.Calculate(500_000_00, "TJS")
	// fmt.Println(min, max)

	card := types.Card {
		ID: 123,
		PAN: "5833 xxxx xxxx 9999",
		Balance: 999_99,
		Currency: "TJS",
		Color: "white",
		Name: "infinity",
		Active: true,
	}
	handle(card)
	fmt.Printf("%+v", card)
}

func handle(card types.Card) {
	// to do
}