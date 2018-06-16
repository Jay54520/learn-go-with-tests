package main

type BitCoin int

type Wallet struct {
	balance BitCoin
}

func (w *Wallet)  Deposit(amount BitCoin) {
	w.balance += amount
}

func (w *Wallet) Balance() (balance BitCoin)  {
	return w.balance
}