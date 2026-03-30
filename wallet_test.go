package wallet

import (
	"errors"
	"testing"
)

func TestDeposit_Success(t *testing.T) {
	w := NewWallet("w1", "owner1", 100)

	err := w.Deposit(50)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if w.Balance() != 150 {
		t.Fatalf("expected balance 150, got %d", w.Balance())
	}
}

func TestDeposit_InvalidAmount(t *testing.T) {
	w := NewWallet("w1", "owner1", 100)

	err := w.Deposit(0)
	if !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("expected ErrInvalidAmount, got %v", err)
	}
}

func TestDeposit_FrozenWallet(t *testing.T) {
	w := NewWallet("w1", "owner1", 100)
	_ = w.Freeze()

	err := w.Deposit(10)
	if !errors.Is(err, ErrWalletFrozen) {
		t.Fatalf("expected ErrWalletFrozen, got %v", err)
	}
}

func TestWithdraw_Success(t *testing.T) {
	w := NewWallet("w1", "owner1", 100)

	err := w.Withdraw(40)
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if w.Balance() != 60 {
		t.Fatalf("expected balance 60, got %d", w.Balance())
	}
}

func TestWithdraw_InvalidAmount_CheckedBeforeFrozen(t *testing.T) {
	w := NewWallet("w1", "owner1", 100)
	_ = w.Freeze()

	err := w.Withdraw(-50)
	if !errors.Is(err, ErrInvalidAmount) {
		t.Fatalf("expected ErrInvalidAmount, got %v", err)
	}
}

func TestWithdraw_InsufficientBalance_IsAndAs(t *testing.T) {
	w := NewWallet("w1", "owner1", 100)

	err := w.Withdraw(150)
	if !errors.Is(err, ErrInsufficientBalance) {
		t.Fatalf("expected ErrInsufficientBalance, got %v", err)
	}

	var ibErr *InsufficientBalanceError
	if !errors.As(err, &ibErr) {
		t.Fatal("expected *InsufficientBalanceError via errors.As")
	}
	if ibErr.Required != 150 || ibErr.Available != 100 {
		t.Fatalf("unexpected fields: required=%d available=%d", ibErr.Required, ibErr.Available)
	}
}

func TestFreeze(t *testing.T) {
	w := NewWallet("w1", "owner1", 100)

	err := w.Freeze()
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
	if w.Status() != StatusFrozen {
		t.Fatalf("expected status %s, got %s", StatusFrozen, w.Status())
	}
}
