package wallet

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	bt := Bitcoin(10)
	wallet.Deposit(bt)

	got := wallet.Balance()
	want := bt

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}
