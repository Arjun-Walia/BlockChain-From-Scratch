package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func ComputeSHA256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
