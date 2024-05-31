package main

import "fmt"

type dog struct {
	Name string
}
type cat struct {
	Name string
}
type person struct {
	Name string
}

func (d dog) say() {
	fmt.Println(d.Name, ":汪汪汪~")
}
func (c cat) say() {
	fmt.Println(c.Name, ":喵喵喵~")
}
func (p person) say() {
	fmt.Println(p.Name, ":rnm~")
}

type sayer interface {
	say()
}

func da(arg sayer) {
	arg.say()
}
func main() {
	var s sayer
	c := cat{Name: "猫"}
	s = c
	p := person{Name: "人"}
	s = p
	fmt.Println(s)
}
