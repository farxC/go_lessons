package exercises

func IsEvenAndPositive(n int) (bool, bool) {
	isEven := n%2 == 0
	isPositive := n > 0

	return isEven, isPositive
}
