package fusion

import (
	"github.com/sirupsen/logrus"
	//"graduation_system_api/internal/db"
	//"graduation_system_api/internal/errors"
	//"graduation_system_api/internal/global"
	"graduation_system_api/internal/database"
	"graduation_system_api/internal/errors"
)

func login(phoneNumber, password string) (int, string, error) {
	u, err := database.Login(phoneNumber, password)
	if err != nil {
		logrus.Errorf("select db failed error is :%s",err.Error())
		return -2,"", errors.New(errors.ServerError,"Server failed")
	}
	if u == nil {
		return -2, "", errors.New(errors.UserUndefinedError, "用户密码错误或未注册账号")
	} else {
		return u.IsAdmin, u.Name, nil
	}
	//db select
	//conn := db.GetDb()
	//var uu User
	//conn.("phone = ?", phoneNumber).Find(&uu)
	//
	//if password == uu.Pwd {
	//	if uu.Is_admin == global.Administrator {
	//		return global.Administrator, nil
	//	}
	//	return global.NonAdministrator, nil
	//} else if password != uu.Pwd {
	//	return -2, errors.New(errors.UserUndefinedError, "UserUndefinedError")
	//} else {
	//	return -2, errors.New(errors.ServerError, "服务错误")
	//}
}
