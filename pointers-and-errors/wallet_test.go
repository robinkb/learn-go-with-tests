package wallet

import (
	"errors"
	"testing"
)

func TestBitcoin(t *testing.T) {
	got := Bitcoin(10).String()
	want := "10 BTC"

	if got != want {
		t.Errorf("Bitcoin(10).String() = %s, want %s", got, want)
	}
}

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(5)
		withdrawAmount := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(withdrawAmount)

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, w Wallet, want Bitcoin) {
	t.Helper()

	got := w.Balance()

	if got != want {
		t.Errorf("wallet.Balance() = %s, want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("wanted an error but did not get one")
	}

	if !errors.Is(got, want) {
		t.Errorf("unexpected error: got %q, want %q", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatalf("got an error but did not want one: %s", got)
	}
}
