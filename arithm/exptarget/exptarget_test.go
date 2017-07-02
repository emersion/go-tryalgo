package exptarget_test

import (
	"testing"

	"github.com/emersion/go-tryalgo/arithm/exptarget"
)

func TestExpTarget(t *testing.T) {
	numbers := []int{1, 2, 3}
	results := map[int]bool{
		1: true,
		2: true,
		3: true,
		4: true, // 1 + 3
		5: true, // 2 + 3
		6: true, // 2 * 3
		7: true, // 2 * 3 + 1
		8: true, // (1 + 3) * 2
		9: true, // 3 * (1 + 2)
		10: false,
		11: false,
		12: false,
	}
	targets := exptarget.ExpTarget(numbers)
	for target, want := range results {
		if got := targets[target]; (got != nil) != want {
			t.Errorf("ExpTarget(%v, %v) = %v, want %v", numbers, target, got, want)
		}
	}
}
