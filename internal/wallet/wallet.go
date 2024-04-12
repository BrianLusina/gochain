package wallet

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/brianlusina/gochain/internal/transaction"
)

// Wallet structure representing a user's wallet with access to assets.
// RSA is sued for key generation, signing of transactions, and verification of transaction signatures.
type Wallet struct {
	// userID owner of this wallet
	userID string

	// privateKey that in combination with the publicKey is used to perform decryption
	privateKey *rsa.PrivateKey

	// publicKey that is shared in the network used to encrypt data which is sent over to this wallet and used in combination
	// with the privateKey to perform decryption
	publicKey *rsa.PublicKey
}

// New creates a new wallet structure
func New(userID string) (*Wallet, error) {
	privateKey, publicKey, err := GenerateRSAKeys()
	if err != nil {
		return nil, fmt.Errorf("failed to create new wallet %s", err.Error())
	}

	return &Wallet{
		userID:     userID,
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}

// UserID is the ID of the user who owns this wallet
func (w *Wallet) UserID() string {
	return w.userID
}

// PublicKey returns the publicKey associated with this wallet
func (w *Wallet) PublicKey() *rsa.PublicKey {
	return w.publicKey
}

// SignTransaction signs a given transaction and returns the signature of the signing or an error
// the signature is then used to validate a transaction
func (w *Wallet) SignTransaction(transaction *transaction.Transaction) (string, error) {
	// concatenate fields of transaction
	data := fmt.Sprintf("%s%s%f%t", transaction.Sender(), transaction.Receiver(), transaction.Amount(), transaction.Coinbase())

	// hash string;
	hashedData := sha256.Sum256([]byte(data))

	// sign the hash using wallet private key
	signature, err := rsa.SignPKCS1v15(rand.Reader, w.privateKey, crypto.SHA256, hashedData[:])
	if err != nil {
		return "", err
	}

	// encode signature as base64 string & returns the signature
	return base64.StdEncoding.EncodeToString(signature), nil
}

// VerifyTransaction verifies a given transaction with a public key and the signature returning an error if verification fails
func VerifyTransaction(transaction *transaction.Transaction, publicKey *rsa.PublicKey, signature string) error {
	dataString := fmt.Sprintf("%s%s%f%t", transaction.Sender(), transaction.Receiver(), transaction.Amount(), transaction.Coinbase())

	hashedData := sha256.Sum256([]byte(dataString))

	signatureBytes, err := base64.StdEncoding.DecodeString(signature)

	if err != nil {
		return err
	}

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashedData[:], signatureBytes)

	if err != nil {
		return errors.New("transaction Signature not valid")
	}
	return nil
}
