package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql" // init()
)

//sqlx是sql的超集 除连接操作 与sql操作基本一致
// Go连接MySQL示例

var db *sqlx.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	// 用户名:密码@tcp(ip:端口)/数据库的名字
	dsn := "root:kwok2012@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	// 也可以使用MustConnect连接不成功就panic
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

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
		return
	}
	sqlStr1 := `select id, name, age from users where id=1`
	var u user

	err = db.Get(&u, sqlStr1)
	if err != nil {
		return
	}
	fmt.Printf("u:%#v\n", u)

	var userList []user
	sqlStr2 := `select id,name, age from users`
	err = db.Select(&userList, sqlStr2)
	if err != nil {
		fmt.Printf("select failed, err:%v\n", err)
		return
	}
	fmt.Printf("userList:%#v\n", userList)
}
