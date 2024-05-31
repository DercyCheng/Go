package main

import (
	"fmt"
	"reflect"
)

// Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空
func main() {
	x := 1
	fmt.Println(reflect.TypeOf(x))
	reflectType(x)

	type myint int
	var y myint
	y = 10
	reflectType(y)

	v := "hello"
	reflectValue(v)

	reflectSetvalue(&x)
	fmt.Println(x)

	//IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("Name:%v\nKind:%v\n", v.Name(), v.Kind())
}

// 通过反射获取值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int, value is %d\n", int(v.Int()))
	case reflect.Float32:
		// v.Float()从反射中获取浮点型的原始值，然后通过float32()强制类型转换
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	case reflect.String:
		// v.Float()从反射中获取浮点型的原始值，然后通过float64()强制类型转换
		fmt.Printf("type is string, value is %s\n", string(v.String()))
	}
}

// 通过反射设置变量的值
func reflectSetvalue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int {
		v.Elem().SetInt(100)
	}
}
