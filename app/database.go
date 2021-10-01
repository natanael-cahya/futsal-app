package app

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	dbs := "root:@tcp(127.0.0.1:3306)/db_futsal?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dbs), &gorm.Config{})

	if err != nil {
		panic("Error Mysql Database NotFound")
	} else {
		fmt.Println("Connected to Mysql Database")
	}
	return db
}
