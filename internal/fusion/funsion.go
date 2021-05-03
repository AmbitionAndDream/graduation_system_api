package fusion

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal/auth"
	"graduation_system_api/internal/domain"
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/global"
	"net/http"
	"reflect"
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
		return nil, err
	}
	//生成token
	token := auth.GetToken(role, u.PhoneNumber)
	resp = struct {
		Token string `json:"token"`
	}{Token: token}
	return
}

// 一个kind  对应一个Handle 多个action
func (f *newFusionHandler) HandleBusinessEvent(ctx *gin.Context) (resp interface{}, err error) {
	action := ctx.Param("action")
	switch action {
	case global.Create:
		name := &domain.RequestBusinessName{}
		if err := ctx.BindJSON(name); err != nil {
			logrus.Errorf("business create param invalid ,the param is %v", name)
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		if err = createBusiness(name.Name); err != nil {
			return
		}
		return "business create successful", nil
	case global.Delete:
		ids := new(domain.RequestDeleteBusinessId)
		if err = ctx.ShouldBind(ids); err != nil {
			logrus.Errorf("business del param invalid ,the param is %v", ids)
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		if err = deleteBusiness(ids.BusinessId); err != nil {
			return
		}

		return "business delete successful", nil
	case global.Select:
		limit := ctx.Query("limit")
		offset := ctx.Query("offset")
		if !checkParam(limit, offset) {
			logrus.Errorf("business list param invalid ,the param limit: %s,offset:%s", limit, offset)
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return selectBusiness(limit, offset)
	}
	return
}

func (f *newFusionHandler) HandlePeopleEvent(ctx *gin.Context) (resp interface{}, err error) {
	// del
	action := ctx.Param("action")
	//if action == "del" {
	//	err = peopleDel(ctx.PostForm("phone"))
	//	if(err != nil) {
	//		logrus.Errorf("del people error: %s",err.Error())
	//
	//		//util.BuildFailedResp(ctx, http.StatusBadRequest, err)
	//		return
	//	}
	//}
	switch action {
	case global.Select:
		limit := ctx.Query("limit")
		offset := ctx.Query("offset")
		if !checkParam(limit, offset) {
			logrus.Errorf("people list param invalid ,the param limit: %s,offset:%s", limit, offset)
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return selectPeople(limit, offset)
	case global.Create:
		u := new(domain.RequestPeople)
		if err = ctx.ShouldBind(u); err != nil {
			logrus.Errorf("parse people list params failed ,error: %s", err.Error())
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		if err = createPeople(u); err != nil {
			return
		}
		return "people create successful", nil
	case global.Delete:
		phone := &domain.RequestPeoplePhone{}
		if err := ctx.BindJSON(phone) ; err !=nil {
			logrus.Errorf("people del params invalid, param:  %v",phone )
			return nil,errors.New(errors.ParamInvalidError, "param invalid error")
		}
		if err = peopleDel(phone.Phone); err != nil{
			logrus.Errorf("people del error: %s", err.Error())
			return
		}
		return "people del successful", nil
	}

	return
}

func checkParam(date ...interface{}) bool {
	if len(date) == 0 {
		return false
	}
	k := reflect.ValueOf(date[0]).Kind()
	switch k {
	case reflect.String:
		for _, element := range date {
			if len(element.(string)) == 0 {
				return false
			}
		}
	}
	return true
}
