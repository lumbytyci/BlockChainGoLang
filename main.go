package main

import (
	"fmt"
)

var blockChain []Block

func main() {
	calculateBlockHash(Block{"1", 0, "12:00:00", "NoHash", "NoHash", "NoMerkelRoot", []string{"A"}, 12, 1})
	fmt.Println(createGenesisBlock().nonce)
	fmt.Println("Simple Block Chain Implementation")
}
