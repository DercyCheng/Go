package main

import (
	"fmt"
	"sort"
	"testing"
)

type Bylength []string

// 全局函数 需要大写调用
func (b Bylength) Len() int {
	return len(b)
}

func (b Bylength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b Bylength) Less(i, j int) bool {
	return len(b[i]) < len(b[j])
}

// 我们现在可以通过将原始的 `fruits` 切片转型成 `ByLength` 来实现我们的自定排序了
// 然后对这个转型的切片使用 `sort.Sort` 方法
func TestSorting_function(t *testing.T) {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(Bylength(fruits))
	fmt.Println(fruits)
}
