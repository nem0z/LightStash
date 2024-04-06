package wallet

import (
	"crypto/sha256"

	"github.com/ltcsuite/ltcd/ltcutil/base58"
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

func (wallet *Wallet) P2PKH(mainet bool) (string, error) {
	hash160, err := hash160(wallet.Public)
	if err != nil {
		return "", err
	}

	checksum := utils.Checksum(hash160)

	prefix := address.LitecoinP2PKHTestnetPrefix
	if mainet {
		prefix = address.LitecoinP2PKHMainnetPrefix
	}

	data := []byte{}
	data = append(data, prefix)
	data = append(data, hash160...)
	data = append(data, checksum...)

	return base58.Encode(data), nil
}
