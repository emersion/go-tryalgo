package arithm_test

import (
	"testing"

	"github.com/emersion/go-tryalgo/arithm"
)

func TestExp(t *testing.T) {
	if r := arithm.Exp(2, 10); r != 1024 {
		t.Errorf("Exp(2, 10) = %v, want 1024", r)
	}
	if r := arithm.Exp(3, 7); r != 2187 {
		t.Errorf("Exp(3, 7) = %v, want 2187", r)
	}
}

func TestExpMod(t *testing.T) {
	if r := arithm.ExpMod(2, 10, 17); r != 4 {
		t.Errorf("ExpMod(2, 10, 11) = %v, want 4", r)
	}
	if r := arithm.ExpMod(3, 7, 11); r != 9 {
		t.Errorf("ExpMod(3, 7, 11) = %v, want 9", r)
	}
}
