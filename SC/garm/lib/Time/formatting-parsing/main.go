package main

import (
	"fmt"
	"time"
)

func main() {
	// 这里是一个基本的按照 RFC3339 进行格式化的例子，使用对应模式常量
	fmt.Println(time.Now().Format(time.RFC3339))

	// 时间解析使用同 `Format` 相同的形式值
	parse, err := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	if err != nil {
		return
	}
	fmt.Println(parse)
	// `Format` 和 `Parse` 使用基于例子的形式来决定日期格式
	// 一般你只要使用 `time` 包中提供的模式常量就行了，但是你也可以实现自定义模式
	// 模式必须使用时间 `Mon Jan 2 15:04:05 MST 2006` 来指定给定时间/字符串的格式化/解析方式
	// 时间一定要按照如下所示：2006为年，15 为小时，Monday 代表星期几，等等
	fmt.Println(time.Now().Format("2006-01-02T15:04:05.999999-07:00"))

	// 对于纯数字表示的时间，你也可以使用标准的格式化字符串来提出时间值的组成
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		time.Now().Year(), time.Now().Month(), time.Now().Day(),
		time.Now().Hour(), time.Now().Minute(), time.Now().Second())

	// `Parse` 函数在输入的时间格式不正确时会返回一个错误
	ansic := "Mon Jan _2 15:04:05 2006"
	_, err = time.Parse(ansic, "8:41PM")
	if err != nil {
		return
	}
}
