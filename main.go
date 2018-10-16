package main

import (
	"fmt"
	"time"
)

func main() {
	BlockChain = initializeBlockChain()
	newBlock1 := Block{"1", len(BlockChain), time.Now().String(), "noHash", "prevHash", "notx", []string{"notx"}, 0, 4}
	newBlock2 := Block{"1", len(BlockChain), time.Now().String(), "noHash", "prevHash", "notx", []string{"notx"}, 0, 4}

	appendBlock(newBlock1)
	appendBlock(newBlock2)

	for i := 0; i < len(BlockChain); i++ {
		fmt.Println("HASH: ", BlockChain[i].hash)
	}

	// calculateBlockHash(Block{"1", 0, "12:00:00", "NoHash", "NoHash", "NoMerkelRoot", []string{"A"}, 12, 1})
	// fmt.Println("Simple Block Chain Implementation")
	fmt.Println(verifyBlockchainIntegrity(BlockChain))
}
