package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Blockchain struct {
	chain      []Block
	mine_hash  string
	difficulty uint64
}

func NewBlockchain() *Blockchain {
	genesis := NewGenesisBlock()
	h := sha256.New()
	h.Write([]byte(genesis.Hash))
	mine_hash := fmt.Sprintf("%x", h.Sum(nil))

	return &Blockchain{
		chain:      []Block{genesis},
		mine_hash:  mine_hash,
		difficulty: 3,
	}
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

func (bc *Blockchain) Mine(mine_id uint64, data string) *Block {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%d", mine_id)))
	mine_id_hash := fmt.Sprintf("%x", h.Sum(nil))

	hash := mine_id_hash[:bc.difficulty]

	if strings.HasPrefix(bc.mine_hash, hash) {

		block := NewBlock(
			uint64(len(bc.chain)),
			data,
			bc.chain[len(bc.chain)-1].Hash,
		)

		bc.chain = append(bc.chain, block)

		// set the new mine_hash based on a random number
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		bc.mine_hash = fmt.Sprintf("%x", r.Int63())

		return &block
	}

	return nil

}

func (bc *Blockchain) SetDifficulty(difficulty uint64) {
	bc.difficulty = difficulty
}
