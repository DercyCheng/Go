package main

import (
	"container/list"
	"fmt"
)

// 列表
func main() {
	// 双链表实现
	l := list.New()
	l.PushBack("2")
	l.PushFront("1")

	element := l.PushBack("4")
	l.InsertAfter("5", element)
	l.InsertBefore("3", element)
	l.Remove(element)

	/*
		遍历
			l.Front()表示初始赋值
			i!=nil每次进入循环都会进行 返回false就会退出循环
			继续执行i.next()
	*/
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
