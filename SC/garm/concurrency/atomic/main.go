package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// atomic包提供了底层的原子级内存操作
// 对于同步算法的实现很有用  这些函数必须谨慎地保证正确使用
// 除了某些特殊的底层应用   使用通道或者sync包的函数/类型实现同步更好

// Inc 比较下互斥锁和原子操作的性能
func (c CommonCounter) Inc() {
	c.counter++
}

func (c CommonCounter) Load() int64 {
	return c.counter
}

type Counter interface {
	Inc()
	Load() int64
}

// CommonCounter 普通版
type CommonCounter struct {
	counter int64
}

// MutexCounter 互斥锁版
type MutexCounter struct {
	counter int64
	lock    sync.Mutex
}

func (m *MutexCounter) Inc() {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.counter++
}

func (m *MutexCounter) Load() int64 {
	m.lock.Lock()
	defer m.lock.Unlock()
	return m.counter
}

// AtomicCounter 原子操作版
type AtomicCounter struct {
	counter int64
}

func (a *AtomicCounter) Inc() {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicCounter) Load() int64 {
	return atomic.LoadInt64(&a.counter)
}

func test(c Counter) {
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc()
			wg.Done()
		}()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(c.Load(), end.Sub(start))
}

func main() {
	c1 := CommonCounter{} // 非并发安全
	test(c1)
	c2 := MutexCounter{} // 使用互斥锁实现并发安全
	test(&c2)
	c3 := AtomicCounter{} // 原子操作并发安全且比互斥锁效率更高
	test(&c3)

	// 比较,交换
	var x int64 = 100
	ok := atomic.CompareAndSwapInt64(&x, 100, 200)
	fmt.Println(ok, x)
}
