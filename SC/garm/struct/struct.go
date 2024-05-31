package main

import (
	"fmt"
)

type Province struct {
	address string
}
type person struct {
	name, city string
	age        int8
	Province
}

type test struct {
	a int8
	b int8
	c int8
	d int8
}
type student struct {
	name string
	age  int8
}

/*
	func (接收者变量 接收者类型) 方法名(参数列表) 返回参数 {
		函数体
	}
*/
func newStudent(name string, age int8) *student {
	return &student{
		name: name,
		age:  age,
	}
}

func newPerson(name, city string, age int8) *person {
	return &person{
		name: name,
		city: city,
		age:  age,
	}
}

func (p student) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func (p *person) SetAge(newAge int8) {
	p.age = newAge
}
func (p person) _SetAge(newAge int8) {
	p.age = newAge
}

func main() {
	//基本实例化
	// var p1 person
	// p1.name = "name"
	// p1.city = "city"
	// p1.age = 18
	// fmt.Printf("p1=%v\n", p1)
	// fmt.Printf("p1=%#v\n", p1)

	//匿名结构体  函数内声明
	// var users struct {
	// 	Name string
	// 	Age  int
	// }
	// users.Name = "name"
	// users.Age = 18
	// fmt.Printf("%#v\n", users)

	//创建指针类型结构体
	// var p2 = new(person)
	// p2.name = "name"
	// p2.city = "city"
	// p2.age = 18
	// fmt.Println(reflect.TypeOf(p2))
	// fmt.Printf("p2=%#v\n", p2)

	//取结构体的地址实例化
	// p3 := &person{}
	// p3.name = "name"
	// p3.city = "city"
	// p3.age = 18
	// fmt.Println(reflect.TypeOf(p3))
	// fmt.Printf("p3=%#v\n", p3)

	//结构体初始化
	// var p4 person
	// fmt.Printf("p4=%#v\n", p4)

	//使用键值对初始化
	// p5 := person{
	// 	name: "name",
	// 	city: "city",
	// 	age:  18,
	// }
	// fmt.Printf("p5=%#v\n", p5)

	//对结构体指针进行键值对初始化
	// p6 := &person{
	// 	name: "name",
	// 	city: "city",
	// 	age:  18,
	// }
	// fmt.Printf("p6=%#v\n", p6)

	//初始化结构体的时候可以简写,也就是初始化的时候不写键,直接写值
	/**
	*必须初始化结构体的所有字段。
	*初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	*该方式不能和键值初始化方式混用。
	 */
	// p8 := &person{
	// 	"name",
	// 	"city",
	// 	18,
	// }
	// fmt.Printf("p8=%#v\n", p8)

	//结构体内存布局
	// n := &test{
	// 	1, 2, 3, 4,
	// }
	// fmt.Printf("n.a %p\n", &n.a)
	// fmt.Printf("n.b %p\n", &n.b)
	// fmt.Printf("n.c %p\n", &n.c)
	// fmt.Printf("n.d %p\n", &n.d)

	// interview
	// 会全部指向最后一个 stu

	m := make(map[string]*student)
	member := []student{
		{name: "first", age: 18},
		{name: "second", age: 23},
		{name: "third", age: 28},
	}

	// for _, stu := range member {
	// 	m[stu.name] = &stu
	// 	fmt.Println(stu)
	// }
	// FIXED
	for k, v := range member {
		m[v.name] = &member[k]
		fmt.Println(v)
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.age)
	}

	//!Go语言的结构体没有构造函数
	// p9 := newPerson("name", "city", 18)
	// fmt.Printf("%#v\n", p9)

	//函数不属于任何类型,方法属于特定的类型
	// p10 := newStudent("我", 25)
	// p10.Dream()

	// p11 := newPerson("name", "city", 20)
	// //!指针类型的接收者
	// p11.SetAge(30)
	// //!值类型接收者
	// p11._SetAge(30)
	// fmt.Println(p11.age)

	//嵌套
	// var p13 person
	// p13.city = "sz"
	// fmt.Println(p13)

}
