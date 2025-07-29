package main

import (
	"fmt"
	"testing"
	"unsafe"
)

func Test(t *testing.T) {
	
}

func main() {
	categories := []string{"auto", "food", "beauty", "mobile", "fun"}
	top3 := categories[0:3]
	fmt.Println(categories)
	fmt.Println(len(categories))
	fmt.Println(cap(categories))
	fmt.Println(top3)
	fmt.Println(len(top3))
	fmt.Println(cap(top3))
	categories[1] = "it"
	fmt.Println(top3)
	fmt.Println(len(top3))
	fmt.Println(cap(top3))
	fmt.Println(categories)
	fmt.Println(len(categories))
	fmt.Println(cap(categories))
	categories = append(categories, "gadgets")
	categories[1] = "food"
	fmt.Println(categories)
	fmt.Println(cap(categories))
	fmt.Println(top3)
	fmt.Println(cap(top3))
}

type slice struct {
	array unsafe.Pointer
	len   int
	cap   int
}
