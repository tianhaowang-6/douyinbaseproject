package db

import (
	"douyin/cmd/feed/dal/query"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Q *query.Query

func Init() {
	dsn := "root:admin@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	query.SetDefault(db)
	Q = query.Q
	if err != nil {
		panic(err)
	}
}
