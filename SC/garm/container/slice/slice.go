package main

import (
	"fmt"
)

func main() {
	// var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// var slice0 []int = arr[2:10]
	// var slice1 []int = arr[:9]
	// var slice2 []int = arr[0:]
	// var slice3 []int = arr[:]
	// var slice4 = arr[:len(arr)-1]
	// fmt.Println(slice0, slice1, slice2, slice3, slice4)

	// data := [...]int{0, 1, 2, 3, 4, 5}
	// s := data[2:4]
	// s[0] += 100
	// s[1] += 200
	// fmt.Println(s)
	// fmt.Println(data)

	/*
		对于底层数组容量是k的切片slice[i:j:k]来说
		长度:j-i
		容量:k-i

	*/
	// slice := []int{1, 2, 3, 4, 5}
	// slice1 := slice[:]
	// slice2 := slice[0:]
	// slice3 := slice[:5]

	// fmt.Println(slice)
	// fmt.Println(reflect.TypeOf(slice))
	// fmt.Println(`slice1:`, slice1)
	// fmt.Println(`slice2:`, slice2)
	// fmt.Println(`slice3:`, slice3)

	// newSlice := []int{1: 3}
	// newSlice[0] = 10
	// fmt.Println(slice)
	// fmt.Println(newSlice)

	//!新的切片和原切片共用的是一个底层数组,所以当修改的时候,底层数组的值就会被改变,所以原切片的值也改变,基于数组的切片也一样
	// newSlice := slice[1:2:3]
	// fmt.Printf(`newSlice.length:%d capacity:%d`, len(newSlice), cap(newSlice))

	//!using slice
	// fmt.Println(slice[2])
	// slice[2] = 10
	// fmt.Println(slice[2])

	//!append  末尾追加一个元素
	// newSlice := slice[1:2:3]
	// fmt.Println(newSlice)
	// // newSlice = append(newSlice, 10, 20, 30)
	// newSlice = append(newSlice, slice...)
	// fmt.Println(newSlice)
	// fmt.Println(slice)
	// //在有可用变量的时候不会创建新的切片来满足追加

	//!迭代切片
	//遍历
	// for i, v := range slice {
	// 	fmt.Printf("index:%d, value:%d\n", i, v)
	// }

	//忽略index
	// for _, v := range slice {
	// 	fmt.Printf("value:%d\n", v)
	// }

	//length迭代
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Printf("value:%d\n", slice[i])
	// }

	//!函数之间传递切片
	//切片在函数中传递是复制,修改一个index的值后,原切片的值被修改,因为共用一个底层数组
	// fmt.Printf("%p\n", &slice)
	// modify(slice)
	// fmt.Println(slice)

}
func modify(slice []int) {
	fmt.Printf("%p\n", &slice)
	slice[1] = 10
}
