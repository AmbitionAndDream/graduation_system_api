package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"graduation_system_api/internal/fusion"
	"graduation_system_api/internal/util"
	"net/http"
	// "log"
)

var fusionHandler = fusion.GetFusionHandler()

func Login(ctx *gin.Context) {
	resp,err := fusionHandler.HandlerLoginEvent(ctx)
	if err != nil {
		return
	}
	util.BuildSuccessResp(ctx,resp)
}


func Handler(ctx *gin.Context) {
	//从Header 获取token 进行check
	//fusionHandler.HandleFrontierEvent(ctx)
  // 判断kind 找对应handle
	kind := ctx.Param("kind")
	if kind == "people" {
		resp,err := fusionHandler.HandlePeopleEvent(ctx)
		if err != nil {
			return
		}
		util.BuildSuccessResp(ctx,resp)
		return
	}

	if http.MethodGet == ctx.Request.Method {
		util.BuildSuccessResp(ctx, struct {
			Name string `json:"name"`
			Age int `json:"age"`
		}{"heihei",15})

	}else if http.MethodPost == ctx.Request.Method {
		util.BuildFailedResp(ctx,702,errors.New("token expired"))
	}else {
		//构造response

	}
}

