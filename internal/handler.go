package internal

import (
	"github.com/gin-gonic/gin"
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/fusion"
	"graduation_system_api/internal/util"
	"net/http"
)

var fusionHandler =fusion.GetFusionHandler()

func Login(ctx *gin.Context){
	resp,err := fusionHandler.HandlerLoginEvent(ctx)
	if err != nil{
		er := err.(*errors.Error)
		util.BuildFailedResp(ctx,er.Code(),er)
		return
	}
	util.BuildSuccessResp(ctx,resp)
}


func Handler(ctx *gin.Context){
	//从Header 获取token 进行check
	//fusionHandler.HandleFrontierEvent(ctx)
	if http.MethodGet == ctx.Request.Method {
		util.BuildSuccessResp(ctx, struct {
			Name string `json:"name"`
			Age int `json:"age"`
		}{"heihei",15})

	}else if http.MethodPost == ctx.Request.Method{
		//util.BuildFailedResp(ctx,702,errors.New("token expired"))

	}else {
		//构造response

	}
}

