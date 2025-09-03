package exercises

func FindVariable(v any) (any, *any) {
	return v, &v
}
