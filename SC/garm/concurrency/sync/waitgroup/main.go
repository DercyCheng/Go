package main

import (
	"fmt"
	"sync"
)

// 在代码中生硬的使用time.Sleep肯定是不合适的可以使用sync.WaitGroup.Wait()
// Go语言中可以使用sync.WaitGroup来实现并发任务的同步
// 同步等待组
var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("a", i)
	}
}
func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println("b", i)
	}
}

func main() {
	wg.Add(2)
	go a()
	go b()
	fmt.Println("进入阻塞")
	wg.Wait()
	fmt.Println("解除阻塞")
}
