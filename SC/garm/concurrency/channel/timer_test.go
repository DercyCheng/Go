package channel

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestNewTimer(t *testing.T) {
	/*
		func NewTimer(d Duration) *Timer
			创建一个计时器：d时间以后触发，go触发计时器的方法比较特别，就是在计时器的channel中发送值
	*/
	//新建一个计时器：timer
	timer := time.NewTimer(3 * time.Second)
	fmt.Println(reflect.TypeOf(timer), "\n", time.Now())

	ch1 := timer.C
	fmt.Println(<-ch1)
}

// 停止前打点3次
func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for i := range ticker.C {
			fmt.Println("Tick at", i)
		}
	}()
	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("ticker stopped")
}

// 等待持续时间过去后发送当前时间
func TestAfter(t *testing.T) {
	fmt.Println(time.Now())
	ch := time.After(3 * time.Second)
	timer := <-ch
	fmt.Println(timer)
}

// 等待3秒结束
func TestTimerStop(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	go func() {
		<-timer.C
		fmt.Println("timer done")
	}()
	time.Sleep(3 * time.Second)
	stop := timer.Stop()
	if stop {
		fmt.Println("timer stop")
	}
}
