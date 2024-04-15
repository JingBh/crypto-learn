package des

import "crypto/rand"

func GenerateKey() []byte {
	key := make([]byte, 8)
	_, err := rand.Read(key)
	if err != nil {
		panic("failed to generate key")
	}

	// set parity bit for each byte
	for i := 0; i < 8; i++ {
		setParity(key[i])
	}

	return key
}

func setParity(b byte) byte {
	parity := 0
	for j := 1; j < 8; j++ {
		if b&(1<<uint(j)) != 0 {
			// test parity of each bit
			parity++
		}
	}
	return (b & ^byte(1)) | byte((parity+1)%2)
}
