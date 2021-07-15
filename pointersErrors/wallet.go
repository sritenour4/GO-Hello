package pointersErrors

import (
	"errors"
	"fmt"
)

type Bitcoin int

// lowercase field to keep it private outside the package it's defined in
type Wallet struct {
	balance Bitcoin
}

func NewWallet(initialBalance Bitcoin) *Wallet{
	return &Wallet{
		balance: initialBalance,
	}
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
