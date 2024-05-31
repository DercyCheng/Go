package string_test

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	s := "中"
	c := []rune(s)
	//编码规则不一致
	fmt.Printf("中的unicode %x\n", c[0])
	fmt.Printf("中的UTF8 %x\n", s)
}

func TestStringToRune(t *testing.T) {
	s := "helloworld"
	for _, v := range s {
		fmt.Printf("%[1]c %[1]x", v)
	}
}
