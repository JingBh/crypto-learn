package des

import "testing"

func TestSetParity(t *testing.T) {
	testCases := []struct {
		name     string
		input    byte
		expected byte
	}{
		{
			name:     "even parity",
			input:    0b10111000,
			expected: 0b10111001,
		},
		{
			name:     "odd parity",
			input:    0b10101011,
			expected: 0b10101011,
		},
		{
			name:     "all zero",
			input:    0b00000000,
			expected: 0b00000001,
		},
		{
			name:     "all one",
			input:    0b11111111,
			expected: 0b11111110,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := setParity(tc.input)
			if result != tc.expected {
				t.Errorf("Expected: %08b, Got: %08b", tc.expected, result)
			}
		})
	}
}
