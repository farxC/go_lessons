package exercises

// Eucliden Algorithm to find the greatest common divisor.
func mdc(a int, b int) int {
	for b != 0 { // For in this case acts like a while loop
		a, b = b, a%b
	}
	return a
}

func Mmc(n1 int, n2 int) int {
	if n1 == 0 || n2 == 0 {
		return 0
	}

	return (n1 / mdc(n1, n2)) * n2
}
