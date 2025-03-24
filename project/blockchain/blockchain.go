package blockchain

import (
	"blockchain/block"
	"blockchain/utils"
	"time"
)

type Blockchain struct {
	Blocks       []*block.Block
	PendingTrans []block.Transaction
	Difficulty   int
}

func CreateGenesisBlock() *block.Block {
	return &block.Block{
		Index:        0,
		Timestamp:    time.Now(),
		Transactions: []block.Transaction{},
		PrevHash:     "0",
		MerkleRoot:   "",
	}
}

func (bc *Blockchain) AddBlock(newBlock *block.Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, CreateGenesisBlock())
	}
	newBlock.MerkleRoot = utils.ComputeMerkleRoot(newBlock.Transactions)
	newBlock.Mine(bc.Difficulty)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		prevBlock := bc.Blocks[i-1]

		if currentBlock.Hash != block.CalculateHash(*currentBlock) {
			return false
		}
		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
	}
	return true
}
