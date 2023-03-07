package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	conn := "root:KdSXfshvP@tcp(localhost:3306)/sakila"

	db, err := gorm.Open(
		mysql.Open(conn+"?parseTime=true"),
		&gorm.Config{},
	)

	if err != nil {
		panic(err)
	}

	DB = db
}
