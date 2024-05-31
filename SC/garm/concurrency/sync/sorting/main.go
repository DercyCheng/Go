package main

import (
	"fmt"
	"sort"
)

// Go 的 sort 包实现了内置和用户自定义数据类型的排序功能
func main() {
	strs := []string{"b", "a", "d", "c"}
	ints := []int{5, 1, 2, 4, 3}

	sort.Strings(strs)
	sort.Ints(ints)

	fmt.Println(strs)
	fmt.Println(ints)

	//检查函数 true/false
	t := sort.IntsAreSorted(ints)
	fmt.Println(t)
}
