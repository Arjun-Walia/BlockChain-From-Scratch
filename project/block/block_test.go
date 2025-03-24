package block

import (
	"testing"
	"time"
)

func TestCalculateHash(t *testing.T) {
	b := Block{
		Index:      1,
		Timestamp:  time.Now(),
		PrevHash:   "0",
		MerkleRoot: "test",
		Nonce:      0,
	}
	hash := CalculateHash(b)
	if len(hash) != 64 {
		t.Error("Invalid SHA-256 hash length")
	}
}

func TestMineBlock(t *testing.T) {
	b := Block{
		Index:      1,
		Timestamp:  time.Now(),
		PrevHash:   "0",
		MerkleRoot: "test",
		Nonce:      0,
	}
	b.Mine(4)
	if b.Hash[:4] != "0000" {
		t.Error("Block was not mined correctly")
	}
}
