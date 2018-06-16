package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want BitCoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got.Error(), want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		wallet.Withdraw(10)
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := BitCoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(BitCoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
