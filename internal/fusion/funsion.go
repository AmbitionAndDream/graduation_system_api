package fusion

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"graduation_system_api/internal/auth"
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/util"
	"io/ioutil"
	"net/http"
)

type newFusionHandler struct {
}

func GetFusionHandler() *newFusionHandler {
	return &newFusionHandler{}
}

func (f *newFusionHandler) HandlerLoginEvent(ctx *gin.Context) (resp interface{}, err error) {
	//获取登陆信息
	type user struct {
		PhoneNumber string `json:"phoneNumber"`
		PassWord    string `json:"passWord"`
	}
	u := new(user)
	b, _ := ioutil.ReadAll(ctx.Request.Body)
	if err = json.Unmarshal(b, u); err != nil {
		util.BuildFailedResp(ctx, http.StatusBadRequest, errors.New(http.StatusBadRequest,"url参数有误"))
		return
	}
	//登陆
	var role int
	if role, err = login(u.PhoneNumber, u.PassWord); err != nil {
		if er,ok:=err.(*errors.Error);ok{
			util.BuildFailedResp(ctx, er.Code(), er)
		}
		return
	}
	//生成token
    token :=auth.GetToken(role,u.PhoneNumber)
    resp = struct {
		Token string `json:"token"`
	}{Token: token}
	return
}

func (f *newFusionHandler) HandleFrontierEvent(ctx *gin.Context) {

}
