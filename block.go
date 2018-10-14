package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	version      int
	index        int
	timestamp    string
	hash         string
	previousHash string
	merkleroot   string
	tx           []string
	nonce        int
	difficulty   float64
}

/* Calculating the block hash (SHA256) of the block from the following data
1. The Bitcoin version number.
2. The previous block hash.
3. The Merkle Root of all the transactions selected to be in that block.
4. The timestamp.
5. The difficulty target.
6. The Nonce.
*/
func calculateBlockHash(block Block) string {
	hashInput := string(block.version) + block.previousHash +
		block.merkleroot + block.timestamp + fmt.Sprint(block.difficulty) + string(block.nonce)
	hash := sha256.New()
	hash.Write([]byte(hashInput))

	return hex.EncodeToString(hash.Sum(nil))
}
