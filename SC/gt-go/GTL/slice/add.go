package slice

import "GTL/internal/slice"

// Add 是一个通用函数，用于在特定索引处向切片添加元素。
// 它使用 internal/slice 包中的 Add 函数。
// 该函数接受三个参数：
// - s：要添加元素的切片。
// - element：要添加到切片的元素。
// - index：应插入元素的位置。
// 该函数返回一个新的切片，其中添加了元素，如果在操作过程中出现任何错误，还会返回错误。
func Add[T any](s []T, element T, index int) ([]T, error) {
	res, err := slice.Add(s, element, index)
	return res, err
}
