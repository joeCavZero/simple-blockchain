package blockchain

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Block struct {
	Index     uint64 `json:"index"`
	Timestamp int64  `json:"timestamp"`
	Data      string `json:"data"`
	PrevHash  string `json:"prevHash"`
	Hash      string `json:"hash"`
}

func NewBlock(index uint64, data string, prevHash string) Block {
	block := Block{
		Index:     index,
		Timestamp: time.Now().UnixMilli(),
		Data:      data,
		PrevHash:  prevHash,
		Hash:      "",
	}
	block.Hash = block.CalculateHash()
	return block
}

func NewGenesisBlock() Block {
	return NewBlock(0, "Genesis Block", "")
}

func (b *Block) CalculateHash() string {
	h := sha256.New()
	h.Write(
		[]byte(
			fmt.Sprintf(
				"%d%d%s%s",
				b.Index,
				b.Timestamp,
				b.Data,
				b.PrevHash,
			),
		),
	)
	return fmt.Sprintf("%x", h.Sum(nil))
}
