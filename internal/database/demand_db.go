package database

import (
	"graduation_system_api/internal/database/domain"
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/global"
)

func DeleteDemand(ids []int) error {
	db := global.GetDb()
	result := db.Table("demand").Delete(&domain.Demand{}, ids)
	if db.RowsAffected == 0 {
		return errors.New(errors.ServerError, "查无此需求")
	}
	return result.Error
}

func SelectDemandById(demandId int) (*domain.Demand, error) {
	db := global.GetDb()
	demand := new(domain.Demand)
	result := db.Table("demand").Find(demand, demandId)
	return demand, result.Error
}

func UpdateDemandItemById(demandId int, itemInfo string) error {
	db := global.GetDb()
	result := db.Table("demand").Where("id=?", demandId).Update("info", itemInfo)
	return result.Error
}

func SelectDemandList(limit, offset int) ([]domain.Demand, error) {
	db := global.GetDb()
	var demand []domain.Demand
	result := db.Table("demand").Order("id asc").Limit(limit).Offset(offset).Find(&demand)
	return demand, result.Error
}

func SelectAllDemandListCount() (int64, error) {
	db := global.GetDb()
	var demand []domain.Demand
	result := db.Table("demand").Order("id asc").Find(&demand)
	return result.RowsAffected, result.Error
}
func SelectAllDemandList() ([]domain.Demand, error) {
	db := global.GetDb()
	var demand []domain.Demand
	result := db.Table("demand").Order("id asc").Find(&demand)
	return demand, result.Error
}

func SelectDemandMyList(limit, offset int, demandPhone string) ([]domain.Demand, error) {
	db := global.GetDb()
	var demand []domain.Demand
	result := db.Table("demand").Where("user_id=?", demandPhone).Order("id asc").Limit(limit).Offset(offset).Find(&demand)
	return demand, result.Error
}
func SelectAllDemandMyList(demandPhone string) (int64, error) {
	db := global.GetDb()
	var demand []domain.Demand
	result := db.Table("demand").Where("user_id=?", demandPhone).Order("id asc").Find(&demand)
	return result.RowsAffected, result.Error
}

func CreateDemand(demand *domain.Demand) error {
	db := global.GetDb()
	result := db.Table("demand").Create(demand)
	return result.Error
}
