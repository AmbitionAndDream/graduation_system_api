package main

import (
	"github.com/gin-gonic/gin"
	"graduation_system_api/internal"
)

func main() {
	r := gin.Default()
	r.Any("/login", internal.Login)
	g:=r.Group("/api")
	g.Any("/:kind/:action", internal.Handler)
	_ = r.Run()
}
