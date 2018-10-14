package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Block struct {
	version      string
	index        int
	timestamp    string
	hash         string
	previousHash string
	merkleroot   string
	tx           []string
	nonce        int
	difficulty   int
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

// Mine new block by finding the proper nonce and hash according to the difficulty level
func mineBlock(block Block, difficulty int) (int, string) {
	nonce := 0
	for string([]rune(block.hash)[0:difficulty]) != strings.Repeat("0", difficulty) {
		block.nonce = nonce
		block.hash = calculateBlockHash(block)
		nonce++
	}

	return block.nonce, block.hash
}

func generateRandomNumber() int {
	// Generating new seed for the nonce
	rand.Seed(time.Now().UTC().UnixNano())
	return 10000000000 + rand.Intn(9999999999-1000000000)
}
