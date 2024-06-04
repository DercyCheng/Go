package slices

// Contains 检查切片中是否包含指定的元素
// 参数:
//
//	slice []T: 需要检查的切片
//	element T: 要检查的元素
//
// 返回值:
//
//	bool: 如果切片中包含该元素则返回 true，否则返回 false
func Contains[T comparable](slice []T, element T) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}
