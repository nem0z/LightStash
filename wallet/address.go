package wallet

import (
	"crypto/sha256"

	"github.com/ltcsuite/ltcd/ltcutil/base58"
	"github.com/ltcsuite/ltcd/ltcutil/bech32"
	"github.com/nem0z/LightStash/utils"
	"github.com/nem0z/LightStash/wallet/address"
	"golang.org/x/crypto/ripemd160"
)

func hash160(payload []byte) ([]byte, error) {
	hash256 := sha256.Sum256(payload)

	hasher := ripemd160.New()
	_, err := hasher.Write(hash256[:])
	if err != nil {
		return nil, err
	}

	return hasher.Sum(nil), nil
}

func (wallet *Wallet) P2PKH() (string, error) {
	hash160, err := hash160(wallet.PubKey())
	if err != nil {
		return "", err
	}

	pubKeyHash := utils.Prepend(hash160, address.P2PKHPrefix)
	checksum := utils.Checksum(pubKeyHash)

	return base58.Encode(append(pubKeyHash, checksum...)), nil
}

func (wallet *Wallet) Bench32() (string, error) {
	hash160, err := hash160(wallet.PubKey())
	if err != nil {
		return "", err
	}

	witnessProgram, err := bech32.ConvertBits(hash160, 8, 5, true)
	if err != nil {
		return "", err
	}

	witnessProgram = utils.Prepend(witnessProgram, address.WitnessVersion)
	return bech32.Encode(address.Bench32Prefix, witnessProgram)
}
