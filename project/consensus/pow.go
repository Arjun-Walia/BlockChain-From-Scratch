package consensus

import "blockchain/block"

const Difficulty = 4

func MineBlock(b *block.Block) {
	target := string(make([]byte, Difficulty))
	for i := 0; i < Difficulty; i++ {
		target += "0"
	}
	for {
		b.Hash = block.CalculateHash(*b)
		if b.Hash[:Difficulty] == target {
			break
		}
		b.Nonce++
	}
}
