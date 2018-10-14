package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
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

func createGenesisBlock() Block {

	// Generating new seed for the nonce
	rand.Seed(time.Now().UTC().UnixNano())

	genesisBlock := Block{
		version:      os.Getenv("block_chain_version"),
		index:        0,
		timestamp:    time.Now().String(),
		hash:         "",
		previousHash: "0",
		merkleroot:   "", // TODO: TRANSACTION INTERFACE
		tx:           []string{},
		nonce:        10000000000 + rand.Intn(9999999999-1000000000), // Horrible solution for nonce generation
		difficulty:   2}

	// Calculating the hash for the genesis block
	genesisBlock.hash = calculateBlockHash(genesisBlock)
	return genesisBlock
}
