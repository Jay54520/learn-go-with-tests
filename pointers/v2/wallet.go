package main

type Wallet struct {
	balance int
}

func (w *Wallet)  Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() (balance int)  {
	return w.balance
}