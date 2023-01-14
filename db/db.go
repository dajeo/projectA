package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func e(key string) string {
	return os.Getenv("PG_" + key)
}

func InitDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=require options=project=ep-floral-haze-264573", e("HOST"), e("USER"), e("PASS"), e("NAME"))
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
}

func GetDB() *gorm.DB {
	return db
}
