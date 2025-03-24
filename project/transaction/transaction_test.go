package transaction

import (
	"blockchain/wallet"
	"testing"
)

func TestSignAndVerifyTransaction(t *testing.T) {
	w := wallet.NewWallet()
	tx := Transaction{
		Sender:   w.GetAddress(),
		Receiver: "receiver_address",
		Amount:   1.0,
	}
	SignTransaction(&tx, w.PrivateKey)
	if !VerifyTransaction(tx) {
		t.Error("Transaction verification failed")
	}
}
