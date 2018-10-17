package main

import (
	"os"
	"strings"
	"time"
)

var BlockChain []Block

// Generate a BlockChain that includes the genesis block
func initializeBlockChain() []Block {
	BlockChain = append(BlockChain, createGenesisBlock())
	return BlockChain
}

func createGenesisBlock() Block {

	genesisBlock := Block{
		version:      os.Getenv("block_chain_version"),
		index:        0,
		timestamp:    time.Now().String(),
		hash:         "",
		previousHash: strings.Repeat("0", 64), // String consisting of 64 zeros
		merkleroot:   "",                      // TODO: TRANSACTION INTERFACE
		tx:           []string{},
		nonce:        generateRandomNumber(), // Horrible solution for nonce generation
		difficulty:   4}

	// Calculating the hash for the genesis block
	genesisBlock.hash = calculateBlockHash(genesisBlock)

	return genesisBlock
}

func appendBlock(block Block) {

	block = mineBlock(block, 4)
	BlockChain = append(BlockChain, block)
}

// Iterate through the chain to verify if the integrity of the blockchain holds
func verifyBlockchainIntegrity(Blockchain []Block) bool {

	for i := 1; i < len(Blockchain); i++ {
		if calculateBlockHash(Blockchain[i-1]) != BlockChain[i].previousHash {
			return false
		}
	}

	return true
}
