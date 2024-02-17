package pointerserrors

import "testing"

func TestDeposite(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposite(Etherneum(10))
	wallet.Deposite(Etherneum(10))

	want := Etherneum(20)

	assertWalletBalance(t, wallet, want)
}

func TestWithdraw(t *testing.T) {
	t.Run("withdraw correct amount of ETH", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(Etherneum(20))
		err := wallet.Withdraw(Etherneum(10))

		want := Etherneum(10)

		assertNoError(t, err)
		assertWalletBalance(t, wallet, want)
	})

	t.Run("withdraw too much ETH", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(Etherneum(20))
		err := wallet.Withdraw(Etherneum(30))

		want := Etherneum(20)

		assertError(t, err, ErrInsufficientFunds)
		assertWalletBalance(t, wallet, want)
	})
}

func assertWalletBalance(t testing.TB, w Wallet, want Etherneum) {
	// With this in the code (t.Helper) when assert here do error
	// go test shows error at place where this func is called
	// not in this function body t.Errorf
	t.Helper()
	got := w.Balance()
	if want != got {
		t.Errorf("expected %s but got %s", want, got)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		// t.Fatal is like t.Errorf + return (?)
		t.Fatal("expected to get error but did not")
	}
	if want.Error() != got.Error() {
		t.Errorf("exprected %q but got %q", want, got)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("expected no error but got %q", got)
	}
}
