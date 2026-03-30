# REVIEW

# Why `errors.Is` returns false

`errors.Is(err, target)` matches when one of these is true:

1.`err == target` (same concrete error value for comparable errors), or
2.`err` has `Is(error) bool` and it says true, or
3.any wrapped inner error matches.

In the buggy code:

1.`ErrInsufficientBalance` is a `WalletError{Code:"E001", Message:"insufficient balance"}` value.
2`Withdraw` returns a *different* `WalletError` value with the same code but different message (`"need 500, have 100"`).

Because `WalletError` does not implement `Is(error) bool`, `errors.Is` falls back to direct value equality. The values are not equal (different `Message`), so matching fails.

# Exact fix (corrected code)

Make a structured insufficient-balance error that carries context and implements `Is` against the sentinel:

```go
package wallet

import "errors"

var ErrInsufficientBalance = errors.New("insufficient balance")

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
```

And return it from `Withdraw`:

```go
func (w *Wallet) Withdraw(amount int64) error {
	if amount > w.balance {
		return &InsufficientBalanceError{
			Required:  amount,
			Available: w.balance,
		}
	}
	w.balance -= amount
	return nil
}
```

Now:

- `errors.Is(err, ErrInsufficientBalance)` works.
- `errors.As(err, &ibErr)` extracts `Required` and `Available`.
