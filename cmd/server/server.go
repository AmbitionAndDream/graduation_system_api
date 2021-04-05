package main

import (
	"github.com/gin-gonic/gin"
	"graduation_system_api/internal/handler"
)

func main() {
	r := gin.Default()

	g:=r.Group("/action_h")
	g.GET("/:act",handler.Handler)
	f:=r.Group("/action_m")
	f.POST("/:act2", handler.Handler_02)
	_ = r.Run()
}
