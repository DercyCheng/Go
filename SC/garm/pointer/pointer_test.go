package printer_test

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	a := 10
	b := &a //取地址
	c := *b //取值 注意要使用指针
	fmt.Println(a, b, &b, c)
}

func TestPointerPass(t *testing.T) {
	modify1 := func(x int) {
		x = 100
	}
	modify2 := func(x *int) {
		*x = 100

	}
	a := 10
	modify1(a)
	fmt.Println(a) // 10
	modify2(&a)
	fmt.Println(a)
}

// !不常用
func TestNew(t *testing.T) {
	a := new(int) //得到的只是一个指针
	*a = 10
	fmt.Println(*a)
}

func adder(x int) func(int) int { //主要返回值
	return func(y int) int {
		x += y
		return x
	}
}
func TestMake(t *testing.T) {
	b := make(map[int]func(int) int)
	b[1] = adder(10)
	fmt.Println(b)
}

/**
 * map常用于slice,map,channel的分配
 * new分配内存但是没有初始化
 * 二者都是分配内存的
 */
