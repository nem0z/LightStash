package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"fmt"
)

type Wallet struct {
	private *ecdsa.PrivateKey
	Public  []byte
}

func New() (*Wallet, error) {
	pKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	pubKey := append(pKey.PublicKey.X.Bytes(), pKey.PublicKey.Y.Bytes()...)

	return &Wallet{pKey, pubKey}, nil
}

func NewFromKey(pKey *ecdsa.PrivateKey) (*Wallet, error) {
	pubKey := append(pKey.PublicKey.X.Bytes(), pKey.PublicKey.Y.Bytes()...)
	return &Wallet{pKey, pubKey}, nil
}
