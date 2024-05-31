package channel

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	// 直接创建 返回值为nil 没有任何意义不建议使用
	var c1 chan int
	fmt.Println(c1) // <nil>

	// make声明
	c2 := make(chan int) // 类型可以是 bool []int ...
	fmt.Println(c2)

	c3 := <-c2
	fmt.Println(c3)
	close(c1) // all goroutines are sleep
	/*
		对一个关闭的通道再发送值就会导致panic。
		对一个关闭的通道进行接收会一直获取值直到通道为空。
		对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
		关闭一个已经关闭的通道会导致panic。
	*/
}

// 无缓冲死锁:启用一个goroutine去接收值
func TestUn_buffer(t *testing.T) {
	ch1 := make(chan int)
	// 无缓冲报错
	// ch1 <- 10
	// fmt.Println("close")

	recv := func(c chan int) {
		ret := <-c
		fmt.Println("成功接收", ret)
	}
	go recv(ch1) // 启用goroutine从通道接收值
	ch1 <- 10
	fmt.Println("done")
}

func TestBuffer(t *testing.T) {
	// 缓存空间为1
	ch := make(chan int, 1)
	ch <- 105
	fmt.Println(len(ch), cap(ch))
	fmt.Println("发送成功")
}

// 通道循环取值
func TestRange(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~5的数发送到ch1中
	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			} // without break before ch1 will always 0
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

// 双向通道
// 没有指定方向 那么Channel就是双向的 既可以接收数据也可以发送数据

// 单项通道改写TestRange
func TestSingle(t *testing.T) {
	/*
	   chan<- int是一个只写单向通道 可以发送但是不能接收
	   <-chan int是一个只读单向通道 可以接收但是不能发送
	*/
	counter := func(out chan<- int) {
		for i := 0; i < 5; i++ {
			out <- i
		}
		close(out)

	}
	squarer := func(out chan<- int, in <-chan int) {
		for i := range in {
			out <- i * i
		}
		close(out)
	}
	printer := func(in <-chan int) {
		for i := range in {
			fmt.Println(i)
		}
	}
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

// channel阻塞
func TestBlock(t *testing.T) {
	// rand.Seed(time.Now().UnixNano())
	ch1 := make(chan int)
	done := make(chan bool)
	go func() {
		fmt.Println("son goroutine start")
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		data := <-ch1
		fmt.Println("data", data)
		done <- true
		fmt.Println("son goroutine done")
	}()
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)
	ch1 <- 100
	<-done
	fmt.Println("done")
}

// 测试阻塞:打印一个数字的个位数的平方和
func TestBlockPrint(t *testing.T) {
	calcSquares := func(num int, squares chan int) {
		sum := 0
		for num != 0 {
			digit := num % 10
			sum += digit * digit
			num /= 10
		}
		squares <- sum
	}
	calcCubes := func(num int, cubes chan int) {
		sum := 0
		for num != 0 {
			digit := num % 10
			sum += digit * digit * digit
			num /= 10
		}
		cubes <- sum
	}
	num := 123
	squares := make(chan int)
	cubes := make(chan int)
	go calcSquares(num, squares)
	go calcCubes(num, cubes)
	square, cube := <-squares, <-cubes
	fmt.Println(square, cube)
}

// 死锁
func TestDeadBlock(t *testing.T) {
	ch1 := make(chan int)
	ch1 <- 5
	fmt.Println(ch1)
}

func TestRoutine(t *testing.T) {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
func TestSelect(t *testing.T) {
	/*
		分支语句：if，switch，select
		select 语句类似于 switch 语句，
			但是select会随机执行一个可运行的case。
			如果没有case可运行，它将阻塞，直到有case可运行。
	*/
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 200
	}()
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- 100
	}()
	// 随机执行
	select {
	case num1 := <-ch1:
		fmt.Println("ch1中取得数据", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2中取得数据", num2)
		} else {
			fmt.Println("error")
		}
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}
}
