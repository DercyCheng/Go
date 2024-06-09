package slice

func toMap[T comparable](src []T) map[T]struct{} {
	var dataMap = make(map[T]struct{}, len(src))
	for _, v := range src {
		dataMap[v] = struct{}{}
	}
	return dataMap
}

func DeduplicateFunc[T any](data []T, equal equalFunc[T]) []T {
	newData := make([]T, 0, len(data))
	for k, v := range data {
		if !ContainsFunc[T](data[k+1:], func(ts T) bool {
			return equal(ts, v)
		}) {
			newData = append(newData, v)
		}
	}
	return newData
}

func Deduplicate[T comparable](data []T) []T {
	dataMap := toMap[](data)
	newData := make([]T, 0, len(dataMap))
	for key := range dataMap {
		newData = append(newData, key)
	}
	return newData
}
