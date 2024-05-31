package main

import "fmt"

func main() {
	//空接口不需要提前定义
	var x interface{}
	x = "hello"
	// fmt.Println(x)
	x = 123
	// fmt.Println(x)

	ret, ok := x.(string) //类型断言猜的类型不对时，ok=false,ret=string类型的零值
	if !ok {
		fmt.Println("not string")
	} else {
		fmt.Println("type is string", ret)
	}

	switch v := x.(type) {
	case string:
		fmt.Printf("是字符串类型, value:%v\n", v)
	case bool:
		fmt.Printf("是布尔类型, value:%v\n", v)
	case int:
		fmt.Printf("是int类型, value:%v\n", v)
	default:
		fmt.Printf("猜不到了, value:%v\n", v)
	}

	person := make(map[string]interface{}, 16)
	person["name"] = "name"
	person["hobby"] = []string{"篮球", "足球"}
	fmt.Println(person)
}
