package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
type class struct {
	Title   string     `json:"title"`
	Student []*Student `json:"student"`
}

func main() {
	//创建班级
	c := &class{
		Title:   "一班",
		Student: make([]*Student, 0, 10),
	}
	//添加学生
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name: fmt.Sprintf("stu%d", i),
			ID:   i,
		}
		c.Student = append(c.Student, stu)
	}
	//序列化
	data, err := json.Marshal(c) //此处的data类型是string
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Printf("%s\n", data)
	//反序列化
	c1 := &class{}
	err = json.Unmarshal([]byte(data), c1)
	if err != nil {
		fmt.Println("error")
		return
	}
}
