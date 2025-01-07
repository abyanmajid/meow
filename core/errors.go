package meow

import "errors"

var (
	ErrInvalidString  = errors.New("not a valid string")
	ErrInvalidInteger = errors.New("not a valid integer")
	ErrInvalidFloat   = errors.New("not a valid float")
	ErrInvalidBoolean = errors.New("not a valid boolean")
	ErrInvalidDate    = errors.New("not a valid date")
	ErrInvalidNil     = errors.New("not null")
	ErrInvalidAny     = errors.New("is null")
	ErrInvalidNever   = errors.New("never allowed")
)
