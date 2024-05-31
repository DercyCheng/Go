package main

import (
	"fmt"
	"sync"
)

var (
	x      int
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

//随机值 不加锁产生混乱
// func add() {
// 	for i := 0; i < 50000; i++ {
// 		x = x + 1
// 	}
// 	wg.Done()
// }

/*
	使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区
	其他的goroutine则在等待锁
	当互斥锁释放后,等待的goroutine才可以获取锁进入临界区
	多个goroutine同时等待一个锁时,唤醒的策略是随机的
*/
//10000
func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}
