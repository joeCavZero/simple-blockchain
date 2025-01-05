package blockchain

type MiningResult struct {
	MinedBlock *Block
	MinedHash  string
	MineTime   uint64
}

func NewMiningResult(minedBlock *Block, minedHash string, mineTime uint64) *MiningResult {
	return &MiningResult{
		MinedBlock: minedBlock,
		MinedHash:  minedHash,
		MineTime:   mineTime,
	}
}
