package slices

/**
 * Union 返回两个切片的并集
 * @param slice1 []T 第一个切片
 * @param slice2 []T 第二个切片
 * @return []T 两个切片的并集
 */
func Union[T comparable](slice1, slice2 []T) []T {
	res := make([]T, len(slice1))
	copy(res, slice1)
	for _, v := range slice2 {
		if !Contains(res, v) {
			res = append(res, v)
		}
	}
	return res
}
