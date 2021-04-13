package internal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"graduation_system_api/internal/fusion"
	"graduation_system_api/internal/util"
	"net/http"
)

var fusionHandler =fusion.GetFusionHandler()

func Login(ctx *gin.Context){
	resp,err := fusionHandler.HandlerLoginEvent(ctx)
	if err != nil{
		return
	}
	util.BuildSuccessResp(ctx,resp)
}


func Handler(ctx *gin.Context){
	//从Header 获取token 进行check
	//fusionHandler.HandleFrontierEvent(ctx)
<<<<<<< HEAD



=======
>>>>>>> 95cb1272897492c242c503157d93aa1a8055cd25
	if http.MethodGet == ctx.Request.Method {
		util.BuildSuccessResp(ctx, struct {
			Name string `json:"name"`
			Age int `json:"age"`
		}{"heihei",15})

	}else if http.MethodPost == ctx.Request.Method{
		util.BuildFailedResp(ctx,702,errors.New("token expired"))

	}else {
		//构造response

	}
}

