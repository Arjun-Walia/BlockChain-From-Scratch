package block

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index        int
	Timestamp    time.Time
	Transactions []Transaction
	PrevHash     string
	MerkleRoot   string
	Nonce        int
	Hash         string
}

type Transaction struct {
	Sender    string
	Receiver  string
	Amount    float64
	Signature string
}

func CalculateHash(b Block) string {
	record := strconv.Itoa(b.Index) + b.Timestamp.String() + b.PrevHash + b.MerkleRoot + strconv.Itoa(b.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (b *Block) Mine(difficulty int) {
	target := string(make([]byte, difficulty))
	for i := 0; i < difficulty; i++ {
		target += "0"
	}
	for {
		b.Hash = CalculateHash(*b)
		if b.Hash[:difficulty] == target {
			break
		}
		b.Nonce++
	}
}
