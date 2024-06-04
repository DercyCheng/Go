package maps

// Add 向 map 中添加一个键值对，并返回新的 map
// 参数:
//
//	m map[K]V: 需要添加键值对的 map
//	key K: 要添加的键
//	value V: 要添加的值
//
// 返回值:
//
//	map[K]V: 添加键值对后的新 map
func Add[K comparable, V any](m map[K]V, key K, value V) map[K]V {
	m[key] = value
	return m
}
