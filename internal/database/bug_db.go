package database

import (
	"graduation_system_api/internal/database/domain"
	req "graduation_system_api/internal/domain"
	"graduation_system_api/internal/global"
)

func SolveBug(bugId, status, solveType int) error {
	db := global.GetDb()
	result := db.Table("bug").Where("id=?", bugId).Updates(map[string]interface{}{"status": status, "solve_type": solveType})
	return result.Error
}

func CreateBug(bug *domain.Bug) error {
	db := global.GetDb()
	result := db.Table("bug").Create(bug)
	return result.Error
}

func SelectBugListHandler(bug *req.BugList) ([]domain.Bug, error) {
	db := global.GetDb()
	var b []domain.Bug
	db = db.Table("bug")
	if bug.Type != 0  {
		db = db.Where("type=?", bug.Type)
	}
	if bug.BeginTime != 0 && bug.EndTime != 0 {
		db = db.Where("begin_time>? AND begin_time<?", bug.BeginTime, bug.EndTime)
	}
	if bug.Status != 0 {
		db = db.Where("status=?", bug.Status)
	}
	if bug.PeoplePhone != ""{
		db = db.Where("handler_id=?", bug.PeoplePhone)
	}
	if bug.Opportunity != 0  {
		db = db.Where("opportunity=? ", bug.Opportunity)
	}
	if bug.PriorityStatus != 0  {
		db = db.Where("priority_status=?", bug.PriorityStatus)
	}
	if bug.SolveType != 0  {
		db = db.Where("solve_type=?", bug.SolveType)
	}
	if bug.SystemId != 0 {
		db = db.Where("system_id=?", bug.SystemId)
	}

	result := db.Order("id asc").Limit(bug.Limit).
		Offset(bug.Offset).Find(&b)
	//result := db.Table("bug").Where("handler_id=? AND status=? AND opportunity=? ", bug.PeoplePhone, bug.Status, bug.Opportunity).
	//	Where("AND solve_type=? AND priority_status=? AND type=?", bug.SolveType, bug.PriorityStatus, bug.Type).
	//	Where("begin_time>? AND begin_time<?", bug.BeginTime, bug.EndTime).
	//	Order("id asc").Limit(bug.Limit).
	//	Offset(bug.Offset).Find(b)
	return b, result.Error
}

func SelectBugListReporter(bug *req.BugList) ([]domain.Bug, error) {
	db := global.GetDb()
	var b []domain.Bug
	db = db.Table("bug")

	if bug.Type != 0  {
		db = db.Where("type=?", bug.Type)
	}
	if bug.BeginTime != 0 && bug.EndTime != 0 {
		db = db.Where("begin_time>? AND begin_time<?", bug.BeginTime, bug.EndTime)
	}
	if bug.Status != 0 {
		db = db.Where("status=?", bug.Status)
	}
	if bug.PeoplePhone != ""{
		db = db.Where("reporter_id=?", bug.PeoplePhone)
	}
	if bug.Opportunity != 0  {
		db = db.Where("opportunity=? ", bug.Opportunity)
	}
	if bug.PriorityStatus != 0  {
		db = db.Where("priority_status=?", bug.PriorityStatus)
	}
	if bug.SolveType != 0  {
		db = db.Where("solve_type=?", bug.SolveType)
	}
	if bug.SystemId != 0 {
		db = db.Where("system_id=?", bug.SystemId)
	}

	result := db.Order("id asc").Limit(bug.Limit).
		Offset(bug.Offset).Find(&b)
	//result := db.Table("bug").Where("reporter_id=? AND status=? AND opportunity=? ", bug.PeoplePhone, bug.Status, bug.Opportunity).
	//	Where("AND solve_type=? AND priority_status=? AND type=?", bug.SolveType, bug.PriorityStatus, bug.Type).
	//	Where("begin_time>? AND begin_time<?", bug.BeginTime, bug.EndTime).
	//	Order("id asc").Limit(bug.Limit).
	//	Offset(bug.Offset).Find(b)
	return b, result.Error
}

func SelectBugAll (bug *req.BugList) ([]domain.Bug, error){
	db := global.GetDb()
	var b []domain.Bug
	db = db.Table("bug")

	if bug.SystemId != 0 {
		db = db.Where("system_id=?", bug.SystemId)
	}
	if bug.Type != 0  {
		db = db.Where("type=?", bug.Type)
	}
	if bug.BeginTime != 0 && bug.EndTime != 0 {
		db = db.Where("begin_time>? AND begin_time<?", bug.BeginTime, bug.EndTime)
	}
	if bug.Status != 0 {
		db = db.Where("status=?", bug.Status)
	}
	if bug.Opportunity != 0  {
		db = db.Where("opportunity=? ", bug.Opportunity)
	}
	if bug.PriorityStatus != 0  {
		db = db.Where("priority_status=?", bug.PriorityStatus)
	}
	if bug.SolveType != 0  {
		db = db.Where("solve_type=?", bug.SolveType)
	}


	result := db.Order("id asc").Limit(bug.Limit).
		Offset(bug.Offset).Find(&b)
	return b, result.Error
}