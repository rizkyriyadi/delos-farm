package repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dbHandler struct {
	db *gorm.DB
}

func NewdbHandler() (*dbHandler, error) {
	dsn := "root:2113@tcp(127.0.0.1:3306)/delos_farm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return &dbHandler{db}, err

}
