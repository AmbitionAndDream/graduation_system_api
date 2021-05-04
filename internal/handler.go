package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal/auth"
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/fusion"
	"graduation_system_api/internal/global"
	"graduation_system_api/internal/util"
	"net/http"

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
	logrus.Infof("request token is %s", token)
	if token == "" {
		util.BuildFailedResp(ctx, errors.TokenMissError, errors.New(errors.TokenMissError, "token miss errored"))
		return
	}
	if err := auth.CheckToken(token); err != nil {
		logrus.Errorf("check token error :%s", err.Error())
		er := err.(*errors.Error)
		util.BuildFailedResp(ctx, er.Code(), er)
		return
	}

	// 判断kind 找对应handle
	kind := ctx.Param("kind")
	var resp interface{}
	var err error
	switch kind {
	case global.PeoPle:
		resp, err = fusionHandler.HandlePeopleEvent(ctx)
	case global.Business:
		resp, err = fusionHandler.HandleBusinessEvent(ctx)
	case global.Demand:
		resp, err = fusionHandler.HandleDemandEvent(ctx)
	case global.Bug:
		resp, err = fusionHandler.HandleBugEvent(ctx)
	default:
		logrus.Errorf("url param error ,url:%s", ctx.Request.URL.String())
		util.BuildFailedResp(ctx, http.StatusBadRequest, errors.New(http.StatusBadRequest, "url 参数有误"))
		return
	}
	if err != nil {
		er := err.(*errors.Error)
		util.BuildFailedResp(ctx, er.Code(), er)
		return
	}
	logrus.Infof("response data is %v", resp)
	util.BuildSuccessResp(ctx, resp)
	return

}
