package db

import (
	"log"

	"github.com/Fermekoo/orderin-api/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(config utils.Config) *gorm.DB {
	dsn := config.DSN

	DB, errDB := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errDB != nil {
		log.Fatalf("DB connection failed %v:", errDB)
	}

	return DB
}
