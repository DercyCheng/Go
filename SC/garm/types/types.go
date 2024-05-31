package main

import "fmt"

func main() {
	//!基本类型:数值类型,浮点类型,字符类型,布尔类型
	//@modify
	//本身的值没有改变,传递的是一个副本,并且返回一个新创建的字符串
	// name := `张三`
	// fmt.Println(modify(name))
	// fmt.Println(name)

	//!应用类型:切片,map,接口,函数类型,chan
	//@_modify
	//引用类型是一个指向底层数据的指针,操作的时候可以修改共享的底层数据的值,进而影响到所有引用到这个共享底层数据的变量
	ages := map[string]int{"one": 1}
	fmt.Println(ages)
	_modify(ages)
	fmt.Println(ages)
}

func modify(s string) string {
	s = s + s
	return s
}

func _modify(m map[string]int) {
	m["one"] = 10
}
