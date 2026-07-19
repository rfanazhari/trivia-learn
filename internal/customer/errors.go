package customer

import "errors"

var (
	ErrInvalidNik    = errors.New("nik must not be empty")
	ErrInvalidLength = errors.New("nik must be exactly 16 digits")
	ErrInvalidType   = errors.New("nik must be a numeric")
)
