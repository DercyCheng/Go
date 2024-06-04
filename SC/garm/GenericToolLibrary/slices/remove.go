package slices

// Remove 从切片中移除第一个匹配的元素，并返回新的切片
// 参数:
//
//	slice []T: 需要移除元素的切片
//	element T: 要移除的元素
//
// 返回值:
//
//	[]T: 移除元素后的新切片
func Remove[T comparable](slice []T, element T) []T {
	for i, v := range slice {
		if v == element {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
