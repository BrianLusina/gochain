package blockchain

import (
	"testing"

	"github.com/brianlusina/gochain/internal/block"
	"github.com/brianlusina/gochain/internal/transaction"
	"github.com/stretchr/testify/assert"
)

type testCase[T any] struct {
	blocks   []*block.Block[T]
	expected *BlockChain[T]
}

var testCases = []testCase[string]{
	{
		blocks: []*block.Block[string]{
			block.New(block.BlockParams[string]{
				Hash:     "hash",
				Data:     "data",
				PrevHash: "",
			}),
			block.New(block.BlockParams[string]{
				Hash:     "hash",
				Data:     "data-1",
				PrevHash: "prevHash",
			}),
			block.New(block.BlockParams[string]{
				Hash:     "hash",
				Data:     "data-2",
				PrevHash: "prevHash-2",
			}),
		},
		expected: &BlockChain[string]{
			blocks: []*block.Block[string]{},
		},
	},
}

func TestNewBlock(t *testing.T) {
	for _, tc := range testCases {
		params := BlockChainParams[string]{
			Blocks: tc.blocks,
		}
		actual := New(params)

		assert.Equal(t, tc.blocks, actual.Blocks())

		// adding a new block
		newBlockData := "data-3"
		recipient := "recipient"
		amount := 10.0
		transactions := []*transaction.Transaction{}
		actual.AddBlock(newBlockData, recipient, amount, transactions)
		assert.Equal(t, len(tc.blocks)+1, len(actual.blocks))
	}
}
