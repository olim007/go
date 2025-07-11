package main

import "fmt"

func sum(operations []int64) int64 {
	sum := int64(0)
	for _, operation := range operations {
		sum += operation
	} 
	return sum
}

func max(operations []int64) int64 {
	max := operations[0]
	for _, operoperation := range operations {
		if max <= operoperation {
			max = operoperation
		}
	}
	return max
}

func main() {

	var operations []int64

	operations = append(operations, 9)
	operations = append(operations, -3)
	operations = append(operations, 0)

	fmt.Println(operations[0])
	// operations := []int64{3, -3, 0}
	// fmt.Println(operations[0])

	// var operations [3]int64
	// operations[0] = 2
	// operations[1] = 1
	// operations[2] = 0
	// last := operations[2]
	// fmt.Println(last)
	// fmt.Println(len(operations))

	// arr := [3]int64{}
	// first := arr[0]
	// last := arr[len(arr) - 1]
	// fmt.Println(first, last)

	// for i := 0; i < len(operations); i++ {
	// 	fmt.Println(i)
	// }

	// index, value = range(operoperations)

	// for _, operation := range operations {
	// 	fmt.Println(operation)
	// }

	// var ssff [3]int64
	// fmt.Printf("%d\n", ssff[2])
	// fmt.Println()
	// gr := make([]int64, 3)
	// fmt.Println(gr[0])
}
