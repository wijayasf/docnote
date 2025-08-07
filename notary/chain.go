package notary

import "time"

type Blockchain struct {
	Blocks     []Block
	Difficulty int
}

func NewBlockchain(difficulty int) *Blockchain {
	genesis := Block{0, time.Now().Unix(), "Genesis Block", "", "", 0}
	genesis.Mine(difficulty)
	return &Blockchain{[]Block{genesis}, difficulty}
}

func (bc *Blockchain) AddBlock(data string) Block {
	prev := bc.Blocks[len(bc.Blocks)-1]
	block := Block{
		Index:     prev.Index + 1,
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prev.Hash,
	}
	block.Mine(bc.Difficulty)
	bc.Blocks = append(bc.Blocks, block)
	return block
}

func (bc *Blockchain) GetChain() []Block {
	chain := make([]Block, len(bc.Blocks))
	copy(chain, bc.Blocks)
	return chain
}
