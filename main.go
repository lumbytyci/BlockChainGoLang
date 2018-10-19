package main

import (
	"fmt"
	"time"
)

func main() {
	BlockChain = initializeBlockChain()
	newBlock1 := Block{"1", len(BlockChain), time.Now().String(), "noHash", "prevHash", "notx", []string{"notx"}, 0, 4}
	newBlock2 := Block{"1", len(BlockChain), time.Now().String(), "noHash", "prevHash", "notx", []string{"notx"}, 0, 4}
	newBlock3 := Block{"1", len(BlockChain), time.Now().String(), "noHash", "prevHash", "notx", []string{"notx"}, 0, 4}
	appendBlock(newBlock1)
	appendBlock(newBlock2)
	appendBlock(newBlock3)

	for i := 0; i < len(BlockChain); i++ {
		fmt.Println("HASH: ", BlockChain[i].hash)
	}

	// fmt.Println("Simple Block Chain Implementation")
	fmt.Println(verifyBlockchainIntegrity(BlockChain))
}
