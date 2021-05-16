package rdb

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3310)/test_db?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
