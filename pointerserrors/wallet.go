package pointerserrors

import (
	"fmt"
	"errors"
)

var ErrInsufficientFunds = errors.New("cannot widthdraw, insufficent funds")

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}



type Wallet struct{
	balance Bitcoin
}


func(w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("Address of balance in deposit is %p \n", &w.balance)
	(*w).balance += amount
}

func(w *Wallet) Widthdraw(amount Bitcoin) error{

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

func(w *Wallet) Balance() Bitcoin{
	return w.balance
}