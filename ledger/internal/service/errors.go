package service

import "errors"

var (
	ErrValidation     = errors.New("validation error")
	ErrBudgetExceeded = errors.New("budget exceeded")
)

func IsValidationError(err error) bool {
	return errors.Is(err, ErrValidation)
}

func IsBudgetExceeded(err error) bool {
	return errors.Is(err, ErrBudgetExceeded)
}
