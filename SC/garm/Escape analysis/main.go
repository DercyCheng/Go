package main

import "fmt"

// 了解就好,go自行决定

// !变量逃逸
// func dummy(b int) int {
// 	c := b
// 	return c
// }
// func void() {}

// func main.exe() {

// 	var a int
// 	void()
// 	fmt.Println(a, dummy(0))
// }

/*
$go run -gcflags "-m -l" session.go

# command-line-arguments
.\session.go:14:13: ... argument does not escape
.\session.go:14:13: a escapes to heap
.\session.go:14:22: dummy(0) escapes to heap
0 0
*/

// !取地址逃逸
type data struct{}

func dummy() *data {
	c := data{}
	return &c
}
func main() {
	fmt.Println(dummy())
}

// 将c分配到堆上 然后垃圾回收掉c
/*
$go run -gcflags "-m -l" session.go

# command-line-arguments
.\session.go:33:2: moved to heap: c
.\session.go:37:13: ... argument does
not escape
&{}
*/
