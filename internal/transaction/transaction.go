package transaction

type (
	// Transaction structure representing a transfer between 2 participants in a network
	Transaction struct {
		// sender is the sender address
		sender string

		// receiver is the receiver's address
		receiver string

		// amount is the amount being transferred
		amount float64

		// coinbase is an indicator of whether the transaction is a coinbase transaction or not
		// A coinbase transaction is a special type of transaction that is primarily used to reward miners who successfully mine a new block.
		// It is the first transaction in each block and does not have a specific sender like regular transactions.
		// Instead, it is created by the miner who successfully solves the cryptographic puzzle associated with mining a block.
		// It rewards the miner by specifying the recipient’s address and an amount as a reward for their mining efforts.
		coinbase bool
	}

	// TransactionParams structure representing a transfer between 2 participants in a network
	TransactionParams struct {
		// sender is the sender address
		Sender string

		// receiver is the receiver's address
		Receiver string

		// amount is the amount being transferred
		Amount float64

		// coinbase is an indicator of whether the transaction is a coinbase transaction or not
		Coinbase bool
	}
)

// New creates a new transaction with the given parameters
func New(params TransactionParams) *Transaction {
	return &Transaction{
		sender:   params.Sender,
		receiver: params.Receiver,
		amount:   params.Amount,
		coinbase: params.Coinbase,
	}
}
