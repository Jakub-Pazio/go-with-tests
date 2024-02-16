package pointerserrors

import (
	"errors"
	"fmt"
)

type Etherneum int

var ErrInsufficientFunds = errors.New("insufficient balance on account to perform operation")

type Wallet struct {
	balance Etherneum
}

func (e Etherneum) String() string {
	return fmt.Sprintf("%d ETH", e)
}

func (w *Wallet) Deposite(amount Etherneum) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Etherneum) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w Wallet) Balance() Etherneum {
	return w.balance
}
