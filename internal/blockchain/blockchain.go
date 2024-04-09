package blockchain

import (
	"github.com/brianlusina/gochain/internal/block"
	"github.com/brianlusina/gochain/internal/factory"
)

// BlockChain represents the block chain
type BlockChain[T any] struct {
	blocks []*block.Block[T]
}

// BlockChainParams represents a param struct to create a block chain
type BlockChainParams[T any] struct {
	// Blocks is a slice of blocks
	Blocks []*block.Block[T]
}

// New creates a new block chain
func New[T any](params BlockChainParams[T]) *BlockChain[T] {
	return &BlockChain[T]{
		blocks: params.Blocks,
	}
}

// Blocks returns a slice of blocks in this chain
func (chain *BlockChain[T]) Blocks() []*block.Block[T] {
	return chain.blocks
}

// AddBlock adds a new block to the block chain
func (chain *BlockChain[T]) AddBlock(data T) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	newBlock := factory.CreateBlock(data, prevBlock.PrevHash())
	chain.blocks = append(chain.blocks, newBlock)
}
