package main

import (
	"fmt"
	"strings"
)

//通过使用组合函数来代替泛型
//TODO: 了解泛型
//运行情况: index()判断是否存在

func Index(ss []string, i string) int {
	for k, v := range ss {
		if v == i {
			return k
		}
	}
	return -1
}

func Include(ss []string, i string) bool {
	return Index(ss, i) >= 0
}

func Any(ss []string, f func(string) bool) bool {
	for _, v := range ss {
		if f(v) {
			return true
		}
	}
	return false
}

func All(ss []string, f func(string) bool) bool {
	for _, v := range ss {
		if f(v) {
			return false
		}
	}
	return true
}

func Filter(ss []string, f func(string) bool) []string {
	ssf := make([]string, 0)
	for _, v := range ss {
		if f(v) {
			ssf = append(ssf, v)
		}
	}
	return ssf
}

func Map(ss []string, f func(string) string) []string {
	ssm := make([]string, len(ss))
	for i, v := range ss {
		ssm[i] = f(v)
	}
	return ssm
}

func main() {

	// 这里试试这些组合函数。
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	// 上面的例子都是用的匿名函数，但是你也可以使用类型正确命名函数
	fmt.Println(Map(strs, strings.ToUpper))

}
