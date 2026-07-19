package account_test

import (
	"errors"
	"learn-trivia/internal/account"
	"learn-trivia/internal/customer"
	"testing"

	"github.com/google/uuid"
)

func TestNewAccount_ShouldSuccess(t *testing.T) {
	personName, _ := customer.NewPersonName("John", "Doe")
	owner := account.NewOwnerSnapshot(personName)
	balance, _ := account.NewMoney(0)

	acc, err := account.NewAccount("580fbc95-5b2b-4b8d-a685-6c7d788e4b68", owner, balance)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if acc.ID() == "" {
		t.Errorf("expected ID to be non-empty, got %s", acc.ID())
	}

	if acc.CustomerID() == "" {
		t.Error("expected CustomerID to be non-empty")
	}

	if acc.Owner() != owner {
		t.Errorf("expected owner %s , got %s", owner, acc.Owner())
	}

	if acc.Balance() != balance {
		t.Errorf("expected balance %v , got %v", balance, acc.Balance())
	}

	if acc.Status() != account.StatusPending {
		t.Errorf("expected status Pending, got: %v", acc.Status())
	}
}

func TestNewAccount_EmptyCustomerID_ShouldFail(t *testing.T) {
	personName, _ := customer.NewPersonName("John", "Doe")
	owner := account.NewOwnerSnapshot(personName)
	balance, _ := account.NewMoney(0)

	_, err := account.NewAccount("", owner, balance)

	if !errors.Is(err, account.ErrEmptyCustomerID) {
		t.Errorf("expected error %v, got %v", account.ErrEmptyCustomerID, err)
	}
}

func TestNewAccount_TwoAccounts_ShouldHaveDifferentIDs(t *testing.T) {
	personName, _ := customer.NewPersonName("John", "Doe")
	owner := account.NewOwnerSnapshot(personName)
	balance, _ := account.NewMoney(0)

	acc1, _ := account.NewAccount(uuid.New().String(), owner, balance)
	acc2, _ := account.NewAccount(uuid.New().String(), owner, balance)

	if acc1.ID() == acc2.ID() {
		t.Error("expected different IDs for two separate accounts, got same ID")
	}
}
