package main

import (
	"fmt"
	"sync"
)

// 并发安全
func main() {
	var scene sync.Map //不能make创建
	scene.Store("0", 1)
	scene.Store(1, 2)
	scene.Store("2", "3")
	scene.Store(3, "4")

	fmt.Println(scene.Load("2"))
	scene.Delete(3)

	scene.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
}
