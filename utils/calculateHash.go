package utils

import (
	"blockchain_votation_system/entities"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

func CalculateHash(block *entities.Block) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash
	hash := sha256.New()

	hash.Write([]byte(record))

	hashResult := hex.EncodeToString(hash.Sum(nil))

	return hashResult
}
