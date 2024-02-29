package entity

import "github.com/google/uuid"

type Currency struct {
	CurrencyId string
	Name       string
	Value      float64
}

func CreateCurrency(currencyId string, name string, value float64) *Currency {
	return &Currency{
		CurrencyId: currencyId,
		Name:       name,
		Value:      value,
	}
}

func NewCurrency(name string, value float64) *Currency {
	return &Currency{
		CurrencyId: uuid.New().String(),
		Name:       name,
		Value:      value,
	}
}
