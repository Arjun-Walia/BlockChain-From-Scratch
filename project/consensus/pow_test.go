package consensus

import (
	"blockchain/block"
	"testing"
	"time"
)

func TestMineBlock(t *testing.T) {
	b := block.Block{
		Index:      1,
		Timestamp:  time.Now(),
		PrevHash:   "0",
		MerkleRoot: "test",
		Nonce:      0,
	}
	MineBlock(&b)
	if b.Hash[:4] != "0000" {
		t.Error("Block was not mined correctly")
	}
}
