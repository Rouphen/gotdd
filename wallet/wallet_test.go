package wallet

import "testing"

func TestWallet(t *testing.T) {
	checkWallet := func(t *testing.T, got, want Bitcoin) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		bt := Bitcoin(10)
		wallet.Deposit(bt)

		got := wallet.Balance()
		want := bt

		checkWallet(t, got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)

		got := wallet.Balance()
		want := Bitcoin(10)

		checkWallet(t, got, want)
	})
}
