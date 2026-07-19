package account

import "github.com/google/uuid"

type Account struct {
	id      string
	owner   Owner
	balance Money
	status  AccountStatus
}

func NewAccount(owner Owner) (*Account, error) {
	newBalance, err := NewMoney(0)
	if err != nil {
		return nil, err
	}
	return &Account{
		id:      uuid.New().String(),
		owner:   owner,
		balance: newBalance,
		status:  StatusPending,
	}, nil
}

func (a *Account) ID() string {
	return a.id
}

func (a *Account) Owner() Owner {
	return a.owner
}

func (a *Account) Balance() Money {
	return a.balance
}

func (a *Account) Status() AccountStatus {
	return a.status
}
