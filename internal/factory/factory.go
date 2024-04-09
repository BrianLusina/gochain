package factory

import (
	"math/rand"
	"time"

	"github.com/brianlusina/gochain/internal/block"
	"github.com/brianlusina/gochain/internal/proof"
)

// CreateBlock creates a new block given a data item and a previous hash. This will have the
// hash of the block calculated. All new blocks will have a has value provided
func CreateBlock[T any](data T, prevHash string) *block.Block[T] {
	// Seed the random number generator
	rand.New(rand.NewSource(time.Now().UnixNano()))
	initialNonce := rand.Intn(10000)

	b := block.New[T](block.BlockParams[T]{
		Hash:     "",
		Data:     data,
		PrevHash: prevHash,
		Nonce:    initialNonce,
	})

	pow := proof.NewProofOfWork[T](*b)
	nonce, hash := pow.MineBlock()

	bloc := b.
		WithHash(string(hash)).
		WithNonce(nonce)

	return &bloc
}
