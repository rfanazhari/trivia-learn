package account_test

import (
	"learn-trivia/internal/account"
	"testing"
)

func TestNewAccount_ShouldSuccess(t *testing.T) {
	owner, _ := account.NewOwner("John", "Doe")

	acc, err := account.NewAccount(owner)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if acc.ID() == "" {
		t.Errorf("expected ID to be non-empty, got %s", acc.ID())
	}

	if acc.Owner().GetFirstName() != "John" {
		t.Errorf("expected first name to be John, got %s", acc.Owner().GetFirstName())
	}

	if acc.Owner().GetLastName() != "Doe" {
		t.Errorf("expected last name to be Doe, got %s", acc.Owner().GetLastName())
	}

	if acc.Balance().Amount() != 0 {
		t.Errorf("expected initial balance 0, got: %d", acc.Balance().Amount())
	}

	if acc.Status() != account.StatusPending {
		t.Errorf("expected status Pending, got: %v", acc.Status())
	}
}

func TestNewAccount_TwoAccounts_ShouldHaveDifferentIDs(t *testing.T) {
	owner, _ := account.NewOwner("Budi", "")

	acc1, _ := account.NewAccount(owner)
	acc2, _ := account.NewAccount(owner)

	if acc1.ID() == acc2.ID() {
		t.Error("expected different IDs for two separate accounts, got same ID")
	}
}
