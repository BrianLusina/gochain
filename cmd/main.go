package main

import (
	"fmt"
	"strconv"

	"github.com/brianlusina/gochain/internal/block"
	"github.com/brianlusina/gochain/internal/blockchain"
	"github.com/brianlusina/gochain/internal/factory"
	"github.com/brianlusina/gochain/internal/proof"
	"github.com/brianlusina/gochain/internal/transaction"
)

func main() {
	coinbaseTransaction := transaction.New(transaction.TransactionParams{
		Sender:   "Coinbase",
		Receiver: "Genesis",
		Amount:   0.0,
		Coinbase: true,
	})

	genesisBlock := factory.CreateBlock("Genesis", "", []*transaction.Transaction{coinbaseTransaction})

	chain := blockchain.New(blockchain.BlockChainParams[string]{
		Blocks: []*block.Block[string]{genesisBlock},
	})

	chain.AddBlock("Block 1")
	chain.AddBlock("Block 2")
	chain.AddBlock("Block 3")

	for _, block := range chain.Blocks() {
		fmt.Printf("Previous hash: %x\n", block.PrevHash())
		fmt.Printf("Data in Block: %s\n", block.Data())
		fmt.Printf("Hash of block: %x\n", block.Hash())

		pow := proof.NewProofOfWork(*block)
		fmt.Printf("IsValidPoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}
