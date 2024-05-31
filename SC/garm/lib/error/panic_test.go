package panic_test

import (
	"fmt"
	"testing"
)

// !异常处理
// panic
func TestPanic(t *testing.T) {
	funcA := func() {
		fmt.Println("A")
	}

	funcB := func() {
		// panic("error")
		fmt.Println("B")
	}
	funcA()
	funcB()

	defer func() {
		if err := recover(); err != nil { //掩盖
			fmt.Println(err)
		}
	}()
	ch := make(chan int, 10)
	close(ch)
	ch <- 1

	defer func() {
		panic("defer panic")
	}()
}
