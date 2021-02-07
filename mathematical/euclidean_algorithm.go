package mathematical

// Euclidean finds the gcd of two numbers
func Euclidean(a, b int) int {
	for {
		m := a % b
		if m == 0 {
			break
		}
		a = b
		b = m
	}
	return b
}

// EuclideanRecursive finds the gcd of two numbers recursively
func EuclideanRecursive(a, b int) int {
	m := a % b
	if m == 0 {
		return b
	}
	return Euclidean(b, m)
}
