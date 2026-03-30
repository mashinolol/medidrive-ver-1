# Wallet Entity Assessment

This repository contains a solution for the Backend Developer Assessment (Junior, Task 2).

## Implemented

- `Wallet` domain entity with fields:
  - `id` (`WalletID`)
  - `ownerID` (`OwnerID`)
  - `balance` (in cents, `int64`)
  - `status` (`ACTIVE`, `FROZEN`)
- Methods:
  - `Deposit(amount int64) error`
  - `Withdraw(amount int64) error`
  - `Freeze() error`
- Business rules:
  - amount must be positive (`> 0`) for deposit and withdraw
  - operations are blocked for frozen wallets
  - withdrawal cannot exceed available balance
- Domain errors:
  - sentinel errors (`ErrInsufficientBalance`, `ErrInvalidAmount`, `ErrWalletFrozen`)
  - structured `InsufficientBalanceError` compatible with `errors.Is` and `errors.As`

## Files

- `wallet.go` - wallet entity implementation
- `errors.go` - domain error definitions
- `wallet_test.go` - unit tests
- `REVIEW.md` - bug analysis (`errors.Is` mismatch explanation + fix)
- `ANSWERS.md` - answers to conceptual questions

## Run tests

```bash
go test ./...
```
