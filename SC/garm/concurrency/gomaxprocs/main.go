package main

import (
	"fmt"
	"runtime"
	"time"
)

//workerpool可以用runtime.GOMAXPROCS代替
/*
	一个操作系统线程对应用户态多个goroutine。
	go程序可以同时使用多个操作系统线程。
	goroutine和OS线程是多对多的关系，即m:n
*/
func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("A", i)
	}
}
func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("B", i)
	}
}
func main() {
	var x int
	fmt.Scanf("%d", x)
	//线程
	runtime.GOMAXPROCS(x) //x个任务同时进行

	//随机执行
	go a()
	go b()
	time.Sleep(time.Second)
}
