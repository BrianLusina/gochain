package block

import "github.com/brianlusina/gochain/internal/transaction"

type (
	// Block represents a block in the block chain
	Block[T any] struct {
		// hash is the hash of the current block
		hash string

		// data is the data stored in the block
		data T

		// prevHash is the hash of the previous block
		prevHash string

		// nonce is a number that is used once (nonce) is a value that miners update and include in the data they are hashing in
		// order to find a valid block hash. It is continuously changed for the same input data until a hash is generated that meets the target.
		nonce int

		// transactions represents a slice of the transactions held by this block
		transactions []*transaction.Transaction
	}

	// BlockParams is a structure with the parameters to create a block
	BlockParams[T any] struct {
		// Hash is the hash of the current block
		Hash string

		// Data is the data stored in the block
		Data T

		// PrevHash is the hash of the previous block
		PrevHash string

		// Nonce is a number used once
		Nonce int

		// Transactions is the slice of transactions for this block
		Transactions []*transaction.Transaction
	}
)

// New creates a new Block with a given struct of params. This is useful for reconstituting a given block from say storage with
// all the fields provided.
func New[T any](params BlockParams[T]) *Block[T] {
	return &Block[T]{
		hash:         params.Hash,
		data:         params.Data,
		prevHash:     params.PrevHash,
		nonce:        params.Nonce,
		transactions: params.Transactions,
	}
}

// Hash returns the hash of the current block
func (b *Block[T]) Hash() string {
	return b.hash
}

// WithHash updates a blocks hash & returns it with the updated hash
func (b Block[T]) WithHash(hash string) Block[T] {
	b.hash = hash
	return b
}

// Data returns the data of the current block
func (b *Block[T]) Data() any {
	return b.data
}

// PrevHash returns the previous has of the previous block in the chain
func (b *Block[T]) PrevHash() string {
	return b.prevHash
}

// Nonce returns the number only used once used for computing the hash of the block
func (b *Block[T]) Nonce() int {
	return b.nonce
}

// WithNonce returns a block with the updated nonce
func (b Block[T]) WithNonce(nonce int) Block[T] {
	b.nonce = nonce
	return b
}

// Transactions lists the transactions for this block
func (b Block[T]) Transactions() []*transaction.Transaction {
	return b.transactions
}
