package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	DNS := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("APP_DATABASE_USERNAME"), os.Getenv("APP_DATABASE_PASSWORD"), os.Getenv("APP_DATABASE_PORT"), os.Getenv("APP_DATABASE_NAME"))
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}

func InitDBTest() *gorm.DB {
	DNS := fmt.Sprintf("root:@tcp(127.0.0.1:3307)/desa_maju?charset=utf8mb4&parseTime=True&loc=Local")
	db, err := gorm.Open(mysql.Open(DNS), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		log.Fatal(err.Error())
	}

	return db
}
