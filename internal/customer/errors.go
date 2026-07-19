package customer

import "errors"

var (
	ErrInvalidNik    = errors.New("invalid nik")
	ErrInvalidLength = errors.New("invalid length")
)
