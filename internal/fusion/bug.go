package fusion

import (
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal/database"
	"graduation_system_api/internal/database/domain"
	req "graduation_system_api/internal/domain"
	resp "graduation_system_api/internal/domain"
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/global"
	"time"
)

func selectBugList(bug *req.BugList) (*resp.ResponseBugList, error) {
	bugDb := new(domain.Bug)
	bugDb.BeginTime = bug.BeginTime
	var b []domain.Bug
	var err error
	if bug.IsAssign == 3 {
		b, err = database.SelectBugAll(bug)
	}else if bug.IsAssign == global.ReporterId {
		b, err = database.SelectBugListHandler(bug)
	} else {
		b, err = database.SelectBugListReporter(bug)
	}
	if err != nil {
		logrus.Errorf("select bug failed param:%v,error:%s", bugDb, err.Error())
		return nil, errors.New(errors.ServerError, "select bug failed")
	}
	return buildBugResultList(b, int64(len(b)), bug.Limit, bug.Offset)
	//var bugList []resp.ResponseBug
	//
	//return &resp.ResponseBug{
	//	BugId:          b.BugId,
	//	SystemId:       b.SystemId,
	//	DemandId:       b.DemandId,
	//	BugName:        b.BugName,
	//	PriorityStatus: b.PriorityStatus,
	//	ReporterId:     b.ReporterId,
	//	HandlerId:      b.HandlerId,
	//	Type:           b.Type,
	//	Opportunity:    b.Opportunity,
	//	BeginTime:      b.BeginTime,
	//	SolveType:      b.SolveType,
	//}, nil
}
func buildBugResultList(result []domain.Bug, total int64, limit, offset int) (*resp.ResponseBugList, error) {
	var r []resp.ResponseBug
	for _, element := range result {
		demand, err := buildBugResult(&element)
		if err != nil {
			return nil, err
		}
		r = append(r, *demand)
	}
	//构造返回体
	res := &resp.ResponseBugList{
		Total:   total,
		BugList: r,
		Limit:   limit,
		Offset:  offset,
	}
	return res, nil
}

func buildBugResult(element *domain.Bug) (*resp.ResponseBug, error) {
	return &resp.ResponseBug{
		BugId:          element.BugId,
		SystemId:       element.SystemId,
		DemandId:       element.DemandId,
		BugName:        element.BugName,
		PriorityStatus: element.PriorityStatus,
		ReporterId:     element.ReporterId,
		HandlerId:      element.HandlerId,
		Type:           element.Type,
		Opportunity:    element.Opportunity,
		BeginTime:      element.BeginTime,
		SolveType:      element.SolveType,
		Status:		    element.Status,
	}, nil
}

func createBug(reqBug *req.RequestBug) error {
	bug := new(domain.Bug)
	bug.SystemId = reqBug.SystemId
	bug.DemandId = reqBug.DemandId
	bug.Opportunity = reqBug.Opportunity
	bug.ReporterId = reqBug.ReporterId
	bug.HandlerId = reqBug.HandlerId
	bug.Type = reqBug.Type
	bug.BeginTime = time.Now().UnixNano() / 1000000
	bug.BugName = reqBug.BugName
	bug.PriorityStatus = reqBug.PriorityStatus
	bug.Status = 1
	if err := database.CreateBug(bug); err != nil {
		logrus.Errorf("create bug :%v,failed error :%s", bug, err.Error())
		return errors.New(errors.ServerError, "create bug failed")
	}
	return nil
}

func solveBug(bug *req.RequestBugSolve) error {
	if err := database.SolveBug(bug.BugId, bug.Status, bug.SolveType); err != nil {
		logrus.Errorf("solve bug bug_id:%d,failed error :%s", bug.BugId, err.Error())
		return errors.New(errors.ServerError, "solve bug failed")
	}
	return nil
}
