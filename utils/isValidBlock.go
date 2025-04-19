package utils

import (
	"blockchain_votation_system/entities"
)

func IsBlockValid(newBlock entities.Block, prevBlock entities.Block) bool {
	if newBlock.Index != prevBlock.Index+1 {
		return false
	}

	if newBlock.PrevHash != prevBlock.Hash {
		return false
	}

	if newBlock.Hash != CalculateHash(&newBlock) {
		return false
	}

	return true
}
