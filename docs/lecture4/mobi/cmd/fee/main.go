package main

import (
	"fmt"
	"mobi/pkg/fee"
)
func main() {
	fmt.Println(fee.Calculate(9_999_99))
}