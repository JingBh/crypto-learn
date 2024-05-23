package des

import (
	"bytes"
	"testing"
)

var modeECBTests = []struct {
	key []byte
	in  []byte
	out []byte
}{{
	[]byte{0x6e, 0x5e, 0xe2, 0x47, 0xc4, 0xbf, 0xf6, 0x51},
	[]byte{0x11, 0xc9, 0x57, 0xff, 0x66, 0x89, 0x0e, 0xf0, 0xff, 0x3d, 0x25, 0x50, 0x12, 0xe3},
	[]byte{0x94, 0xc5, 0x35, 0xb2, 0xc5, 0x8b, 0x39, 0x72, 0xc4, 0x70, 0xe8, 0x30, 0xd0, 0x41, 0x8d, 0x6c},
}, {
	[]byte{0x6e, 0x5e, 0xe2, 0x47, 0xc4, 0xbf, 0xf6, 0x51},
	[]byte{0x11, 0xc9, 0x57, 0xff, 0x66, 0x89, 0x0e, 0xf0, 0xff, 0x3d, 0x25, 0x50, 0x12, 0xe3, 0x4a, 0xc5},
	[]byte{0x94, 0xc5, 0x35, 0xb2, 0xc5, 0x8b, 0x39, 0x72, 0x92, 0x19, 0xb4, 0xc5, 0x96, 0xb1, 0xec, 0x94},
}}

var modeCBCTests = []struct {
	key []byte
	iv  []byte
	in  []byte
	out []byte
}{{
	[]byte{0x6e, 0x5e, 0xe2, 0x47, 0xc4, 0xbf, 0xf6, 0x51},
	[]byte{0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef},
	[]byte{0x11, 0xc9, 0x57, 0xff, 0x66, 0x89, 0x0e, 0xf0, 0xff, 0x3d, 0x25, 0x50, 0x12, 0xe3, 0x4a, 0xc5},
	[]byte{0x23, 0x4b, 0x2c, 0x45, 0xfc, 0xb3, 0x9d, 0x0c, 0xe8, 0xa4, 0x1a, 0x0f, 0x53, 0xe8, 0x59, 0x36},
}}

func TestModeECB_Encipher(t *testing.T) {
	for i, tt := range modeECBTests {
		out := ECB(tt.key).Encipher(tt.in)
		if !bytes.Equal(out, tt.out) {
			t.Errorf("#%d: key: %x in: %x out: %x want: %x", i, tt.key, tt.in, out, tt.out)
		}
	}
}

func TestModeECB_Decipher(t *testing.T) {
	for i, tt := range modeECBTests {
		out := ECB(tt.key).Decipher(tt.out)
		if !bytes.Equal(out, tt.in) {
			t.Errorf("#%d: key: %x in: %x out: %x want: %x", i, tt.key, tt.out, out, tt.in)
		}
	}
}

func TestModeCBC_Encipher(t *testing.T) {
	for i, tt := range modeCBCTests {
		out := CBC(tt.key, tt.iv).Encipher(tt.in)
		if !bytes.Equal(out, tt.out) {
			t.Errorf("#%d: key: %x in: %x out: %x want: %x", i, tt.key, tt.in, out, tt.out)
		}
	}
}

func TestModeCBC_Decipher(t *testing.T) {
	for i, tt := range modeCBCTests {
		out := CBC(tt.key, tt.iv).Decipher(tt.out)
		if !bytes.Equal(out, tt.in) {
			t.Errorf("#%d: key: %x in: %x out: %x want: %x", i, tt.key, tt.out, out, tt.in)
		}
	}
}
