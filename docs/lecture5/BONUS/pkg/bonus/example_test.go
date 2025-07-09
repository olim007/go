package bonus

import "fmt"

func ExampleCalculate() {
	fmt.Println(Calculate(0))
	fmt.Println(Calculate(5000))
	fmt.Println(Calculate(10000))

	// Output:
	// 0
	// 5024
	// 10049
}
