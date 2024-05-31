package main

//链式处理器
import (
	"fmt"
	"strings"
)

// 移除指定前缀
func removePrefix(str string) string { return strings.TrimPrefix(str, "go") }
func stringproccess(list []string, chain []func(string) string) {
	for index, str := range list {
		//将当前字符串保存到result中
		result := str
		//遍历每一个处理链
		for _, proc := range chain {
			//输入一个字符串进行处理,返回数据作为下一个处理链的输入
			result = proc(result)
		}
		list[index] = result
	}
}
func main() {
	list := []string{
		"go scanner",
		"go parser",
		"go compiler",
		"go printer",
		"go formater",
	}
	chain := []func(string) string{
		removePrefix,
		strings.TrimSpace, //移除空格
		strings.ToUpper,   //首字母大写
	}
	stringproccess(list, chain)
	for _, str := range list {
		fmt.Println(str)
	}
}
