package fusion

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal/auth"
	"graduation_system_api/internal/errors"
	"net/http"
)

type newFusionHandler struct {
}

func GetFusionHandler() *newFusionHandler {
	return &newFusionHandler{}
}

func (f *newFusionHandler) HandlerLoginEvent(ctx *gin.Context) (resp interface{}, err error) {
	type User struct {
		PhoneNumber string `json:"phoneNumber" binding:"required"`
		PassWord    string `json:"passWord" binding:"required"`
	}
	u := new(User)
	if err = ctx.ShouldBind(u); err != nil {
		logrus.Errorf("parse user login params failed ,error: %s", err.Error())
		err = errors.New(http.StatusBadRequest, "url 参数有误")
		return
	}

	//登陆
	var role int
	if role, err = login(u.PhoneNumber, u.PassWord); err != nil {
		logrus.Errorf("login failed:%s", err.Error())
		return
	}
	//生成token
	token := auth.GetToken(role, u.PhoneNumber)
	resp = struct {
		Token string `json:"token"`
	}{Token: token}
	return
}

func (f *newFusionHandler) HandleFrontierEvent(ctx *gin.Context) {

}
