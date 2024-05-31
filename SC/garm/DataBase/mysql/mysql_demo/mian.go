package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // init()
)

// Go连接MySQL示例

var db *sql.DB // 是一个连接池对象

func initDB() (err error) {
	// 数据库信息
	// 用户名:密码@tcp(ip:端口)/数据库的名字
	dsn := "root:kwok2012@tcp(127.0.0.1:3306)/sql_test"
	// 连接数据库
	db, err = sql.Open("mysql", dsn) // 不会校验用户名和密码是否正确
	if err != nil {                  // dsn格式不正确的时候会报错
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(5)  // 设置最大空闲连接数
	return
}

type user struct {
	id   int
	name string
	age  int
}

// 查询单个记录
func queryOne(id int) {
	var u user
	// 1. 写查询单条记录的sql语句
	sqlStr := `select id, name, age from users where id=?;`
	// 2. 执行并拿到结果
	// 必须对rowObj对象调用Scan方法,因为该方法会释放数据库链接 // 从连接池里拿一个连接出来去数据库查询单条记录
	db.QueryRow(sqlStr, id).Scan(&u.id, &u.name, &u.age)
	// 打印结果
	fmt.Printf("one: %#v\n", u)
}

// 查询多条 遍历
func queryMore(n int) {
	// 1. SQL语句
	sqlStr := `select id, name, age from users where id > ?;`
	// 2. 执行
	rows, err := db.Query(sqlStr, n)
	if err != nil {
		fmt.Printf("exec %s query failed, err:%v\n", sqlStr, err)
		return
	}
	// 3. 一定要关闭rows
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			panic(err)
		}
	}(rows)
	// 4. 循环取值
	for rows.Next() {
		var u user
		//与单行的scan不同
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
		}
		fmt.Printf("more: %#v\n", u)
	}
}

// 插入数据
func insert() {
	// 1. 写SQL语句
	sqlStr := `insert into users(name, age) values("a", 28)`
	// 2. exec
	ret, err := db.Exec(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	// 如果是插入数据的操作，能够拿到插入数据的id
	id, err := ret.LastInsertId()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Println("id:", id)
}

// 更新操作
func updateRow(newAge int, id int) {
	sqlStr := `update users set age=? where id > ?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Printf("get id failed,err:%v\n", err)
		return
	}
	fmt.Printf("更新了%d行数据\n", n)
}

// 删除
func deleteRow(id int) {
	sqlStr := `delete from users where id=?;`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("删除了%d行\n", n)
}

// 预处理方式插入多条数据
func prepareInsert() {
	sqlStr := `insert into users(name,age)values(?,?);`
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var m = map[string]int{
		"b": 30,
		"c": 32,
		"d": 72,
		"e": 40,
	}
	for k, v := range m {
		_, err := stmt.Exec(k, v)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	fmt.Println("预处理插入成功")
}
func transctionSql() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return
	}
	sqlStr1 := `update users set age=age-2 where id=1`
	sqlStr2 := `update users set age=age+2 where id=2`
	_, err = db.Exec(sqlStr1)
	if err != nil {
		tx.Rollback()
		log.Fatalln("sqlStr1 error Must rollback")
		return
	}
	_, err = db.Exec(sqlStr2)
	if err != nil {
		tx.Rollback()
		log.Fatalln("sqlStr2 error Must rollback")
		return
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Fatalln("commit error")
		return
	}
	fmt.Println("执行事务成功")
}
func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init DB failed, err:%v\n", err)
	}
	fmt.Println("连接数据库成功!")
	//queryOne(52)
	//queryMore(2)
	//insert()
	//updateRow(9000, 2)
	//deleteRow(30)
	//prepareInsert()
	transctionSql()
}
