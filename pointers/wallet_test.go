package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, e Bitcoin) {
		t.Helper()
		a := wallet.Balance()

		if a != e {
			t.Errorf("expected: %s actual: %s", e, a)
		}
	}

	assertError := func(t *testing.T, e, a error) {
		t.Helper()

		if a != e {
			t.Errorf("Expected error: %q actual: %q", e, a)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		e := Bitcoin(10)
		assertBalance(t, wallet, e)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		e := Bitcoin(10)
		assertBalance(t, wallet, e)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}
