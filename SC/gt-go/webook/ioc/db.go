package ioc

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gt-go/webook/config"
	"gt-go/webook/internal/repository/dao"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
		panic(err)
	}
	err = dao.InitTables(db)
	if err != nil {
		panic(err)
	}
	return db
}