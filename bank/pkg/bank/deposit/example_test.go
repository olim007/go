package deposit

import "fmt"

func ExampleCalculate() {
	fmt.Println(Calculate(500_000, "TJS"))
	fmt.Println(Calculate(500_000, "USD"))

	// Output:
	// 50000 60000
	// 15000 20000
}