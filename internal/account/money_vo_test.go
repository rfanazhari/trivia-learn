package account_test

import (
	"errors"
	"learn-trivia/internal/account"
	"testing"
)

func TestNewMoney_ZeroAmount_ShouldSuccess(t *testing.T) {
	money, err := account.NewMoney(0)
	if err != nil {
		t.Fatalf("expected no error with amount 0, but error: %v", err)
	}
	if money.Amount() != 0 {
		t.Errorf("expected amount 0, got %d", money.Amount())
	}
}

func TestNewMoney_PositiveAmount_ShouldSuccess(t *testing.T) {
	money, err := account.NewMoney(100)
	if err != nil {
		t.Fatalf("expected no error with amount 100, but error: %v", err)
	}
	if money.Amount() != 100 {
		t.Errorf("expected amount 100, got %d", money.Amount())
	}
}

func TestNewMoney_NegativeAmount_ShouldFail(t *testing.T) {
	_, err := account.NewMoney(-100)
	if err == nil {
		t.Fatalf("expected error with amount -100, but no error")
	}
	if !errors.Is(err, account.ErrNegativeAmount) {
		t.Errorf("expected error %v, got %v", account.ErrNegativeAmount, err)
	}
}
