package ops

import (
	"log"
	"testing"
)

func TestSum(t *testing.T) {
	cases := []struct {
		a, b     int
		expected int
	}{
		{0, 0, 0},
		{2, 2, 5},
		{0, 1, 1},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			s := Sum(tt.a, tt.b)
			if s != tt.expected {
				log.Fatalf("Sum(%d, %d) != %d", tt.a, tt.b, tt.expected)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	cases := []struct {
		a, b     int
		expected int
	}{
		{0, 0, 0},
		{0, 1, -1},
		{3, 2, 1},
	}

	for _, tt := range cases {
		t.Run("", func(t *testing.T) {
			s := Diff(tt.a, tt.b)
			if s != tt.expected {
				t.Fatalf("Diff(%d, %d) != %d", tt.a, tt.b, tt.expected)
			}
		})
	}
}
