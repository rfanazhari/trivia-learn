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
		t.Errorf("NewCustomer() error = %v", err)
	}
	if cust == nil {
		t.Errorf("NewCustomer() customer = nil")
	}

	if cust.GetID() == "" {
		t.Errorf("NewCustomer() id = %v, want not empty", cust.GetID())
	}

	if cust.GetNik() != nik.String() {
		t.Errorf("NewCustomer() nik = %v, want %v", cust.GetNik(), nik)
	}

	if cust.GetFirstName() != personName.GetFirstName() {
		t.Errorf("NewCustomer() firstName = %v, want %v", cust.GetFirstName(), personName.GetFirstName())
	}

	if cust.GetLastName() != personName.GetLastName() {
		t.Errorf("NewCustomer() lastName = %v, want %v", cust.GetLastName(), personName.GetLastName())
	}
}
