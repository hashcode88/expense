package entity_test

import (
	"testing"

	entity "github.com/hashcode88/expense/internal/domain/entity/currency"
)

func TestNewCurrency(t *testing.T) {
	currency := entity.NewCurrency("Dolar", 4.90)
	if currency == nil {
		t.Errorf("expected new currency but reciverd nil value")
	}
}

func TestCreateCurrency(t *testing.T) {
	currency := entity.CreateCurrency("850dd894-7587-466c-85b6-61c54e822bdf", "Dolar", 4.90)
	if currency.CurrencyId == "" {
		t.Errorf("curreny id cannot be empty value")
	}
}
