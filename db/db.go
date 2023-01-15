package db

import (
	"log"
	"projectA/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	c := config.GetConfig()
	dsn := "host=" + c.GetString("pg.host") +
		" user=" + c.GetString("pg.user") +
		" password=" + c.GetString("pg.pass") +
		" dbname=" + c.GetString("pg.name") +
		" sslmode=require" +
		" options=project=ep-floral-haze-264573"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("error on initializing database: ", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
