package main

import "fmt"

// recover 会把错误吞掉 panic 会抛出错误不会继续执行
func TestRecover() {
	s := []string{"1", "2", "3", "4", "5"}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("defer finished")
	}()
	fmt.Println(s[5])
}
func main() {
	TestRecover()
}
