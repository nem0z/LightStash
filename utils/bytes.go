package utils

import "crypto/sha256"

func Prepend(payload []byte, b byte) []byte {
	slice := make([]byte, len(payload)+1)
	slice[0] = b
	copy(slice[1:], payload)
	return slice
}

func Checksum(payload []byte) []byte {
	checksumLen := 4
	hash := sha256.Sum256(payload)
	tmp := sha256.Sum256(hash[:])
	return tmp[:checksumLen]
}

func Join(data ...byte) []byte {
	return data
}
