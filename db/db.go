package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDb() (err error) {
	dsn := "root:sp112233@tcp(127.0.0.1:53306)/FocusLZ2023a?charset=utf8mb4"

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Sprintf("数据库连接失败，%v\n", err)
	}
	return
}
