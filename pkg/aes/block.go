package aes

import "github.com/JingBh/crypto-learn/pkg/utils"

func Cipher(input, key []byte) []byte {
	if len(input) != 16 {
		panic("invalid aes cipher input")
	}
	nk, nr := getNkAndNr(key)
	keys := keyExpansion(key, nk, nr)
	state := utils.Xor(input, keys[:16])
	for i := 1; i < nr; i++ {
		state = subBytes(state)
		state = shiftRows(state)
		state = mixColumns(state)
		state = utils.Xor(state, keys[16*i:16*i+16])
	}
	state = subBytes(state)
	state = shiftRows(state)
	state = utils.Xor(state, keys[16*nr:16*nr+16])
	return state
}

func InvCipher(input, key []byte) []byte {
	if len(input) != 16 {
		panic("invalid aes inv cipher input")
	}
	nk, nr := getNkAndNr(key)
	keys := keyExpansion(key, nk, nr)
	state := utils.Xor(input, keys[16*nr:16*nr+16])
	for i := nr - 1; i > 0; i-- {
		state = invShiftRows(state)
		state = invSubBytes(state)
		state = utils.Xor(state, keys[16*i:16*i+16])
		state = invMixColumns(state)
	}
	state = invShiftRows(state)
	state = invSubBytes(state)
	state = utils.Xor(state, keys[:16])
	return state
}

func keyExpansion(key []byte, nk, nr int) []byte {
	res := make([]byte, 16*(nr+1))

	for i := 0; i <= nk-1; i++ {
		copy(res[4*i:4*i+4], key[4*i:4*i+4])
	}
	for i := nk; i <= 4*nr+3; i++ {
		temp := make([]byte, 4)
		copy(temp, res[4*(i-1):4*(i-1)+4])
		if i%nk == 0 {
			temp = subBytes(rotWord(temp))
			temp[0] ^= rCon[i/nk-1]
		} else if nk > 6 && i%nk == 4 {
			temp = subBytes(temp)
		}
		copy(res[4*i:4*i+4], utils.Xor(res[4*(i-nk):4*(i-nk)+4], temp))
	}
	return res
}

func subBytes(state []byte) []byte {
	res := make([]byte, len(state))
	for i, v := range state {
		res[i] = sBox[v]
	}
	return res
}

func invSubBytes(state []byte) []byte {
	res := make([]byte, len(state))
	for i, v := range state {
		res[i] = invSBox[v]
	}
	return res
}

func shiftRows(state []byte) []byte {
	res := make([]byte, len(state))
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			res[c*4+r] = state[(c+r)%4*4+r]
		}
	}
	return res
}

func invShiftRows(state []byte) []byte {
	res := make([]byte, len(state))
	for r := 0; r < 4; r++ {
		for c := 0; c < 4; c++ {
			res[c*4+r] = state[(c-r+4)%4*4+r]
		}
	}
	return res
}

func mixColumns(state []byte) []byte {
	res := make([]byte, len(state))
	for c := 0; c < 4; c++ {
		copy(res[4*c:4*c+4], utils.GMulWord(state[4*c:4*c+4], []byte{0x02, 0x01, 0x01, 0x03}))
	}
	return res
}

func invMixColumns(state []byte) []byte {
	res := make([]byte, len(state))
	for c := 0; c < 4; c++ {
		copy(res[4*c:4*c+4], utils.GMulWord(state[4*c:4*c+4], []byte{0x0e, 0x09, 0x0d, 0x0b}))
	}
	return res
}

func rotWord(word []byte) []byte {
	return []byte{word[1], word[2], word[3], word[0]}
}

func getNkAndNr(key []byte) (int, int) {
	nk := len(key) / 4
	var nr int
	switch nk {
	case 4:
		nr = 10
	case 6:
		nr = 12
	case 8:
		nr = 14
	default:
		panic("invalid aes key length")
	}
	return nk, nr
}
