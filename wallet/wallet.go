package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
)

type Wallet struct {
	pkey *ecdsa.PrivateKey
}

func New() (*Wallet, error) {
	pKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &Wallet{pKey}, nil
}

func (wallet *Wallet) PubKey() []byte {
	return append(wallet.pkey.PublicKey.X.Bytes(), wallet.pkey.PublicKey.Y.Bytes()...)
}

func (wallet *Wallet) Export(path string) error {
	err := os.MkdirAll(filepath.Dir(path), os.ModePerm)
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encodedPKey, err := x509.MarshalECPrivateKey(wallet.pkey)
	if err != nil {
		return err
	}

	pKeyPEM := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: encodedPKey,
	}
	return pem.Encode(file, pKeyPEM)
}

func Load(path string) (*Wallet, error) {
	pemData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	pKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return &Wallet{pKey}, nil
}
