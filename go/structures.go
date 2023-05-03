package main

// Create the Block data structure
type Block struct {
	Timestamp         int64  // the time when the block was created
	PreviousBlockHash []byte // the hash of the previous block
	MyBlockHash       []byte // the hash of the current block
	AllData           []byte // the data or transactions (body info)
}

// Prepare the Blockchain data structure (a series of blocks):
type Blockchain struct {
	Blocks []*Block
}
