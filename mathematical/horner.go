package mathematical

func horner(coef []int, x int) int {
	n := len(coef) - 1
	p := coef[n]
	for i := n - 1; i >= 0; i-- {
		p = x*p + coef[i]
	}
	return p
}
