package arithm_test

import (
	"testing"

	"github.com/emersion/go-tryalgo/arithm"
)

func TestGCD(t *testing.T) {
	if gcd := arithm.GCD(20, 30); gcd != 10 {
		t.Errorf("GCD(20, 30) = %v, want 10", gcd)
	}
}
