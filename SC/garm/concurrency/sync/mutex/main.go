package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// 对于更加复杂的情况，我们可以使用一个互斥锁 来在 Go 协程间安全的访问数据
// 全局变量
var ticket = 10 // 100张票
var lock sync.Mutex
var wg sync.WaitGroup

// 临界资源安全问题
func main() {
	/*
	   4个goroutine，模拟4个售票口，4个子程序操作同一个共享数据。
	*/
	wg.Add(4)
	go saleTickets("售票口1售出一张") // g1,100
	go saleTickets("售票口2售出一张") // g2,100
	go saleTickets("售票口3售出一张") //g3,100
	go saleTickets("售票口4售出一张") //g4,100

	// time.Sleep(5 * time.Second)
	wg.Wait()
	fmt.Println("done")
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done() //打开了4个goroutine 结束要关闭
	for {
		//ticket=1
		// `Lock()` 这个 `mutex` 来确保对 `state` 的独占访问读取选定的键的值
		//TODO `Unlock()` 这个 mutex 确保 ?
		lock.Lock()
		if ticket > 0 { //g1,g3,g2,g4
			//睡眠持续小于一个1000ms内
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			// g1 ,g3, g2,g4
			fmt.Println(name, "剩余:", ticket)
			ticket--
		} else {
			lock.Unlock() //不满足也要解锁
			fmt.Println(name, "售罄，没有票了")
			break
		}
		lock.Unlock()
		// 为了确保这个 Go 协程不会在调度中饿死
		// 在每次操作后明确的使用 `runtime.Gosched()` 进行释放
		// 这个释放一般是自动处理的
		// 每个通道操作后或者 `time.Sleep` 的阻塞调用后 相似
		// 但是在这个例子需要手动处理
		runtime.Gosched()
	}
}
