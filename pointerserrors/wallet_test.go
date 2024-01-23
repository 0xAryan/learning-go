package pointerserrors

import (
	"testing"
)


func TestWallet(t *testing.T) {


	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{Bitcoin(0)}
		wallet.Deposit(100)
		assertBalance(t, wallet, Bitcoin(100))

	})

	t.Run("widthdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(100)}
		err := wallet.Widthdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(90))
	})

	t.Run("Widthdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Widthdraw(Bitcoin(30))

		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)

	})

}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Error("Got an error but didn't need one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("Got %s, want %s", got, want)
	}
}