package handler

import (
	"github.com/gin-gonic/gin"
	"graduation_system_api/internal/auth"
	"net/http"
)


func Handler(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Handler",
	})
}

func Handler_02(ctx *gin.Context){
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Handler_02",
	})
}
func Handler_03(ctx *gin.Context){
	if err:=auth.CheckToken();err !=nil{
		ctx.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Handler_02",
	})
}