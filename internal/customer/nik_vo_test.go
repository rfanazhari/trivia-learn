package customer_test

import (
	"errors"
	"learn-trivia/internal/customer"
	"testing"
)

func TestNewNik_ShouldSuccess(t *testing.T) {
	nik, err := customer.NewNik("3175983746738887")
	if err != nil {
		t.Errorf("NewNik() error = %v, want nil", err)
	}
	if nik.String() != "3175983746738887" {
		t.Errorf("NewNik() got = %v, want 3175983746738887", nik.String())
	}
}

func TestNewNikEmpty_ShouldFail(t *testing.T) {
	nik, err := customer.NewNik("")
	if err == nil {
		t.Fatalf("NewNik() error = %v, want error", err)
	}
	if nik != nil {
		t.Errorf("NewNik() got = %v, want nil", nik)
	}

	if !errors.Is(err, ErrInvalidNik) {
		t.Errorf("NewNik() error = %v, want ErrInvalidNik", err)
	}
}
