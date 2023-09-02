package main

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("Insufficient funds")

type Wallet struct {
	balance Bitcoin // privado en el paquete main
}

type Bitcoin int

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
