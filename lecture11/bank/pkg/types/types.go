package types

// Money представляет собой денежную сумму в минимальных удиницах (дирамы, копейки, центы и т.д.)
type Money int64

// Currency представляет код валюты
type Currency string

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PAN string

type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}
