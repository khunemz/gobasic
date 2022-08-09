package services

import "errors"

var (
	ErrZeroAmount = errors.New("Cannot be zero amount")
	ErrRepository = errors.New("Unexpected error at repository level")
)
