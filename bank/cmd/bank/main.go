package main

import (
	"bank/pkg/bank/credit"
	"bank/pkg/bank/deposit"
	"fmt"
)

func main() {
	credit.Calculate(1,1,1,"")
	min, max := deposit.Calculate(500_000_00, "TJS")
	fmt.Println(min, max)
}