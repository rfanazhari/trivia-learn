package account

import "errors"

var (
	ErrNegativeAmount  = errors.New("negative amount")
	ErrEmptyCustomerID = errors.New("customer ID must not be empty")
)
