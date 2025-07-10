package types

type Money int64

type Currency string

const (
	USD Currency = "USD"
	TJS Currency = "TJS"
	RUB Currency = "RUB"
)

type Card struct {
	ID int
	PAN string
	Name string
	Balance Money
	Currency Currency
	Color string
	Active bool
}

type Payment struct {
	ID int
	Amount Money
}