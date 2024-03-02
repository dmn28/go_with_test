package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	ballance Bitcoin
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.ballance {
		return ErrInsufficientFunds
	}

	w.ballance -= amount
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.ballance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.ballance += amount
}
