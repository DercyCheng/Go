package main

import (
	"fmt"
	"os"
)

// 运行 :./command-line-arguments a b c d
func main() {

	// `os.Args` 提供原始命令行参数访问功能
	// 切片中的第一个参数是该程序的路径
	// `os.Args[1:]` 保存程序的所有参数
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 使用标准的索引位置方式取得单个参数的值
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
