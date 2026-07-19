package account

import "errors"

var (
	ErrInvalidFirstName  = errors.New("invalid first name")
	ErrFirstNameTooShort = errors.New("first name too short")
	ErrNegativeAmount    = errors.New("negative amount")
)
