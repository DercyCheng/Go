package slices

// Add 向切片中添加一个元素，并返回新的切片
// 参数:
//
//	slice []T: 需要添加元素的切片
//	element T: 要添加的元素
//
// 返回值:
//
//	[]T: 添加元素后的新切片
func Add[T any](slice []T, element T) []T {
	return append(slice, element)
}
