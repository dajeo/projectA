package utils

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var Db *gorm.DB

func e(key string) string {
	return os.Getenv("PG_" + key)
}

func InitDb() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", e("HOST"), e("USER"), e("PASS"), e("NAME"))
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	Db = db
}
