package slice

func Index[T comparable](ts []T, dst T) int {
	return IndexFunc(ts, func(src T) bool {
		return src == dst
	})
}

func IndexFunc[T any](ts []T, matchF matchFunc[T]) int {
	for k, v := range ts {
		if matchF(v) {
			return k
		}
	}
	return -1
}

func LastIndex[T comparable](ts []T, dst T) int {
	return LastIndexFunc[T](ts, func(src T) bool {
		return src == dst
	})
}

func LastIndexFunc[T any](ts []T, matchF matchFunc[T]) int {
	for i := len(ts) - 1; i >= 0; i-- {
		if matchF(ts[i]) {
			return i
		}
	}
	return -1
}

func IndexAll[T comparable](ts []T, dst T) []int {
	return IndexALlFunc(ts, func(src T) bool {
		return src == dst
	})
}

func IndexALlFunc[T any](ts []T, matchF matchFunc[T]) []int {
	indexes := make([]int, 0, len(ts))
	for k, v := range ts {
		if matchF(v) {
			indexes = append(indexes, k)
		}
	}
	return indexes
}
