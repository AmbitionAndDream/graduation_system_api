package database

import (
	"graduation_system_api/internal/database/domain"
	"graduation_system_api/internal/global"
)

func Login(phoneNumber, password string) (*domain.User, error) {
	db := global.GetDb()
	var u domain.User
	result := db.Table("user").Where("phone=? AND pwd=?", phoneNumber, password).Find(&u)
	//通过rowCount 来判别是否有数据
	if result.RowsAffected != 0 {
		return &u, nil
	} else {
		return nil, result.Error
	}
}

func SelectPeople(limit, offset int) ([]domain.User, error) {
	db := global.GetDb()
	var user []domain.User
	result := db.Table("user").Order("id asc").Limit(limit).Offset(offset).Find(&user)
	return user, result.Error
}

func CreatePeople(user *domain.User) error {
	db := global.GetDb()
	result := db.Table("user").Create(user)
	return result.Error
}
