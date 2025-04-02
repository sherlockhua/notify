package entity

import "fmt"

type Money struct {
	Amount   int64
	Currency string
}

func (m *Money) Add(amount int64, currency string) {
	m.Amount += amount
	m.Currency = currency
}

func (m *Money) Subtract(amount int64, currency string) bool {
	if m.Amount < amount {
		return false
	}
	m.Amount -= amount
	m.Currency = currency
	return true
}

func (m *Money) Equal(amount int64, currency string) bool {
	return m.Amount == amount && m.Currency == currency
}

func (m *Money) LessThan(amount int64, currency string) bool {
	return m.Amount < amount && m.Currency == currency
}

func (m *Money) GreaterThan(amount int64, currency string) bool {
	return m.Amount > amount && m.Currency == currency
}

func (m *Money) String() string {
	return fmt.Sprintf("%d %s", m.Amount, m.Currency)
}

func NewMoney(amount int64, currency string) *Money {
	return &Money{Amount: amount, Currency: currency}
}
