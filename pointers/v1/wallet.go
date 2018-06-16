package main

import "fmt"

type Wallet struct {
	balance int
}

func (w Wallet)  Deposit(amount int) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() (balance int)  {
	return w.balance
}