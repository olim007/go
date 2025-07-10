package main

import "fmt"

func main() {
	var pointer *int64
	fmt.Println(pointer)
	fmt.Println(pointer == nil)

	amount := int64(60_000_00)
	pointer = &amount

	fmt.Println(pointer)
	fmt.Println(*pointer)
	fmt.Println(&pointer)
	fmt.Println()	
	first := &amount
	second := &amount
	*first = 30
	fmt.Println(*first)
	fmt.Println(amount)
	fmt.Println(*second)
}