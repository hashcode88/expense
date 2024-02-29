package entity

import (
	"github.com/google/uuid"
	currency "github.com/hashcode88/expense/internal/domain/entity/currency"
)

type Organization struct {
	OrganizationId string
	Name           string
	Currency       *currency.Currency
}

func NewOrganization(name string) *Organization {
	return &Organization{
		OrganizationId: uuid.New().String(),
		Name:           name,
	}
}

func (o *Organization) AddDefaultCurrency(currency *currency.Currency) {
	o.Currency = currency
}
