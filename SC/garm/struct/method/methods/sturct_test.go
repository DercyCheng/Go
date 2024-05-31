package struct_test

import (
	"fmt"
	"testing"
)

// !结构体比较
func TestCompare(t *testing.T) {
	type name struct {
		Fname, Lname string
	}
	name1 := name{
		"c", "y",
	}
	name2 := name{
		"c", "y",
	}
	if name1 == name2 {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

//!不可比较
// func TestUncompare(t *testing.T) {
// 	type image struct {
// 		data map[int]int
// 	}
// 	image1 := image{data: map[int]int{
// 		0: 1,
// 	}}
// 	image2 := image{data: map[int]int{
// 		1: 2,
// 	}}
// 	if image1 == image2 {
// 		fmt.Println("equal")
// 	}
// }

// !结构体作为函数的参数
func TestFuncwithStruct(t *testing.T) {

	type person struct {
		name struct {
			Fname, Lname string
		}
	}
	person1 := person{
		name: struct {
			Fname string
			Lname string
		}{
			"c", "y",
		},
	}

	PrintFS := func(p person) {
		fmt.Println(p.name.Fname)
		fmt.Println(p.name.Lname)
	}
	PrintFS(person1)
}
