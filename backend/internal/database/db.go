package database

import (
	"fmt"
	"gorm.io/gorm"
	"owl-backend/internal/config"
	"time"

	"gorm.io/driver/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	var err error
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=UTC",
		config.C.DBUser,
		config.C.DBPassword,
		config.C.DBHost,
		config.C.DBPort,
		config.C.DBName,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v mysql init successfully\n", time.Now())
}
