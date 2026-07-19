package account_test

import (
	"errors"
	"learn-trivia/internal/account"
	"testing"
)

func TestNewOwner_ValidateFirstNameOnly_ShouldSuccess(t *testing.T) {
	owner, err := account.NewOwner("John", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if owner.FirstName != "John" {
		t.Errorf("unexpected firstName: %s", owner.FirstName)
	}

	if owner.LastName != "" {
		t.Errorf("unexpected lastName: %s", owner.LastName)
	}
}

func TestNewOwner_ValidateFirstNameAndLastName_ShouldSuccess(t *testing.T) {
	owner, err := account.NewOwner("John", "Wick")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if owner.FirstName != "John" {
		t.Errorf("unexpected firstName: %s", owner.FirstName)
	}

	if owner.LastName != "Wick" {
		t.Errorf("unexpected lastName: %s", owner.LastName)
	}
}

func TestNewOwner_FirstNameEmpty_ShouldFail(t *testing.T) {
	_, err := account.NewOwner("", "Wick")
	if err == nil {
		t.Fatal("expected error for empty firstName, got nil")
	}

	if !errors.Is(err, account.ErrInvalidFirstName) {
		t.Errorf("expected ErrInvalidFirstName, got %v", err)
	}
}

func TestNewOwner_FirstNameTooShort_ShouldFail(t *testing.T) {
	_, err := account.NewOwner("Jo", "Wick")
	if err == nil {
		t.Fatal("expected error for firstName < 3 char, got nil")
	}

	if !errors.Is(err, account.ErrFirstNameTooShort) {
		t.Errorf("expected ErrFirstNameTooShort, got %v", err)
	}
}

func TestNewOwner_FirstNameExactlyMinLength_ShouldSuccess(t *testing.T) {
	_, err := account.NewOwner("Joh", "")
	if err != nil {
		t.Fatalf("expected no error with length 3, but error: %v", err)
	}
}
