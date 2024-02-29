package entity_test

import (
	"testing"

	currency "github.com/hashcode88/expense/internal/domain/entity/currency"
	organization "github.com/hashcode88/expense/internal/domain/entity/organization"
)

func TestCreateNewOrganization(t *testing.T) {
	organization := organization.NewOrganization("Hashcode88")
	if organization.Name == "" {
		t.Errorf("Expected value: %s but received empty value", "Hashcode88")
	}
}

func TestAddDefaultCurrency(t *testing.T) {
	organization := organization.NewOrganization("Hashcode88")
	currency := currency.CreateCurrency("6cadb2ba-0e9f-4ced-a2e6-af3fd472a976", "Real", 1.0)
	organization.AddDefaultCurrency(currency)
}
