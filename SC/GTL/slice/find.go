package slice

func Find[T any](ts []T, match matchFunc[T]) (T, bool) {
	for _, val := range ts {
		if match(val) {
			return val, true
		}
	}
	var t T
	return t, false
}

func FindAll[T any](ts []T, match matchFunc[T]) []T {
	res := make([]T, 0, len(ts)>>3+1)
	for _, val := range ts {
		if match(val) {
			res = append(res, val)
		}
	}
	return res
}
