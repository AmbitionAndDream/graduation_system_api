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

func SelectBugListHandler(bug *req.BugList) (*domain.Bug, error) {
	db := global.GetDb()
	b := new(domain.Bug)
	result := db.Table("bug").Where("handler_id=? AND status=? AND opportunity=? ", bug.PeoplePhone, bug.Status, bug.Opportunity).
		Where("AND solve_type=? AND priority_status=? AND type=?", bug.SolveType, bug.PriorityStatus, bug.Type).
		Where("begin_time>? AND begin_time<?", bug.BeginTime, bug.EndTime).
		Order("id asc").Limit(bug.Limit).
		Offset(bug.Offset).Find(b)
	return b, result.Error
}

func SelectBugListReporter(bug *req.BugList) (*domain.Bug, error) {
	db := global.GetDb()
	b := new(domain.Bug)
	result := db.Table("bug").Where("reporter_id=? AND status=? AND opportunity=? ", bug.PeoplePhone, bug.Status, bug.Opportunity).
		Where("AND solve_type=? AND priority_status=? AND type=?", bug.SolveType, bug.PriorityStatus, bug.Type).
		Where("begin_time>? AND begin_time<?", bug.BeginTime, bug.EndTime).
		Order("id asc").Limit(bug.Limit).
		Offset(bug.Offset).Find(b)
	return b, result.Error
}
