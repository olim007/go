package main

import (
	"fmt"
	"megafon/pkg/billing"
)

func main() {
	result := billing.Calculate1000(9999)
	fmt.Println(result)
}