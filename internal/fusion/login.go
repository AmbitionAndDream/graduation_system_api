package fusion

import (
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/global"
	"graduation_system_api/internal/db"
	"fmt"
)
type User struct {
	Id int
	Phone string
	Name string
	Pwd string
	Is_admin int
	Role_type int
}

func login(phoneNumber, password string) (int, error) {
	//db select
	conn := db.GetDb()
	var uu User
	conn.Where("phone = ?", phoneNumber).Find(&uu)
	defer conn.Close()

	if password == uu.Pwd {
		if(uu.Is_admin == global.Administrator) {
			return global.Administrator, nil
		}
		return global.NonAdministrator, nil
	} else if password != uu.Pwd {
		return -2, errors.New(errors.UserUndefinedError, "UserUndefinedError")
	} else {
		return -2, errors.New(errors.ServerError, "服务错误")
	}

}
