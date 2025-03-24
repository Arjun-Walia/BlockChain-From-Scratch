package utils

import (
	"blockchain/block"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func ComputeMerkleRoot(transactions []block.Transaction) string {
	if len(transactions) == 0 {
		return ""
	}

	var hashes [][]byte
	for _, tx := range transactions {
		data := tx.Sender + tx.Receiver + fmt.Sprintf("%f", tx.Amount)
		hash := sha256.Sum256([]byte(data))
		hashes = append(hashes, hash[:])
	}

	for len(hashes) > 1 {
		var newHashes [][]byte
		for i := 0; i < len(hashes); i += 2 {
			if i+1 >= len(hashes) {
				newHashes = append(newHashes, hashes[i])
				continue
			}
			combined := append(hashes[i], hashes[i+1]...)
			hash := sha256.Sum256(combined)
			newHashes = append(newHashes, hash[:])
		}
		hashes = newHashes
	}
	return hex.EncodeToString(hashes[0])
}
