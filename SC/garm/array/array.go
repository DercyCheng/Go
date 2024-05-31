package main

import (
	"fmt"
)

// func TestArray(t *testing.T) {
func main() {

	//!
	// a := [5]string{0: "1", 1: "2"}
	// fmt.Println(a[0])

	//!...拓展语法
	// b := [...]int{1, 2, 3, 4}
	// fmt.Println(b)

	//!未定义的值为0
	// c := [...]int{0: 1, 4: 1, 9: 1}
	// fmt.Println(c)
	//

	//!length
	// d := [...]float64{67, 7, 89.8, 21, 78}
	// fmt.Println(len(d))

	//!遍历数组
	// e := [...]float64{67.7, 89.8, 21, 78}
	// for i := 0; i < len(e); i++ {
	// 	fmt.Printf("%d th element of e is %.2f\n", i, e[i])
	// }

	//!range方法遍历数组
	f := [...]float64{67.7, 79.2, 46.6}
	for _, v := range f {
		fmt.Println(v)
	}

	//多维数组
	g := [3][4]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9, 10, 11}}
	fmt.Println(g[0][1])

}
