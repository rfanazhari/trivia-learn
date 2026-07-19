package account

import "learn-trivia/internal/customer"

type OwnerSnapshot struct {
}

func NewOwnerSnapshot(name customer.PersonName) *OwnerSnapshot {
	return &OwnerSnapshot{}
}
