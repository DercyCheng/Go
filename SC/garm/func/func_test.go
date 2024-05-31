package func_test

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// !函数作为参数
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValues()
	t.Log(a)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))

	fmt.Println("time out")
}

// !可变参数
// 可变参数是指函数的参数数量不固定,Go语言中的可变参数通过在参数名后加...来标识
func sum(x int, ops ...int) int {
	//ops是个切片,遍历所有数然后相加
	ret := 0 + x
	for _, op := range ops {
		ret += op
	}
	return ret
}

// defer
func hello() {
	fmt.Println("hello world")
}

func caclation() {
	//匿名函数保存在变量中
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20)
	//自执行函数
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}
func TestVarparam(t *testing.T) {
	defer hello()
	fmt.Println(sum(100, 1, 2, 3, 4, 5))
	caclation()
}

// !闭包=函数+引用环境
func adder(x int) func(int) int { //主要返回值
	return func(y int) int {
		x += y
		return x
	}
}

// 生成后缀
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func TestClosure(t *testing.T) {
	f := adder(10)
	fmt.Println(f(10)) //20
	jpg := makeSuffixFunc(".jpg")
	fmt.Println(jpg("test"))
}

// !多返回值 并命名
func calc(x int) (add func(int) int, sub func(int) int) {
	add = func(i int) int {
		x += i
		return x
	}

	sub = func(i int) int {
		x -= i
		return x
	}
	return
}
func TestCalc(t *testing.T) {
	f1, f2 := calc(10)
	fmt.Println(f1(1), f2(2)) //11 9
	fmt.Println(f1(3), f2(4)) //12 8
	fmt.Println(f1(5), f2(6)) //13 7
}

func Interviewcalc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func TestInterviewcalc(t *testing.T) {
	x := 1
	y := 2
	defer Interviewcalc("AA", x, Interviewcalc("A", x, y))
	x = 10
	defer Interviewcalc("BB", x, Interviewcalc("B", x, y))
	y = 20
}

// !递归
// 阶乘
func factorial(i int) int {
	if i < 1 {
		return 1
	}
	return i * factorial(i-1)
}

func TestFactorial(t *testing.T) {
	fmt.Println(factorial(10))
}

// fibonaci数列
func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}
func TestFibonaci(t *testing.T) {
	fmt.Println(fibonaci(3))
}
