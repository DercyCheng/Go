package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	//!遍历排序

	// dict := map[string]int{"foo": 1, "_foo": 2}
	// var names []string
	// for name := range dict {
	// 	names = append(names, name)
	// }
	// sort.Strings(names) //排序
	// for _, key := range names {
	// 	fmt.Println(key, dict[key])
	// }

	//?按照指定顺序遍历
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	scoreMap := make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	keys := make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

	//!声明和初始化
	// Init()
	//!遍历
	// Each()
	//!undefined 的值是0 不能通过nil来返回 需要注意
	// exist()

	//!map value为函数
	// mapwithfuncval()

	//!自定义Set方法
	// mapforSet()

	//!map类型的切片
	// mapSlice()

	//!值为切片类型的map : map里面有切片(动态数组)
	slicemap()
}

func modify(dict map[string]int) {
	dict["foo"] = 10
}

func Init() {
	dict := make(map[string]int)

	//添加一个key:foo,value:1的键值对数据
	dict["foo"] = 1
	dict["_foo"] = 2

	//第一个返回值是键的值,第二个返回值标记这个键是否存在
	age, exits := dict["foo"]
	fmt.Printf("value:%d\n是否存在:%t\n", age, exits)

	//第一个参数是要操作的Map,第二个是要删除的Map的键
	delete(dict, "_foo")
}
func Each() {

	dict := make(map[int]int)
	for key, value := range dict {
		fmt.Println(`遍历:`, key, value)
	}
}
func exist() {

	array := map[string]int{"first": 0, "second": 0, "third": 0}
	if v, ok := array["one"]; ok {
		fmt.Printf("value is %d", v)
	} else {

		fmt.Println("Not Existing")
	}
}

func mapwithfuncval() {
	m := map[int]func(op int) int{} //?
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	fmt.Println(m[1](2), m[2](2), m[3](2))
}

func mapforSet() {
	Set := map[int]bool{}
	Set[1] = true
	n := 1
	if Set[n] {
		fmt.Println("existing")
	} else {
		fmt.Println("Not existing")
	}
	fmt.Println(len(Set))
}

func mapSlice() {
	mapSlice := make([]map[string]string, 3)

	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "王五"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "红旗大街"

	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

func slicemap() {
	var sliceMap = make(map[string][]string, 3)
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(key)
	fmt.Println(value)
	fmt.Println(sliceMap)
}
