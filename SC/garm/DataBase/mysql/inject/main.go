package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql" // init()
)

//!我们任何时候都不应该自己拼接SQL语句！

// Go连接MySQL示例
var db *sqlx.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	// 用户名:密码@tcp(ip:端口)/数据库的名字
	dsn := "root:kwok2012@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数
	return
}

type user struct {
	ID   int
	Name string
	Age  int
}

// SQL注入

// sql注入示例
func sqlInjectDemo(name string) {
	// 自己拼接SQL语句的字符串
	//!不要直接使用带生成参数的语句,会产生sql注入风险
	sqlStr := fmt.Sprintf("select id, name, age from users where name='%s'", name)

	fmt.Printf("SQL:%s\n", sqlStr)

	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("exec failed, err:%v\n", err)
		return
	}
	for _, u := range users {
		fmt.Printf("users:%#v\n", u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	// SQL注入的几种示例
	//sqlInjectDemo("图朝阳")
	//常用注入方式 "test or 1=1 "即可获取sql所有数据
	//sqlInjectDemo("xxx' or 1=1 #")
	//连接也成立
	sqlInjectDemo("xxx' union select * from users #")
}
