package slice

func Reduce[T, U any](slice []T, initial U, f func(U, T) U) U {
	res := initial
	for _, v := range slice {
		res = f(res, v)
	}
	return res
}
