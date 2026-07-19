package account

import "learn-trivia/internal/customer"

type OwnerSnapshot struct {
	firstName string
	lastName  string
}

func NewOwnerSnapshot(name customer.PersonName) OwnerSnapshot {
	return OwnerSnapshot{
		firstName: name.FirstName(),
		lastName:  name.LastName(),
	}
}

func (s OwnerSnapshot) FirstName() string {
	return s.firstName
}

func (s OwnerSnapshot) LastName() string {
	return s.lastName
}
