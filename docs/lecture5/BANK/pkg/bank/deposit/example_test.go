package deposit

import "fmt"

func ExampleCalculate() {
	fmt.Println(Calculate(0, "TJS"))
	fmt.Println(Calculate(0, "USD"))
	fmt.Println(Calculate(500_000_00, "TJS"))
	fmt.Println(Calculate(500_000_00, "USD"))
	fmt.Println(Calculate(1_000_000, "TJS"))
	fmt.Println(Calculate(1_000_000, "USD"))

	// Output:
	// 0 0
	// 0 0
	// 2000000 3000000
	// 500000 1000000
	// 40000 60000
	// 10000 20000
}
