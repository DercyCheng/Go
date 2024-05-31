package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}

var (
	CREATE_TABLE = "CREATE TABLE student(" +
		"sid INT(10) NOT NULL AUTO_INCREMENT," +
		"sname VARCHAR(64) NULL DEFAULT NULL," +
		"age INT(10) DEFAULT NULL,PRIMARY KEY (sid))" +
		"ENGINE=InnoDB DEFAULT CHARSET=utf8;"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 建立数据库连接
func setupConnect() *sql.DB {
	db, err := sql.Open("mysql", "root:kwok2012@tcp(127.0.0.1:3306)/test?charset=utf8")
	errorHandler(err)
	return db
}

// 创建表
func CreateTable(db *sql.DB, sql string) {
	_, err := db.Exec(sql)
	errorHandler(err)
}

var INSERT_DATA = `INSERT INTO student(sid,sname,age) VALUES(?,?,?);`

// 插入数据
func Insert(db *sql.DB) {
	db.Exec(INSERT_DATA, 1, "唐僧", 30)
}

var UPDATE_DATA = `UPDATE student SET age=28 WHERE sname="唐僧";`

// 修改数据
func Update(db *sql.DB) {
	db.Exec(UPDATE_DATA)

}

var DELETE_DATA = `DELETE FROM student WHERE age>=30`

// 删除记录
func Delete(db *sql.DB) {
	db.Exec(DELETE_DATA)
}

var DELETE_TABLE = `DROP TABLE student;`

// 删除表
func DeleteTable(db *sql.DB) {
	db.Exec(DELETE_TABLE)
}

var QUERY_DATA = `SELECT * FROM student;`

// 查询数据
func Query(db *sql.DB) {
	rows, err := db.Query(QUERY_DATA)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var name string
		var id int
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s is %d\n", name, age)
	}
}
func MultiInsert(db *sql.DB) {
	// 批量数据插入
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}
	values := [][]interface{}{{2, "孙悟空", 500}, {3, "猪八戒", 200}, {4, "沙悟净", 100}}
	stmt, err := tx.Prepare("INSERT INTO student(sid,sname,age) VALUES(?,?,?);")
	for _, val := range values {
		_, err := stmt.Exec(val...)
		if err != nil {
			fmt.Printf("INSERT failed：%v", err)
			tx.Rollback()
		}
	}
	tx.Commit()
}

func main() {
	// 建立数据连接
	db := setupConnect()
	// 创建数据库表
	CreateTable(db, CREATE_TABLE)
	// 插入数据
	Insert(db)
	// 查询数据
	Query(db)
	// 删除数据
	Delete(db)
	// 插入数据
	Insert(db)
	// 修改数据
	Update(db)
	// 查询数据
	Query(db)
	// 删除表
	DeleteTable(db)

	// 关闭数据库连接
	err := db.Close()
	handleErr(err)
}
