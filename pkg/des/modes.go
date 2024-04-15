package des

type modeECB struct {
	key []byte
}

func ECB(key []byte) modeECB {
	if len(key) != 8 {
		panic("invalid des key")
	}
	return modeECB{
		key: key,
	}
}

func (m modeECB) Encipher(input []byte) []byte {
	padded := PadZero(input, 8)
	output := make([]byte, len(padded))
	for i := 0; i < len(padded)/8; i++ {
		res := EncipherBlock(padded[i*8:(i+1)*8], m.key)
		copy(output[i*8:(i+1)*8], res)
	}
	return output
}

func (m modeECB) Decipher(input []byte) []byte {
	padded := PadZero(input, 8)
	output := make([]byte, len(padded))
	for i := 0; i < len(padded)/8; i++ {
		res := DecipherBlock(padded[i*8:(i+1)*8], m.key)
		copy(output[i*8:(i+1)*8], res)
	}
	return UnpadZero(output)
}

type modeCBC struct {
	iv  []byte
	key []byte
}

func CBC(key, iv []byte) modeCBC {
	if len(key) != 8 {
		panic("invalid des key")
	}
	if len(iv) != 8 {
		panic("invalid cbc initialization vector")
	}
	return modeCBC{
		iv:  iv,
		key: key,
	}
}

func (m modeCBC) Encipher(input []byte) []byte {
	padded := PadZero(input, 8)
	output := make([]byte, len(padded))
	last := make([]byte, 8)
	copy(last, m.iv)
	for i := 0; i < len(padded)/8; i++ {
		res := EncipherBlock(xor(last, padded[i*8:(i+1)*8]), m.key)
		copy(last, res)
		copy(output[i*8:(i+1)*8], res)
	}
	return output
}

func (m modeCBC) Decipher(input []byte) []byte {
	padded := PadZero(input, 8)
	output := make([]byte, len(padded))
	last := make([]byte, 8)
	copy(last, m.iv)
	for i := 0; i < len(padded)/8; i++ {
		res := xor(last, DecipherBlock(padded[i*8:(i+1)*8], m.key))
		copy(last, padded[i*8:(i+1)*8])
		copy(output[i*8:(i+1)*8], res)
	}
	return UnpadZero(output)
}
