package blockchain

import (
	"blockchain/block"
	"testing"
)

func TestAddBlock(t *testing.T) {
	bc := Blockchain{
		Blocks:     []*block.Block{CreateGenesisBlock()},
		Difficulty: 4,
	}

	newBlock := &block.Block{
		Index:        1,
		Transactions: []block.Transaction{},
		PrevHash:     bc.Blocks[0].Hash,
	}
	bc.AddBlock(newBlock)

	if len(bc.Blocks) != 2 {
		t.Error("Block was not added to the blockchain")
	}
}

func TestIsValid(t *testing.T) {
	bc := Blockchain{
		Blocks:     []*block.Block{CreateGenesisBlock()},
		Difficulty: 4,
	}

	newBlock := &block.Block{
		Index:        1,
		Transactions: []block.Transaction{},
		PrevHash:     bc.Blocks[0].Hash,
	}
	bc.AddBlock(newBlock)

	if !bc.IsValid() {
		t.Error("Blockchain is not valid")
	}
}
