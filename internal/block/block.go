package block

// Block represents a block in the block chain
type Block struct {
	// hash is the hash of the current block
	hash string

	// data is the data stored in the block
	data any

	// prevHash is the hash of the previous block
	prevHash string
}

type BlockParams struct {
	// Hash is the hash of the current block
	Hash string

	// Data is the data stored in the block
	Data any

	// PrevHash is the hash of the previous block
	PrevHash string
}

// New creates a new Block with a given struct of params
func New(params BlockParams) *Block {
	return &Block{
		hash:     params.Hash,
		data:     params.Data,
		prevHash: params.PrevHash,
	}
}

// Hash returns the hash of the current block
func (b *Block) Hash() string {
	return b.hash
}

// Data returns the data of the current block
func (b *Block) Data() any {
	return b.data
}

// PrevHash returns the previous has of the previous block in the chain
func (b *Block) PrevHash() string {
	return b.prevHash
}
