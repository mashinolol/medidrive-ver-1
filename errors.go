package wallet

import "errors"

var (
	ErrInsufficientBalance = errors.New("insufficient balance")
	ErrInvalidAmount       = errors.New("invalid amount")
	ErrWalletFrozen        = errors.New("wallet is frozen")
)

type InsufficientBalanceError struct {
	Required  int64
	Available int64
}

func (e *InsufficientBalanceError) Error() string {
	return "insufficient balance"
}

func (e *InsufficientBalanceError) Is(target error) bool {
	return target == ErrInsufficientBalance
}
