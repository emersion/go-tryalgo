package levenshtein_test

import (
	"testing"

	"github.com/emersion/go-tryalgo/levenshtein"
)

func TestLevenshtein(t *testing.T) {
	tests := []struct {
		w1, w2 string
		dist   int
	}{
		{"abc", "abc", 0},
		{"abc", "acc", 1},
		{"aac", "ac", 1},
		{"aaa", "baa", 1},
		{"aac", "ac", 1},
		{"abc", "aaa", 2},
	}

	for _, test := range tests {
		dist := levenshtein.Levenshtein([]byte(test.w1), []byte(test.w2))
		if dist != test.dist {
			t.Errorf("Levenshtein(%v, %v) = %v, want %v", test.w1, test.w2, dist, test.dist)
		}
	}
}
