package main

import (
	"GenericToolLibrary/slice"
	"fmt"
)

func main() {
	s := []int{1, 2, 3}
	e := 4
	slice.Add(s, e)
	fmt.Println(s)
}
