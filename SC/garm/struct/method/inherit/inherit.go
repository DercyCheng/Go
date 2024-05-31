package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human  //匿名字段
	school string
}
type Employee struct {
	Human   //匿名字段
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Employee的method重写Human的method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name, e.company, e.phone) //Yes you can split into 2 lines here.
}
func main() {
	tom := Student{Human{"tom", 18, "123"}, "MIT"}
	tom.SayHi()
	lili := Employee{Human{"lili", 20, "456"}, "Google"}
	lili.SayHi()
}
