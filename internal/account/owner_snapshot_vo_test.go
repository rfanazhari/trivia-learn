package account_test

import (
	"learn-trivia/internal/account"
	"learn-trivia/internal/customer"
	"testing"
)

func TestNewOwnerSnapshot_ShouldMapFieldsCorrectly(t *testing.T) {
	personName, _ := customer.NewPersonName("John", "Doe")

	snapshot := account.NewOwnerSnapshot(personName)

	if snapshot.FirstName() != personName.FirstName() {
		t.Fatalf("Expected first name to be %s, but got %s", personName.FirstName(), snapshot.FirstName())
	}

	if snapshot.LastName() != personName.LastName() {
		t.Fatalf("Expected last name to be %s, but got %s", personName.LastName(), snapshot.LastName())
	}
}
