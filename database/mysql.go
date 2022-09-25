package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mariaDB *gorm.DB

func InitDB() {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		"alan", "pass", "localhost", "3306", "home_work_z_db")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}

	sqlDB, err := db.DB()

	switch {
	case err != nil:
		fmt.Println("get db failed:", err)
		return
	case err == nil:
		sqlDB.SetConnMaxLifetime(time.Duration(2) * time.Minute)
		sqlDB.SetMaxOpenConns(30)
		sqlDB.SetMaxIdleConns(30)
	}

	mariaDB = db

	defer sqlDB.Close()
}

func GetMariaDB() (*gorm.DB, error) {

	if mariaDB == nil {
		return nil, fmt.Errorf("%s \n", "connect: db connection refused")
	}

	return mariaDB, nil
}
