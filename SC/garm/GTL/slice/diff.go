package slice

func DiffSet[T comparable](ts, dst []T) []T {
	srcMap := toMap[T](ts)
	for _, val := range dst {
		delete(srcMap, val)
	}
	ret := make([]T, 0, len(srcMap))
	for key := range srcMap {
		ret = append(ret, key)
	}
	return ret
}

func DiffSetFunc[T any](ts, dst []T, equal equalFunc[T]) []T {
	ret := make([]T, 0, len(ts))
	for _, val := range ts {
		if !ContainsFunc[T](dst, func(ts T) bool {
			return equal(ts, val)
		}) {
			ret = append(ret, val)
		}
	}
	return DeduplicateFunc[T](ret, equal)
}
