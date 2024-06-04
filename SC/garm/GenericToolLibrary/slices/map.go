package slices

/**
 * Map 对切片中的每个元素应用一个函数，并返回一个包含结果的新切片
 * @param slice []T 输入切片
 * @param f func(T) U 应用于每个元素的函数
 * @return []U 应用函数后的新切片
 */
func Map[T, U any](slice []T, f func(T) U) []U {
	res := make([]U, len(slice))
	for i, v := range slice {
		res[i] = f(v)
	}
	return res
}
