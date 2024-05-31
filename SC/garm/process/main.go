package main

import (
	"fmt"
)

type data [2]int

func main() {
	//!if语句
	// _if()

	//!if变体
	// ifVariants()

	//!switch continue 在go语言中的实现
	_fallthrough()

	//!可以判断类型
	// typeSwitch()
}

func _if() {
	a := 9
	if a <= 10 {
		fmt.Println("small than 10")
	}
}
func ifVariants() {
	if num := 10; num%2 == 0 { //num是一个局部变量
		fmt.Println(num)
	} else {
		fmt.Println("odd")
	}
}

// fallthrough不能用在switch的最后一个分支
// 会强制执行下一个case
func _fallthrough() {
	switch x := 5; x {
	default:
		fmt.Println(x)
	case 5:
		x += 5
		fmt.Println(x)
		fallthrough
	case 6:
		x += 10
		fmt.Println(x)
	}
}
func typeSwitch() {
	var x interface{}

	switch i := x.(type) {
	case nil:
		fmt.Printf(" x 的类型 :%T", i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
