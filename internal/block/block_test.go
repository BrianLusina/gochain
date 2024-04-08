package block

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase[T any] struct {
	hash     string
	data     T
	prevHash string
	expected *Block[T]
}

var testCases = []testCase[string]{
	{
		hash:     "hash",
		data:     "data",
		prevHash: "prev-hash",
		expected: &Block[string]{
			hash:     "hash",
			data:     "data",
			prevHash: "prev-hash",
		},
	},
}

func TestNewBlock(t *testing.T) {
	for _, tc := range testCases {
		params := BlockParams[string]{
			Hash:     tc.hash,
			Data:     tc.data,
			PrevHash: tc.prevHash,
		}
		actual := New(params)

		assert.Equal(t, tc.data, actual.Data())
		assert.Equal(t, tc.hash, actual.Hash())
		assert.Equal(t, tc.prevHash, actual.PrevHash())
	}
}
