package main

import (
	"delos-farm-API/handler"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:2113@tcp(127.0.0.1:3306)/delos_farm?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	// Router
	router := gin.Default()

	// Method GET
	// router.GET("/", func(ctx *gin.Context) {

	// 	ctx.JSON(http.StatusOK, gin.H{
	// 		"name": "Rizky Riyadi",
	// 		"fav":  "Bakso",
	// 	})
	// })
	router.POST("/createFarm", handler.CreateFarm)
	router.Run()

}
