package proof

import (
	"bytes"
	"encoding/binary"
	"fmt"
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

// ComputeData takes a nonce integer and returns the computed for hashing the block
// A number used once (nonce) is a value that miners update and include in the data they are hashing in order to find a valid block hash.
// It is continuously changed for the same input data until a hash is generated that meets the target.
func (pow *ProofOfWork[T]) ComputeData(nonce int) []byte {
	blockData := fmt.Sprintf("%s", pow.Block.Data())

	data := bytes.Join(
		[][]byte{
			[]byte(pow.Block.PrevHash()),
			[]byte(blockData),
			make([]byte, 8),
			make([]byte, 8),
		},
		[]byte{},
	)

	binary.BigEndian.PutUint64(data[len(data)-16:], uint64(nonce))
	binary.BigEndian.PutUint64(data[len(data)-8:], uint64(Difficulty))

	return data
}
