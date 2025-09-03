package exercises

func Divide(a int) func(b int) int {
	return func(b int) int {
		if b >= a {
			return b / a
		}
		return a / b
	}
}
