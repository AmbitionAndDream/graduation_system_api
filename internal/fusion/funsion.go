package fusion

import (
	"github.com/gin-gonic/gin"
	"graduation_system_api/internal/auth"
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/util"
	"log"
	"net/http"
)

type newFusionHandler struct {
}

func GetFusionHandler() *newFusionHandler {
	return &newFusionHandler{}
}

func (f *newFusionHandler) HandlerLoginEvent(ctx *gin.Context) (resp interface{}, err error) {
	//获取登陆信息
	//type user struct {
	//	PhoneNumber string `json:"phoneNumber"`
	//	PassWord    string `json:"passWord"`
	//}
	//u := new(user)
	//b, _ := ioutil.ReadAll(ctx.Request.Body)
	//if err = json.Unmarshal(b, u); err != nil {
	//	util.BuildFailedResp(ctx, http.StatusBadRequest, errors.New(http.StatusBadRequest,"url参数有误"))
	//	return
	//}
	if ctx.Request.Method==http.MethodPost {
		type User struct {
			PhoneNumber string `json:"phone_number" binding:"required"`
			PassWord    string `json:"pass_word" binding:"required"`
		}
		u := new(User)
		if err = ctx.ShouldBind(&u); err != nil {
			util.BuildFailedResp(ctx, http.StatusBadRequest, errors.New(http.StatusBadRequest, "url参数有误"))
			return
		}

		log.Println(u)

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
			Is_admin int `json:"is_admin"`
		}{Token: token, Is_admin: role}
	}
	if ctx.Request.Method==http.MethodGet {
		type User1 struct {
			PhoneNumber string `form:"phoneNumber1" binding:"required"`
			PassWord    string `form:"passWord1" binding:"required"`
		}
		u1 := new(User1)
		if err = ctx.ShouldBind(&u1); err != nil {
			util.BuildFailedResp(ctx, http.StatusBadRequest, errors.New(http.StatusBadRequest, "url参数有误"))
			return
		}
		log.Println(u1)

	}
	return
}

func (f *newFusionHandler) HandleFrontierEvent(ctx *gin.Context) {

}
