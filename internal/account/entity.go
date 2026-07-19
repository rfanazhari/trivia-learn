package account

import "github.com/google/uuid"

type Account struct {
	id         string
	customerId string
	owner      OwnerSnapshot
	balance    Money
	status     AccountStatus
}

func NewAccount(customerID string, personName OwnerSnapshot, balance Money) (*Account, error) {
	if customerID == "" {
		return nil, ErrCustomerIDNotEmpty
	}

	return &Account{
		id:         uuid.New().String(),
		customerId: customerID,
		owner:      personName,
		balance:    balance,
		status:     StatusPending,
	}, nil
}

func (a *Account) ID() string {
	return a.id
}

func (a *Account) CustomerID() string {
	return a.customerId
}

func (a *Account) Owner() OwnerSnapshot {
	return a.owner
}

func (a *Account) Balance() Money {
	return a.balance
}

func (a *Account) Status() AccountStatus {
	return a.status
}
