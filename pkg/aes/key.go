package aes

import "crypto/rand"

func GenerateKey(len int) []byte {
	if len != 128 && len != 192 && len != 256 {
		panic("invalid key length")
	}
	key := make([]byte, len/8)
	_, err := rand.Read(key)
	if err != nil {
		panic("failed to generate key")
	}
	return key
}
