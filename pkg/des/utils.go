package des

func rotateLeft28(input uint32, k int) uint32 {
	return ((input << (4 + k)) >> 4) | ((input << 4) >> (32 - k))
}
