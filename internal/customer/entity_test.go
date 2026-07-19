package customer_test

import (
	"learn-trivia/internal/customer"
	"testing"
)

func TestNewCustomer_ShouldSuccess(t *testing.T) {
	nik, _ := customer.NewNik("3175983746738887")
	personName, _ := customer.NewPersonName("John", "Doe")
	cust, err := customer.NewCustomer(nik, personName)

	if err != nil {
		t.Fatalf("NewCustomer() error = %v", err)
	}
	if cust == nil {
		t.Fatalf("NewCustomer() customer = nil")
	}

	if cust.ID() == "" {
		t.Errorf("NewCustomer() id = %v, want not empty", cust.ID())
	}

	if cust.Nik() != nik {
		t.Errorf("NewCustomer() nik = %v, want %v", cust.Nik(), nik)
	}

	if cust.Name() != personName {
		t.Errorf("NewCustomer() firstName = %v, want %v", cust.Name(), personName)
	}
}
