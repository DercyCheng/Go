package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
	//lock   sync.Mutex
	rwlock sync.RWMutex
)

/*
	读锁不能阻塞读锁
	读锁需要阻塞写锁，直到所有读锁都释放
	写锁需要阻塞读锁，直到所有写锁都释放
	写锁需要阻塞写锁
*/

func main() {
	wg.Add(3)
	go writeData(1)
	go readDate(2)
	go writeData(3)
	wg.Wait()
	fmt.Println("done")

}
func writeData(i int) {
	defer wg.Done()
	fmt.Println(i, "写 write start")
	rwlock.Lock()
	fmt.Println(i, "写 writing")
	time.Sleep(time.Second)
	rwlock.Unlock()
	fmt.Println(i, "写 write done")
}
func readDate(i int) {
	defer wg.Done()
	fmt.Println(i, "读 read start")
	rwlock.RLock()
	fmt.Println(i, "读 reading")
	time.Sleep(time.Second)
	rwlock.RUnlock()
	fmt.Println(i, "读 read done")
}
