package aes

import "github.com/JingBh/crypto-learn/pkg/utils"

type modeECB struct {
	key []byte
}

func ECB(key []byte) modeECB {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		panic("invalid aes key")
	}
	return modeECB{
		key: key,
	}
}

func (m modeECB) Encipher(input []byte) []byte {
	padded := utils.PKCS7Pad(input, 16)
	output := make([]byte, len(padded))
	for i := 0; i < len(padded)/16; i++ {
		res := Cipher(padded[i*16:(i+1)*16], m.key)
		copy(output[i*16:(i+1)*16], res)
	}
	return output
}

func (m modeECB) Decipher(input []byte) []byte {
	output := make([]byte, len(input))
	for i := 0; i < len(input)/16; i++ {
		res := InvCipher(input[i*16:(i+1)*16], m.key)
		copy(output[i*16:(i+1)*16], res)
	}
	return utils.PKCS7Unpad(output)
}

type modeCBC struct {
	iv  []byte
	key []byte
}

func CBC(key, iv []byte) modeCBC {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		panic("invalid aes key")
	}
	if len(iv) != 16 {
		panic("invalid cbc initialization vector")
	}
	return modeCBC{
		iv:  iv,
		key: key,
	}
}

func (m modeCBC) Encipher(input []byte) []byte {
	padded := utils.PKCS7Pad(input, 16)
	output := make([]byte, len(padded))
	last := make([]byte, 16)
	copy(last, m.iv)
	for i := 0; i < len(padded)/16; i++ {
		res := Cipher(utils.Xor(last, padded[i*16:(i+1)*16]), m.key)
		copy(last, res)
		copy(output[i*16:(i+1)*16], res)
	}
	return output
}

func (m modeCBC) Decipher(input []byte) []byte {
	output := make([]byte, len(input))
	last := make([]byte, 16)
	copy(last, m.iv)
	for i := 0; i < len(input)/16; i++ {
		res := utils.Xor(last, InvCipher(input[i*16:(i+1)*16], m.key))
		copy(last, input[i*16:(i+1)*16])
		copy(output[i*16:(i+1)*16], res)
	}
	return utils.PKCS7Unpad(output)
}
