package models

import (
	"fmt"
	"index-indicator-apis/server/app/entity"
	"os"

	"github.com/jinzhu/gorm"
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
	USER := "iia"
	PASS := "iia"
	PROTOCOL := "tcp(" + os.Getenv("MYSQL_HOST") + ":3306)"
	DBNAME := "index_indicator_apis"

	CONNECT := (USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo")

	return gorm.Open(DBMS, CONNECT)
}

//AutoMigrate マイグレーション
func AutoMigrate() {
	fmt.Println("migrating database...")
	db, err := SQLConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.AutoMigrate(&entity.Fgi{}, &entity.Like{}, &entity.Ticker{}, &entity.User{})
	fmt.Println("finish migrate!")
}
