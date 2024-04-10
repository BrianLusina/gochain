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

	chain.AddBlock("Block 1", "Alice", 10.0, []*transaction.Transaction{
		transaction.New(transaction.TransactionParams{Sender: "Alice", Receiver: "Bob", Amount: 1.5}),
		transaction.New(transaction.TransactionParams{Sender: "Alice", Receiver: "Charlie", Amount: 19.5}),
	})

	chain.AddBlock("Block 2", "Bob", 10.0, []*transaction.Transaction{
		transaction.New(transaction.TransactionParams{
			Sender: "Bob", Receiver: "Charlie", Amount: 2.3,
		}),
	})

	for _, block := range chain.Blocks() {
		fmt.Printf("Previous hash: %x\n", block.PrevHash())
		fmt.Printf("Data in Block: %s\n", block.Data())
		fmt.Printf("Hash of block: %x\n", block.Hash())

		pow := proof.NewProofOfWork(*block)
		fmt.Printf("IsValidPoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()

		fmt.Println("Transactions:")
		for _, tx := range block.Transactions() {
			fmt.Printf("Sender: %s\n", tx.Sender())
			fmt.Printf("Receiver: %s\n", tx.Receiver())
			fmt.Printf("Amount: %f\n", tx.Amount())
			fmt.Printf("Coinbase: %t\n", tx.Coinbase())
			fmt.Println()
		}
		fmt.Println()
	}
}
