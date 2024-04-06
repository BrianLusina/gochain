package blockchain

import "github.com/brianlusina/gochain/internal/block"

// BlockChain represents the block chain
type BlockChain struct {
	blocks []*block.Block
}

// BlockChainParams represents a param struct to create a block chain
type BlockChainParams struct {
	// Blocks is a slice of blocks
	Blocks []*block.Block
}

// New creates a new block chain
func New(params BlockChainParams) *BlockChain {
	return &BlockChain{
		blocks: params.Blocks,
	}
}

// Blocks returns a slice of blocks in this chain
func (chain *BlockChain) Blocks() []*block.Block {
	return chain.blocks
}
