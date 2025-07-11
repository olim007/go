package types

type Money int64

type PaymentSource struct {
	Type string
	Number string
	Balance Money
}

type PAN string

type Card struct {
	ID int
	Name string
	Number PAN
	Currency string
	Color string
	Balance Money
	Active bool
}
