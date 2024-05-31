package main

import "fmt"

type person struct {
	name string
	age  int
}

// 使用值接收者实现接口: 类型的值和类型的指针都能够保存到接口变量中
func (p person) move() { fmt.Println(p.name, "在动") }
func (p person) say()  { fmt.Println(p.name, "在说") }

// 接口嵌套
type sayer interface{ say() }
type mover interface{ move() }
type Human interface {
	mover
	sayer
}

func main() {
	p := person{ // p1是person类型的值
		name: "1",
		age:  18,
	}
	// p2 := &person{ // p2是person类型的指针
	// 	name: "2",
	// 	age:  18,
	// }

	// var m mover
	// m = p1 //?无法赋值的情况:p1是person类型的值没有实现mover接口 10定义的是指针
	// m = p2
	// m.move()
	// fmt.Println(m)

	var s sayer
	s = p
	s.say()
	fmt.Println(s)

	var a Human
	a = p
	a.move()
	a.say()
	fmt.Println(a)
}
