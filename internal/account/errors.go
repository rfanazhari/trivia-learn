package account

import "errors"

var (
	ErrNegativeAmount     = errors.New("negative amount")
	ErrCustomerIDNotEmpty = errors.New("customer ID must not be empty")
)
