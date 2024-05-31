package main

import (
	"flag"
	"fmt"
)

//!指针获取命令行输入信息

func main() {
	mode := flag.String("mode", "", "process mode") //参数名,默认值,参数说明

	//解析命令行参数
	flag.Parse()
	fmt.Println(*mode) //fast
}

//命令行执行 go run session.go --mode=fast 再接-help就能看到参数说明
/**
 *解析--mode=fast
 *flag.Parse()写入结果
 *mode:=flag.String()提供数据
 *fmt.Println()输出
 */
