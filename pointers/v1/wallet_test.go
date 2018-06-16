package main

import (
	"testing"
	"fmt"
)

func TestWallet(t *testing.T)  {

	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	fmt.Printf("address of blance in test is %v \n", &wallet.balance)
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
