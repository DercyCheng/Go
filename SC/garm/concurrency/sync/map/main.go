package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.map 并发安全的map 不使用make
var (
	m  = sync.Map{}
	wg = sync.WaitGroup{}
)

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m.Store(key, n)
			value, _ := m.Load(key)
			fmt.Printf("k=:%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
