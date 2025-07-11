package main

import "bonus/pkg/bank/types"

func main() {
	card := types.Card{
		ID: 121,
		PAN: "2222 xxxx xxxx 1111",
		Name: "ED",
		Balance: 121212,
		Currency: "TJS",
		Color: "white",
		Active: true,
		MinBalance: 1111,
	}

	card.Active = false
}