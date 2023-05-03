package main

import (
	"fmt"
)

func main() {
	// Initialize the blockchain
	newblockchain := NewBlockchain()

	// create 2 blocks and add 2 transactions
	newblockchain.AddBlock("first transaction")  // first block containing one tx
	newblockchain.AddBlock("Second transaction") // second block containing one tx

	// Now print all the blocks and their contents, iterating per block
	for i, block := range newblockchain.Blocks {
		fmt.Printf("Block ID : %d \n", i)
		fmt.Printf("Timestamp : %d \n", block.Timestamp+int64(i))
		fmt.Printf("Hash of the block %x\n", block.MyBlockHash)                 // print the hash of the block
		fmt.Printf("Hash of the previous Block: %x\n", block.PreviousBlockHash) // print the hash of the previous block
		fmt.Printf("All the transactions: %s\n", block.AllData)                 // print the transactions
	}
}
