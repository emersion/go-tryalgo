package arithm

// GCD computes the Greatest Common Divisor of a and b.
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}
