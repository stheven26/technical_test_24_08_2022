package db

import (
	"technical_test_24_08_2022/models"

	"technical_test_24_08_2022/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func SetupDB() {
	config := config.Env()

	dsn := config.DB_USERNAME + ":" + config.DB_PASSWORD + "@(" + config.DB_HOST + ")/" + config.DB_NAME + "?charset=utf8&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Activities{})
	db.AutoMigrate(&models.ToDo{})
}

func GetConnectionDB() *gorm.DB {
	return db
}
