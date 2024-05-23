package utils

func Xor(input1, input2 []byte) []byte {
	output := make([]byte, len(input1))
	for i := range input1 {
		output[i] = input1[i] ^ input2[i]
	}
	return output
}

func GMul(a, b byte) byte {
	res := byte(0)
	for b != 0 {
		if b&1 == 1 {
			res ^= a
		}
		highBitSet := a & 0b10000000
		a <<= 1
		if highBitSet != 0 {
			a ^= 0b00011011
		}
		b >>= 1
	}
	return res
}

func GMulWord(a, b []byte) []byte {
	res := make([]byte, 4)
	res[0] = GMul(a[0], b[0]) ^ GMul(a[3], b[1]) ^ GMul(a[2], b[2]) ^ GMul(a[1], b[3])
	res[1] = GMul(a[1], b[0]) ^ GMul(a[0], b[1]) ^ GMul(a[3], b[2]) ^ GMul(a[2], b[3])
	res[2] = GMul(a[2], b[0]) ^ GMul(a[1], b[1]) ^ GMul(a[0], b[2]) ^ GMul(a[3], b[3])
	res[3] = GMul(a[3], b[0]) ^ GMul(a[2], b[1]) ^ GMul(a[1], b[2]) ^ GMul(a[0], b[3])
	return res
}
