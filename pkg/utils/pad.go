package utils

import (
	"bytes"
)

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

func PKCS7Pad(input []byte, alignment int) []byte {
	padding := alignment - len(input)%alignment
	padBytes := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(input, padBytes...)
}

func PKCS7Unpad(input []byte) []byte {
	length := len(input)
	if length == 0 {
		return []byte{}
	}
	padding := int(input[length-1])
	return input[:length-padding]
}
