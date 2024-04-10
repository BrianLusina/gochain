package main

import (
	"fmt"
	"strconv"

	"github.com/brianlusina/gochain/internal/block"
	"github.com/brianlusina/gochain/internal/blockchain"
	"github.com/brianlusina/gochain/internal/factory"
	"github.com/brianlusina/gochain/internal/proof"
	"github.com/brianlusina/gochain/internal/transaction"
	"github.com/brianlusina/gochain/internal/wallet"
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

	aliceWallet, err := wallet.New("alice")
	if err != nil {
		fmt.Println("Error creating Alice's wallet:", err)
		return
	}
	fmt.Println("Alice's wallet created successfully")

	// Create a wallet for Bob.
	bobWallet, err := wallet.New("bob")
	if err != nil {
		fmt.Println("Error creating Bob's wallet:", err)
		return
	}
	fmt.Println("Bob's wallet created successfully")

	// Create a transaction from Alice to Bob.
	tx := transaction.New(transaction.TransactionParams{
		Sender:   aliceWallet.PublicKey().N.String(),
		Receiver: bobWallet.PublicKey().N.String(),
		Amount:   5.0,
	})
	fmt.Println("Alice to Bob Transaction created successfully")

	// Sign the transaction with Alice’s wallet.
	signature, err := aliceWallet.SignTransaction(tx)
	if err != nil {
		fmt.Println("Error signing the transaction:", err)
		return
	}

	// Verify the transaction using Alice’s wallet, public key, and the signature.
	err = wallet.VerifyTransaction(tx, aliceWallet.PublicKey(), signature)
	if err != nil {
		fmt.Println("Transaction verification failed:", err)
		return
	}

	fmt.Println("Transaction Verified Successfully")

	chain.AddBlock("Block 1", "Alice", 10.0, []*transaction.Transaction{tx})

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
