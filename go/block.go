package main

import (
	"bytes"         // need to convert data into byte in order to be sent on the network, computer understands better the byte(8bits)language
	"crypto/sha256" // crypto library to hash the data
	"strconv"       // for conversion
	"time"          // the time for our timestamp
)

// Create a function for new block generation and return that block
func NewBlock(data string, prevBlockHash []byte) *Block {
	// instance and hash block
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)}
	block.SetHash()
	return block
}

/* let's now create the genesis block function that will return the first block. The genesis block is the first block on the chain */
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// Now let's create a method for generating a hash of the block
// We will just concatenate all the data and hash it to obtain the block hash
func (block *Block) SetHash() {
	// get the time and convert it into a unique series of digits
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	// concatenate all the block data
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{})
	// hash the whole thing
	hash := sha256.Sum256(headers)

	// now set the hash of the block
	/// @dev the '[:]' are for casting from a fixed byte arr [32]byte to []byte type
	block.MyBlockHash = hash[:]
}

/* Create the function that returns the whole blockchain and add the genesis to it first. the genesis block is the first ever mined block, so let's create a function that will return it since it does not exist yet */
func NewBlockchain() *Blockchain { // the function is created
	// the genesis block is added first to the chain
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// create the method that adds a new block to a blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	// get the previous block
	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	// create a new block containing the data and the hash of the previous block
	newBlock := NewBlock(data, PreviousBlock.MyBlockHash)
	// append it to the blockchain
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}
