package card

import (
	"fmt"
	"total/pkg/types"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID: 1,
			Amount: 1000,
		},
		{
			ID: 2,
			Amount: 3000,
		},
		{
			ID: 3,
			Amount: 2230,
		},
	}
	
	max := Max(payments)
	fmt.Println(max)
}

func Max(payments []types.Payment) types.Payment {
	maxPayment := payments[0]
	for _, papayment := range payments {
		if maxPayment.Amount < papayment.Amount {
			maxPayment = papayment
		}
	}
	return maxPayment
}