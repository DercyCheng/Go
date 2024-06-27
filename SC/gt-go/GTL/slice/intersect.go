package slice

func IntersectSet[T comparable](ts, dst []T) []T {
	srcMap := toMap(ts)
	ret := make([]T, 0, len(ts))
	for _, val := range dst {
		if _, exist := srcMap[val]; exist {
			ret = append(ret, val)
		}
	}
	return Deduplicate[T](ret)
}

func IntersectSetFunc[T any](ts, dst []T, equalF equalFunc[T]) []T {
	ret := make([]T, 0, len(ts))
	for _, v := range dst {
		if ContainsFunc(ts, func(t T) bool {
			return equalF(t, v)
		}) {
			ret = append(ret, v)
		}
	}
	return DeduplicateFunc[T](ret, equalF)
}
