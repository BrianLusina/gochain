package block

import (
	"bytes"
	"crypto/md5"
	"fmt"
)

// Block represents a block in the block chain
type Block[T any] struct {
	// hash is the hash of the current block
	hash string

	// data is the data stored in the block
	data T

	// prevHash is the hash of the previous block
	prevHash string
}

type BlockParams[T any] struct {
	// Hash is the hash of the current block
	Hash string

	// Data is the data stored in the block
	Data T

	// PrevHash is the hash of the previous block
	PrevHash string
}

// New creates a new Block with a given struct of params
func New[T any](params BlockParams[T]) *Block[T] {
	return &Block[T]{
		hash:     params.Hash,
		data:     params.Data,
		prevHash: params.PrevHash,
	}
}

// Hash returns the hash of the current block
func (b *Block[T]) Hash() string {
	return b.hash
}

// Data returns the data of the current block
func (b *Block[T]) Data() any {
	return b.data
}

// PrevHash returns the previous has of the previous block in the chain
func (b *Block[T]) PrevHash() string {
	return b.prevHash
}

// ComputeHash computes the hash of this block from the previous has and the current data
func (b *Block[T]) ComputeHash() {
	dataString := fmt.Sprintf("%v", b.data)
	concatenatedData := bytes.Join([][]byte{[]byte(dataString), []byte(b.prevHash)}, []byte{})
	hash := md5.Sum(concatenatedData)
	b.hash = string(hash[:])
}
