package handler

import (
	"delos-farm-API/Farm"

	"github.com/gin-gonic/gin"
)

func CreateFarm(ctx *gin.Context) {
	var farm Farm.Farms
	err := ctx.ShouldBindJSON(&farm)
	if err != nil {
		panic(err)
	}

	ctx.JSON(200, gin.H{
		"farm": farm.Farm,
		"ID":   farm.ID,
		"pond": farm.Pond,
	})
}
