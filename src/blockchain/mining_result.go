package blockchain

type MiningResult struct {
	MinedBlock *Block
	MineTime   uint64
}

func NewMiningResult(minedBlock *Block, mineTime uint64) *MiningResult {
	return &MiningResult{
		MinedBlock: minedBlock,
		MineTime:   mineTime,
	}
}
