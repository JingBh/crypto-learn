package des

func PadZero(input []byte, alignment int) []byte {
	var paddedLen int
	if len(input)%alignment == 0 {
		paddedLen = len(input)
	} else {
		paddedLen = len(input) + alignment - len(input)%alignment
	}
	output := make([]byte, paddedLen)
	copy(output, input)
	return output
}

func UnpadZero(input []byte) []byte {
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != 0 {
			return input[:i+1]
		}
	}
	return nil
}

func getBit(input []byte, pos byte) byte {
	return (input[pos/8] >> (7 - pos%8)) & 1
}

func xor(input1, input2 []byte) []byte {
	output := make([]byte, len(input1))
	for i := range input1 {
		output[i] = input1[i] ^ input2[i]
	}
	return output
}

func rotateLeft28(input uint32, k int) uint32 {
	return ((input << (4 + k)) >> 4) | ((input << 4) >> (32 - k))
}
