package blockchain

import (
	"errors"
	"fmt"
)

type Blockchain struct {
	chain []Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{
		chain: []Block{NewGenesisBlock()},
	}
}

func (bc *Blockchain) AddBlock(data string) Block {
	lastBlock := bc.chain[len(bc.chain)-1]
	newBlock := NewBlock(
		lastBlock.Index+1,
		data,
		lastBlock.Hash,
	)

	bc.chain = append(bc.chain, newBlock)
	return newBlock
}

func (bc *Blockchain) GetChain() []Block {
	return bc.chain
}

func (bc *Blockchain) GetBlock(index uint64) (Block, error) {
	for _, block := range bc.chain {
		if block.Index == index {
			return block, nil
		}
	}
	return Block{}, errors.New("Block not found")
}

func (bc *Blockchain) ValidateChain() error {

	if len(bc.chain) == 0 {
		return errors.New("the chain is empty")
	}

	// Check the rest of the chain
	for i := 1; i < len(bc.chain); i++ {
		actualBlock := bc.chain[i]
		previousBlock := bc.chain[i-1]

		// Check if the actual~previous hash linking is valid
		if actualBlock.PrevHash != previousBlock.Hash {
			return fmt.Errorf(
				"the chain is invalid, the %d block has an invalid previous hash",
				actualBlock.Index,
			)
		}

		// Check if the actual hash is valid
		if actualBlock.Hash != actualBlock.CalculateHash() {
			return fmt.Errorf(
				"the chain is invalid, the %d block has an invalid hash",
				actualBlock.Index,
			)
		}
	}

	return nil
}
