package models

import (
	"index-indicators/server/app/entity"
	"log"
	"os"

	// import tzdata
	_ "time/tzdata"

	"github.com/jinzhu/gorm"
	// import mysql
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Models DB models
type Models struct {
	DB *gorm.DB
}

// NewModels is constructor
func NewModels() (*Models, error) {
	db, err := SQLConnect()
	if err != nil {
		return &Models{}, err
	}
	return &Models{DB: db}, nil
}

// SQLConnect is starting connection
func SQLConnect() (database *gorm.DB, err error) {
	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	DBNAME := os.Getenv("MYSQL_DATABASE")
	PROTOCOL := "tcp(" + os.Getenv("MYSQL_HOST") + ":3306)"

	CONNECT := (USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo")

	return gorm.Open(DBMS, CONNECT)
}

//AutoMigrate マイグレーション
func AutoMigrate() {
	log.Println("migrating database...")
	db, err := SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&entity.Fgi{}, &entity.Like{}, &entity.Ticker{}, &entity.User{})
	log.Println("finish migrate!")
}
