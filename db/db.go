package db

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3307)/gokapster?parseTime=true"

	DB, errDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatalf("DB connection failed %v:", errDB)
	}

	return DB
}
