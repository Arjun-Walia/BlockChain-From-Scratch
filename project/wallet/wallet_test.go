package wallet

import "testing"

func TestNewWallet(t *testing.T) {
	w := NewWallet()
	if w.PrivateKey == nil || w.PublicKey == nil {
		t.Error("Wallet was not created correctly")
	}
}

func TestGetAddress(t *testing.T) {
	w := NewWallet()
	address := w.GetAddress()
	if len(address) == 0 {
		t.Error("Wallet address is empty")
	}
}
