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

func TestNewNikInvalidType_ShouldFail(t *testing.T) {
	_, err := customer.NewNik("adfvG74389312839")
	if err == nil {
		t.Fatalf("NewNik() error = %v, want error", err)
	}

	if !errors.Is(err, customer.ErrInvalidType) {
		t.Errorf("NewNik() error = %v, want ErrInvalidType", err)
	}
}

func TestNewNikInvalidLength_ShouldFail(t *testing.T) {
	_, err := customer.NewNik("317598374673888")
	if err == nil {
		t.Fatalf("NewNik() error = %v, want error", err)
	}

	if !errors.Is(err, customer.ErrInvalidLength) {
		t.Errorf("NewNik() error = %v, want ErrInvalidLength", err)
	}
}

func TestNewNikEmpty_ShouldFail(t *testing.T) {
	_, err := customer.NewNik("")
	if err == nil {
		t.Fatalf("NewNik() error = %v, want error", err)
	}

	if !errors.Is(err, customer.ErrInvalidNik) {
		t.Errorf("NewNik() error = %v, want ErrInvalidNik", err)
	}
}

func TestNewNikTooLong_ShouldFail(t *testing.T) {
	_, err := customer.NewNik("31759837467388871") // 17 digit
	if !errors.Is(err, customer.ErrInvalidLength) {
		t.Errorf("NewNik() error = %v, want ErrInvalidLength", err)
	}
}
