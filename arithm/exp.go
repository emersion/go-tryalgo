package arithm

// Exp computes x^n.
func Exp(x, n int) int {
	r := 1
	for n > 0 {
		if n%2 == 1 {
			r *= x
		}
		x = x * x
		n /= 2
	}
	return r
}

// ExpMod computes x^n % q.
func ExpMod(x, n, q int) int {
	r := 1
	for n > 0 {
		if n%2 == 1 {
			r = (r * x) % q
		}
		x = (x * x) % q
		n /= 2
	}
	return r
}
