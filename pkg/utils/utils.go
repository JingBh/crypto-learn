package utils

func GetBit(input []byte, pos byte) byte {
	return (input[pos/8] >> (7 - pos%8)) & 1
}
