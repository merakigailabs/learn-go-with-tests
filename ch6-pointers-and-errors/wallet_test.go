package pointersanderrors

import "testing"

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit Bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Widthraw Bitcoin", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Widthraw(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

}
