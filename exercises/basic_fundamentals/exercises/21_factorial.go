package exercises

// Using tail recursion
func factorialHelper(n int, acc int) int {
	if n == 0 {
		return acc
	}
	if n == 1 {
		return acc
	}
	return factorialHelper(n-1, n*acc)
}

func Factorial(n int) int {
	return factorialHelper(n, 1)
}
