package causes

import (
	"errors"
)

var (
	ErrNotFound         = errors.New("not found")
	ErrInvalidData      = errors.New("invalid data")
	ErrAlreadyExists    = errors.New("already exists")
	ErrUnauthenticated  = errors.New("unauthenticated")
	ErrPermissionDenied = errors.New("permission denied")
)
