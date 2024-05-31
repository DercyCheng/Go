package main

import "fmt"

type person struct {
	Name, sex string
	age       int
}

type student struct {
	ID   int
	Name string
}

// !可以使用相同的名字
func (p person) display() {
	fmt.Printf("%s,%s,%d\n", p.Name, p.sex, p.age)
}
func (s student) display() {
	fmt.Printf("%s,%d\n", s.Name, s.ID)
}
func main() {
	p := person{
		Name: "cy",
		sex:  "boy",
		age:  20,
	}
	s := student{
		Name: "cy",
		ID:   1,
	}
	s.display()
	p.display()
}
