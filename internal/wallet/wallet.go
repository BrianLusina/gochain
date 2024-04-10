package wallet

import (
	"crypto/rsa"
	"fmt"
)

// Wallet structure representing a user's wallet with access to assets.
// RSA is sued for key generation, signing of transactions, and verification of transaction signatures.
type Wallet struct {
	// privateKey that in combination with the publicKey is used to perform decryption
	privateKey *rsa.PrivateKey

	// publicKey that is shared in the network used to encrypt data which is sent over to this wallet and used in combination
	// with the privateKey to perform decryption
	publicKey *rsa.PublicKey
}

// New creates a new wallet structure
func New() (*Wallet, error) {
	privateKey, publicKey, err := GenerateRSAKeys()
	if err != nil {
		return nil, fmt.Errorf("failed to create new wallet %s", err.Error())
	}

	return &Wallet{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}
