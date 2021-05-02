package database

import (
	"graduation_system_api/internal/domain/user"
	"graduation_system_api/internal/global"
)

func Login(phoneNumber, password string) (*user.User, error) {
	db := global.GetDb()
	var u user.User
	result := db.Table("user").Where("phone=? AND pwd=?", phoneNumber, password).Find(&u)
	//通过rowCount 来判别是否有数据
	if result.RowsAffected != 0 {
		return &u, nil
	} else {
		return nil, result.Error
	}
}
