package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal"
	"graduation_system_api/internal/global"
)

func main() {
	defer logrus.Error("application start failed")

	r := gin.Default()

	//加载配置文件
	var config = flag.String("config","","server config")
	if !flag.Parsed(){
		flag.Parse()
	}

	options := &global.Options{Log: true}
	if err := global.InitConfig(*config,options);err != nil{
		logrus.Errorf("init server config error %s",err.Error())
		panic(err)
	}

	//配置路由信息
	r.Any("/login", internal.Login)
	g:=r.Group("/api")
	g.Any("/:kind/:action", internal.Handler)
	if err := r.Run();err != nil{
		panic(err)
	}
}
