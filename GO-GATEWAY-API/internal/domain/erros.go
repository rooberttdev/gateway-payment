package domain

import "errors"

var (
	ErrAccountNotFound = errors.New("account not found")
	ErrDuplicateAPIKey = errors.New("duplicate API key")
	ErrInvoiceNotFound = errors.New("invoice not found")
	ErrUnauthorizedAcess = errors.New("unauthorized acess")
	ErrInvalidAmount = errors.New("invalid amount")
	ErrInvalidStatus = errors.New("invalid status")
)
