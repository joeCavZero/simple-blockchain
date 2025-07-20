package blockchain

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/joeCavZero/simple-blockchain/src/logkit"
)

var bclk = logkit.NewLogkit("Blockchain")

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
		bclk.Error("The chain is empty")
		return errors.New("the chain is empty")
	}

	// Check the rest of the chain
	for i := 1; i < len(bc.chain); i++ {
		actualBlock := bc.chain[i]
		previousBlock := bc.chain[i-1]

		// Check if the actual~previous hash linking is valid
		if actualBlock.PrevHash != previousBlock.Hash {
			bclk.Errorf(
				"The chain is invalid, the %d block has an invalid previous hash",
				actualBlock.Index,
			)
			return fmt.Errorf(
				"the chain is invalid, the %d block has an invalid previous hash",
				actualBlock.Index,
			)
		}

		// Check if the actual hash is valid
		if actualBlock.Hash != actualBlock.CalculateHash() {
			bclk.Errorf(
				"The chain is invalid, the %d block has an invalid hash",
				actualBlock.Index,
			)
			return fmt.Errorf(
				"the chain is invalid, the %d block has an invalid hash",
				actualBlock.Index,
			)
		}
	}

	return nil
}

func (bc *Blockchain) CreateBlock(data string) Block {
	return NewBlock(
		uint64(len(bc.chain)),
		data,
		bc.chain[len(bc.chain)-1].Hash,
	)
}

func (bc *Blockchain) Mine(block *Block) *MiningResult {
	start := time.Now().UnixMilli()
	for {
		block.Hash = block.CalculateHash()
		if strings.HasPrefix(block.Hash, strings.Repeat("0", int(bc.difficulty))) {
			time := uint64(time.Now().UnixMilli() - start)
			bc.chain = append(bc.chain, *block)
			bclk.Infof("Block %d hash %s mined in %d ms", block.Index, block.Hash, time)
			return NewMiningResult(
				block,
				uint64(time),
			)
		}
		block.Nonce++
	}
}

func (bc *Blockchain) SetDifficulty(difficulty uint64) {
	bc.difficulty = difficulty
}
