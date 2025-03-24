package transaction

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Transaction struct {
	Sender    string
	Receiver  string
	Amount    float64
	Signature string
}

func SignTransaction(tx *Transaction, privKey *ecdsa.PrivateKey) {
	data := tx.Sender + tx.Receiver + fmt.Sprintf("%f", tx.Amount)
	hash := sha256.Sum256([]byte(data))
	r, s, _ := ecdsa.Sign(rand.Reader, privKey, hash[:])
	tx.Signature = hex.EncodeToString(append(r.Bytes(), s.Bytes()...))
}

func VerifyTransaction(tx Transaction) bool {
	data := tx.Sender + tx.Receiver + fmt.Sprintf("%f", tx.Amount)
	hash := sha256.Sum256([]byte(data))

	sig, _ := hex.DecodeString(tx.Signature)
	r := new(big.Int).SetBytes(sig[:len(sig)/2])
	s := new(big.Int).SetBytes(sig[len(sig)/2:])

	pubKeyBytes, _ := hex.DecodeString(tx.Sender)
	x := new(big.Int).SetBytes(pubKeyBytes[:len(pubKeyBytes)/2])
	y := new(big.Int).SetBytes(pubKeyBytes[len(pubKeyBytes)/2:])
	pubKey := ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}

	return ecdsa.Verify(&pubKey, hash[:], r, s)
}
