package slice

// Reduce 对切片中的每个元素应用函数，累计结果并返回
// 参数:
//
//	slice []T: 要处理的切片
//	initial U: 初始值
//	f func(U, T) U: 应用到每个元素的函数
//
// 返回值:
//
//	U: 累计结果
func Reduce[T any, U any](slice []T, initial U, f func(U, T) U) U {
	result := initial
	for _, v := range slice {
		result = f(result, v)
	}
	return result
}
