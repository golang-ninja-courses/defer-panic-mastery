package idna

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_punyEncode(t *testing.T) {
	cases := []struct {
		name     string
		in       string
		expected string
	}{
		{
			name:     "ascii only",
			in:       "helloworld",
			expected: "helloworld",
		},
		{
			name:     "ascii and unicode",
			in:       "привет world",
			expected: " world-hofyz3el1a",
		},
		{
			name:     "a lot of unicode",
			in:       "光榮歸於蘇共! 向共產主義前進!",
			expected: "! !-i88dp3l4gc57of0ggq2docjbmfd51bb24am7thg1b",
		},
		{
			name:     "unicode only",
			in:       "привет",
			expected: "b1agh1afp",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			out, err := punyEncode(tt.in)
			require.NoError(t, err)
			require.Equal(t, tt.expected, out)
		})
	}
}

func Test_encodeDigit(t *testing.T) {
	cases := []struct {
		name        string
		in          int32
		expectedSym byte
		wantPanic   bool
	}{
		{in: 0, expectedSym: 'a'},
		{in: 1, expectedSym: 'b'},
		{in: 2, expectedSym: 'c'},
		{in: 3, expectedSym: 'd'},
		{in: 4, expectedSym: 'e'},
		{in: 5, expectedSym: 'f'},
		{in: 6, expectedSym: 'g'},
		{in: 7, expectedSym: 'h'},
		{in: 8, expectedSym: 'i'},
		{in: 9, expectedSym: 'j'},
		{in: 10, expectedSym: 'k'},
		{in: 11, expectedSym: 'l'},
		{in: 12, expectedSym: 'm'},
		{in: 13, expectedSym: 'n'},
		{in: 14, expectedSym: 'o'},
		{in: 15, expectedSym: 'p'},
		{in: 16, expectedSym: 'q'},
		{in: 17, expectedSym: 'r'},
		{in: 18, expectedSym: 's'},
		{in: 19, expectedSym: 't'},
		{in: 20, expectedSym: 'u'},
		{in: 21, expectedSym: 'v'},
		{in: 22, expectedSym: 'w'},
		{in: 23, expectedSym: 'x'},
		{in: 24, expectedSym: 'y'},
		{in: 25, expectedSym: 'z'},

		{in: 26, expectedSym: '0'},
		{in: 27, expectedSym: '1'},
		{in: 28, expectedSym: '2'},
		{in: 29, expectedSym: '3'},
		{in: 30, expectedSym: '4'},
		{in: 31, expectedSym: '5'},
		{in: 32, expectedSym: '6'},
		{in: 33, expectedSym: '7'},
		{in: 34, expectedSym: '8'},
		{in: 35, expectedSym: '9'},

		{in: -1, name: "panic from -1", wantPanic: true},
		{in: 36, name: "panic from -36", wantPanic: true},
		{in: 0x7FFFFFFF, name: "panic from max int32", wantPanic: true},
	}

	for _, tt := range cases {
		n := tt.name
		if n == "" {
			n = string(tt.expectedSym)
		}

		t.Run(n, func(t *testing.T) {
			if tt.wantPanic {
				require.PanicsWithValue(t, "logic error in punycode encoding", func() {
					encodeDigit(tt.in)
				})
			} else {
				sym := encodeDigit(tt.in)
				require.Equal(t, string(tt.expectedSym), string(sym))
			}
		})
	}
}
