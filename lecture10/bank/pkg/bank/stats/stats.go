package stats

import "github.com/olim007/go/pkg/bank/types"

func Avg(payments []types.Payment) types.Money {
	sum := 0
	for _, c := range payments {
		sum += int(c.Amount)
	}
	avg := sum / len(payments)
	return types.Money(avg)
}
