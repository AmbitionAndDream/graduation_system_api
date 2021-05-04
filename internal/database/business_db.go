package database

import (
	"graduation_system_api/internal/database/domain"
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/global"
)

/*
	prod 可以这么写，但是单测的时候不能
*/
//var db *gorm.DB
//
//func init() {
//	db = global.GetDb()
//}
func CreateBusiness(name string) error {
	db := global.GetDb()
	business := &domain.Business{
		BusinessName: name,
	}
	result := db.Table("business").Select("name").Create(business)
	return result.Error
}
func SelectAllBusiness() (int64, error) {
	db := global.GetDb()
	var business []domain.Business
	result := db.Table("business").Order("id asc").Find(&business)
	return result.RowsAffected, result.Error
}

func SelectBusiness(limit, offset int) ([]domain.Business, error) {
	db := global.GetDb()
	var business []domain.Business
	result := db.Table("business").Order("id asc").Limit(limit).Offset(offset).Find(&business)
	return business, result.Error
}

func DeleteBusiness(ids []int) error {
	db := global.GetDb()
	result := db.Table("business").Delete(&domain.Business{}, ids)
	if db.RowsAffected == 0 {
		return errors.New(errors.ServerError, "查无此业务线")
	}
	return result.Error
}
