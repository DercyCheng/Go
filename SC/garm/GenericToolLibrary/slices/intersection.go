package slices

/**
 * Intersection 返回两个切片的交集
 * @param slice1 []T 第一个切片
 * @param slice2 []T 第二个切片
 * @return []T 两个切片的交集
 */
func Intersection[T comparable](slice1, slice2 []T) []T {
	var res []T
	for _, v := range slice1 {
		if Contains(slice2, v) {
			res = append(res, v)
		}
	}
	return res
}
