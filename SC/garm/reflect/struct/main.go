package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"`
	Score int    `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "1",
		Score: 90,
	}
	t := reflect.TypeOf(stu1)
	v := reflect.ValueOf(stu1)

	// fmt.Println(t.Name(), t.Kind())

	// 通过for循环遍历结构体的所有字段信息
	//for i := 0; i < t.NumField(); i++ { //返回结构体成员字段数量。
	//	field := t.Field(i)
	//	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	//}

	//通过字段名获取指定结构体字段信息
	//if scoreField, ok := t.FieldByName("score"); ok {
	//	fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	//}

	//返回该类型的方法集中方法的数目
	for i := 0; i < v.NumMethod(); i++ {
		fmt.Printf("method name : %s\n", t.Method(i).Name)
		fmt.Printf("method : %s\n", v.Method(i).Type())
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func (s student) Study() {
	fmt.Println("study")
}
func (s student) Sleep() {
	fmt.Println("sleep")
}
