package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	// 包中的函数  不是字符串对象自身的方法
	// 调用时将字符串作为第一个参数进行传递
	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// 虽然不是 `strings` 的一部分
	// 但是仍然值得一提的是获取字符串长度和通过索引获取一个字符的机制
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
