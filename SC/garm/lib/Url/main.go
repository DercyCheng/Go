package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	s := "postgres://users:pass@host.com:5432/path?k=v#f"
	up, err := url.Parse(s)
	if err != nil {
		return
	}
	// 直接访问 scheme
	fmt.Println(up.Scheme)

	// `User` 包含了所有的认证信息，这里调用 `Username`
	// 和 `Password` 来获取独立值
	fmt.Println(up.User)
	fmt.Println(up.User.Username())
	p, _ := up.User.Password()
	fmt.Println(p)

	// `Host` 同时包括主机名和端口信息，如过端口存在的话，
	// 使用 `strings.Split()` 从 `Host` 中手动提取端口。
	fmt.Println(up.Host)
	h := strings.Split(up.Host, ":")
	fmt.Println(h[0])
	fmt.Println(h[1])

	// 路径和查询片段信息
	fmt.Println(up.Path, up.Fragment)

	// 要得到字符串中的 `k=v` 这种格式的查询参数可以使用 `RawQuery` 函数
	// 你也可以将查询参数解析为一个map
	// 已解析的查询参数 map 以查询字符串为键
	// 对应值字符串切片为值，所以如果只想得到一个键对应的第一个值
	// 将索引位置设置为 `[0]` 就行了
	fmt.Println(up.RawQuery)
	m, _ := url.ParseQuery(up.RawQuery)
	fmt.Println(m, m["k"][0])
}
