package main

import (
	"fmt"
)

func main() {
	BlockChain = initializeBlockChain()
	newBlock := Block{"1", len(BlockChain), "12312321313", "noHash", "prevHash", "notx", []string{"notx"}, 0, 2}
	appendBlock(newBlock)
	fmt.Println(BlockChain[0].hash)
	fmt.Println(BlockChain[1].hash)
	// calculateBlockHash(Block{"1", 0, "12:00:00", "NoHash", "NoHash", "NoMerkelRoot", []string{"A"}, 12, 1})
	fmt.Println("Simple Block Chain Implementation")
}
