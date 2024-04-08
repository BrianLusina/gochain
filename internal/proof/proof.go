package proof

import (
	"math/big"

	"github.com/brianlusina/gochain/internal/block"
)

// Difficulty represents the difficulty level for validating the proof of work algorithm
// The number of leading zeros required in the block’s hash determines the difficulty level.
// The more leading zeros required, the higher the difficulty.
const Difficulty int = 10

// ProofOfWork is the structure representing the PoW algorithm
type ProofOfWork[T any] struct {
	// Block is the block to be validated
	Block *block.Block[T]

	// Target is the difficulty level for mining
	Target *big.Int
}

// NewProofOfWork creates a new PoW structure
func NewProofOfWork[T any](block block.Block[T]) *ProofOfWork[T] {
	// The target for mining a new block will remain constant here; however, in a real-world blockchain implementation,
	// it is typically adjusted periodically to adapt to changes in the network’s computational power and maintain a
	// consistent block creation rate.
	targetVal := big.NewInt(1)

	targetVal.Lsh(targetVal, uint(256-Difficulty))
	return &ProofOfWork[T]{
		Block:  &block,
		Target: targetVal,
	}
}
