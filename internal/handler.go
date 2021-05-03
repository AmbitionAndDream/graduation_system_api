package internal

import (
	"github.com/gin-gonic/gin"
	"graduation_system_api/internal/auth"
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/fusion"
	"graduation_system_api/internal/util"
	// "log"
)

var fusionHandler = fusion.GetFusionHandler()

func Login(ctx *gin.Context) {
	resp, err := fusionHandler.HandlerLoginEvent(ctx)
	if err != nil {
		er := err.(*errors.Error)
		util.BuildFailedResp(ctx, er.Code(), er)
		return
	}
	util.BuildSuccessResp(ctx, resp)
}

func Handler(ctx *gin.Context) {
	//从Header 获取token 进行check
	token := ctx.GetHeader("authorized")
	if token == "" {
		util.BuildFailedResp(ctx, errors.TokenMissError, errors.New(errors.TokenMissError, "token miss errored"))
	}
	if err := auth.CheckToken(token); err != nil {
		er := err.(*errors.Error)
		util.BuildFailedResp(ctx, er.Code(), er)
	}

	// 判断kind 找对应handle
	kind := ctx.Param("kind")
	switch kind {
	case "people":
		resp, err := fusionHandler.HandlePeopleEvent(ctx)
		if err != nil {
			return
		}
		util.BuildSuccessResp(ctx, resp)
		return
	case "business":
		resp, err := fusionHandler.HandleBusinessEvent(ctx)
		if err != nil {
			er := err.(*errors.Error)
			util.BuildFailedResp(ctx, er.Code(), er)
			return
		}
		util.BuildSuccessResp(ctx, resp)
		return
	}
}
