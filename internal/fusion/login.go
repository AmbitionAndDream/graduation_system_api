package fusion

import (
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/global"
)

func login(phoneNumber, password string) (int, error) {
	//db select
	if phoneNumber == "13999271543" && password == "88888888" {
		return global.Administrator, nil
	} else if phoneNumber == "13999271543" && password != "88888888" {
		return -2, errors.New(errors.UserUndefinedError, "UserUndefinedError")
	} else {
		return -2, errors.New(errors.ServerError, "服务错误")
	}
}
