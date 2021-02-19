package pointer_and_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(15))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoErrors(t, err)
	})

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(15)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(20))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds.Error())
	})
}

func assertNoErrors(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but did'nt want one")
	}
}
