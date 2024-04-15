package des

import "encoding/binary"

// function to encipher a 64-bit block
func EncipherBlock(input, key []byte) []byte {
	if len(input) != 8 {
		panic("invalid des encipher input")
	}
	if len(key) != 8 {
		panic("invalid des key")
	}
	permuted := permute(input, initialPermutation)
	keys := scheduleKeys(key)
	l, r := permuted[:4], permuted[4:]
	for i := 0; i < 16; i++ {
		l, r = r, xor(l, process(r, keys[i]))
	}
	return permute(append(r, l...), initialPermutationInverse)
}

// function to decipher a 64-bit block
func DecipherBlock(input, key []byte) []byte {
	if len(input) != 8 {
		panic("invalid des decipher input")
	}
	if len(key) != 8 {
		panic("invalid des key")
	}
	permuted := permute(input, initialPermutation)
	keys := scheduleKeys(key)
	r, l := permuted[:4], permuted[4:]
	for i := 15; i >= 0; i-- {
		l, r = xor(r, process(l, keys[i])), l
	}
	return permute(append(l, r...), initialPermutationInverse)
}

func scheduleKeys(key []byte) [][]byte {
	keys := make([][]byte, 16)
	permuted := permute(key, permutedChoice1)
	value := binary.BigEndian.Uint64(append([]byte{0}, permuted...))
	c, d := uint32(value>>28), uint32(value<<4)>>4
	for i := 0; i < 16; i++ {
		if i == 0 || i == 1 || i == 8 || i == 15 {
			c, d = rotateLeft28(c, 1), rotateLeft28(d, 1)
		} else {
			c, d = rotateLeft28(c, 2), rotateLeft28(d, 2)
		}
		joined := make([]byte, 8)
		binary.BigEndian.PutUint64(joined, uint64(c)<<28|uint64(d))
		keys[i] = permute(joined[1:], permutedChoice2)
	}
	return keys
}

func permute(input, table []byte) []byte {
	output := make([]byte, len(table)/8)
	for i, pos := range table {
		bit := getBit(input, pos-1)
		output[i/8] |= bit << (7 - i%8)
	}
	return output
}

// implementation of the cipher function
func process(input, key []byte) []byte {
	output := make([]byte, 4)
	permuted := permute(input, eSelection)
	added := xor(permuted, key)
	for i := 0; i < 8; i++ {
		row := (getBit(added, byte(i*6)) << 1) |
			getBit(added, byte(i*6+5))
		col := (getBit(added, byte(i*6+1)) << 3) |
			(getBit(added, byte(i*6+2)) << 2) |
			(getBit(added, byte(i*6+3)) << 1) |
			getBit(added, byte(i*6+4))
		output[i/2] = (output[i/2] << 4) | sSelection[i][row*16+col]
	}
	return permute(output, permutation)
}
