package main

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain(data string) *BlockChain {

	block := NewBlock(data, nil)

	blockChain := BlockChain{
		Blocks: []*Block{block},
	}

	return &blockChain
}

func (blockChain *BlockChain) AddBlock(data string) {
	lastBlock := blockChain.Blocks[len(blockChain.Blocks)-1]
	block := NewBlock(data, lastBlock.NowHash)
	blockChain.Blocks = append(blockChain.Blocks, block)
}
