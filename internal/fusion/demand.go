package fusion

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal/database"
	"graduation_system_api/internal/database/domain"
	req "graduation_system_api/internal/domain"
	resp "graduation_system_api/internal/domain"
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/global"
	"strconv"
)

func deleteDemand(ids []int) error {
	return database.DeleteDemand(ids)
}
func selectPoolList() (*resp.ResponseDemandPoolList, error) {
	demandLists, err := database.SelectAllDemandList()
	if err != nil {
		logrus.Errorf("select demand list failed error :%s", err.Error())
		return nil, errors.New(errors.ServerError, "select demand pool list by id failed")
	}
	logrus.Infof("select demand list is :%v", demandLists)
	//数据转换(只为拿到转换后的pojo)
	result, err := buildDemandResultList(demandLists, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	//reviewPool
	reviewPool := new(resp.ResponseReviewPool)
	var reviewPoolDemandList []resp.ResponseDemand
	for _, element := range result.DemandList {
		var itemInfo []resp.ResponseDemandNodeInfo
		for _, item := range element.DemandNodeInfo {
			if item.Status == global.NotFinish && item.ItemType == global.ReviewPoll {
				itemInfo = append(itemInfo, item)
				break
			}
		}
		//赋值
		element.DemandNodeInfo = itemInfo
		reviewPoolDemandList = append(reviewPoolDemandList, element)
	}
	reviewPool.DemandList = reviewPoolDemandList
	reviewPool.Total = int64(len(reviewPoolDemandList))
	//developmentPoll
	developmentPoll := new(resp.ResponseDevelopmentPoll)
	var developmentPollDemandList []resp.ResponseDemand
	for _, element := range result.DemandList {
		var itemInfo []resp.ResponseDemandNodeInfo
		for _, item := range element.DemandNodeInfo {
			if item.Status == global.NotFinish && item.ItemType == global.DevelopmentPoll {
				itemInfo = append(itemInfo, item)
				break
			}
		}
		//赋值
		element.DemandNodeInfo = itemInfo
		developmentPollDemandList = append(developmentPollDemandList, element)
	}
	developmentPoll.DemandList = developmentPollDemandList
	developmentPoll.Total = int64(len(developmentPollDemandList))
	//testPoll
	testPoll := new(resp.ResponseTestPoll)
	var testPollDemandList []resp.ResponseDemand
	for _, element := range result.DemandList {
		var itemInfo []resp.ResponseDemandNodeInfo
		for _, item := range element.DemandNodeInfo {
			if item.Status == global.NotFinish && item.ItemType == global.TestPoll {
				itemInfo = append(itemInfo, item)
				break
			}
		}
		//赋值
		element.DemandNodeInfo = itemInfo
		testPollDemandList = append(testPollDemandList, element)
	}
	testPoll.DemandList = testPollDemandList
	testPoll.Total = int64(len(testPollDemandList))
	//acceptancePoll
	acceptancePoll := new(resp.ResponseAcceptancePoll)
	var acceptancePollDemandList []resp.ResponseDemand
	for _, element := range result.DemandList {
		var itemInfo []resp.ResponseDemandNodeInfo
		for _, item := range element.DemandNodeInfo {
			if item.Status == global.NotFinish && item.ItemType == global.AcceptancePoll {
				itemInfo = append(itemInfo, item)
				break
			}
		}
		//赋值
		element.DemandNodeInfo = itemInfo
		acceptancePollDemandList = append(acceptancePollDemandList, element)
	}
	acceptancePoll.DemandList = acceptancePollDemandList
	acceptancePoll.Total = int64(len(acceptancePollDemandList))
	//completePoll
	completePoll := new(resp.ResponseCompletePoll)
	var completePollDemandList []resp.ResponseDemand
	for _, element := range result.DemandList {
		var itemInfo []resp.ResponseDemandNodeInfo
		for _, item := range element.DemandNodeInfo {
			if item.Status == global.Finish && item.ItemType == global.CompletePoll {
				itemInfo = append(itemInfo, item)
				break
			}
		}
		//赋值
		element.DemandNodeInfo = itemInfo
		completePollDemandList = append(completePollDemandList, element)
	}
	completePoll.DemandList = completePollDemandList
	completePoll.Total = int64(len(completePollDemandList))
	return &resp.ResponseDemandPoolList{
		ReviewPool:      reviewPool,
		DevelopmentPoll: developmentPoll,
		TestPoll:        testPoll,
		AcceptancePoll:  acceptancePoll,
		CompletePoll:    completePoll,
	}, nil
}

func selectDemandById(demandId int) (*resp.ResponseDemand, error) {
	//查出对应的demand
	demand, err := database.SelectDemandById(demandId)
	if err != nil {
		logrus.Errorf("select demand by demand_id:%d failed error :%s", demandId, err.Error())
		return nil, errors.New(errors.ServerError, "select demand by id failed")
	}
	logrus.Infof("select demand by demand_id:%d,is :%v", demandId, demand)
	return buildDemandResult(demand)
}

func updateItem(item *req.RequestDemandItem) error {
	result, err := selectDemandById(item.DemandId)
	if err != nil {
		return err
	}
	//遍历修改数据
	for _, element := range result.DemandNodeInfo {
		if element.ItemId == item.ItemId {
			element.PeoplePhone = item.NodePeoplePhone
		}
	}
	//update
	return UpdateByDemandIdForItemInfo(item.DemandId, result.DemandNodeInfo)
}
func UpdateByDemandIdForItemInfo(demandId int, itemInfo []resp.ResponseDemandNodeInfo) error {
	//struct -> json
	info, err := json.Marshal(itemInfo)
	if err != nil {
		logrus.Errorf("demand item info json marshal error:%s", err.Error())
		return errors.New(errors.ServerError, err.Error())
	}
	logrus.Infof("update demand item info:%s by deman_id:%d", string(info), demandId)

	err = database.UpdateDemandItemById(demandId, string(info))
	if err != nil {
		logrus.Errorf("update demand item info failed error :%s", err.Error())
		return errors.New(errors.ServerError, "update item info failed")
	}
	return nil
}
func updateItemTime(item *req.RequestDemandSetTime) error {
	result, err := selectDemandById(item.DemandId)
	if err != nil {
		return err
	}
	//遍历修改数据
	for _, element := range result.DemandNodeInfo {
		if element.ItemId == item.ItemId {
			element.Time = item.Time
		}
	}
	//update
	return UpdateByDemandIdForItemInfo(item.DemandId, result.DemandNodeInfo)
}

func itemSolve(item *req.RequestDemandItem) error {
	result, err := selectDemandById(item.DemandId)
	if err != nil {
		return err
	}
	//遍历修改数据
	for _, element := range result.DemandNodeInfo {
		if element.ItemId == item.ItemId {
			element.Status = global.Finish
		}
	}
	//update
	return UpdateByDemandIdForItemInfo(item.DemandId, result.DemandNodeInfo)
}

func selectDemandList(limit, offset, demandPhone, isAll, status string) (*resp.ResponseDemandList, error) {
	l, _ := strconv.Atoi(limit)
	o, _ := strconv.Atoi(offset)
	//查出来全部的数据
	result, err := database.SelectDemandList(l, o)
	if err != nil {
		logrus.Errorf("select demand list failed error :%s", err.Error())
		return nil, errors.New(errors.ServerError, "select demand list failed")
	}
	logrus.Infof("the demand my list is %v", result)
	//查totalCount
	totalCount, err := database.SelectAllDemandListCount()
	if err != nil {
		logrus.Errorf("select demand list totalCount failed error :%s", err.Error())
		return nil, errors.New(errors.ServerError, "select demand total failed")
	}
	logrus.Infof("the demand my list totalCount is %d", totalCount)
	response, err := buildDemandResultList(result, totalCount, l, o)
	if isAll == "true" { //所有的需求，直接构造返回体
		return response, err
	} else { //分配给我的需求
		sta, _ := strconv.ParseInt(status, 10, 64)
		return selectForMyDemand(response, l, o, demandPhone, sta)
	}
}

func selectForMyDemand(result *resp.ResponseDemandList, limit, offset int, demandPhone string, status int64) (*resp.ResponseDemandList, error) {
	//筛选出跟我相关的需求
	//构建需求slice
	var demandForAllMyList []resp.ResponseDemand
	for _, demandList := range result.DemandList {
		//构建分配给我的需求节点info slice
		var nodeForMyInfo []resp.ResponseDemandNodeInfo
		for _, node := range demandList.DemandNodeInfo {
			//将分配给我的需求，放入构建的需求节点info slice
			if node.PeoplePhone == demandPhone && node.Status == status {
				nodeForMyInfo = append(nodeForMyInfo, node)
			}
		}
		demandForAllMyList = append(demandForAllMyList, demandList)
	}
	total := int64(len(demandForAllMyList))

	//进行手动分页
	var demandList []resp.ResponseDemand

	if offset < len(demandForAllMyList) {
		var begin, end int
		begin = offset
		if end = offset + limit; end > len(demandForAllMyList) {
			end = len(demandForAllMyList)
		}
		demandList = demandForAllMyList[begin:end]
	}

	return &resp.ResponseDemandList{
		Total:      total,
		Limit:      limit,
		Offset:     offset,
		DemandList: demandList,
	}, nil

}

func selectDemandMyList(limit, offset, demandPhone string) (*resp.ResponseDemandList, error) {
	l, _ := strconv.Atoi(limit)
	o, _ := strconv.Atoi(offset)
	//select
	result, err := database.SelectDemandMyList(l, o, demandPhone)
	if err != nil {
		logrus.Errorf("select demand my list failed error :%s", err.Error())
		return nil, errors.New(errors.ServerError, "select failed")
	}
	logrus.Infof("the demand my list is %v", result)
	//查totalCount
	totalCount, err := database.SelectAllDemandMyList(demandPhone)
	if err != nil {
		logrus.Errorf("select demand my list totalCount failed error :%s", err.Error())
		return nil, errors.New(errors.ServerError, "select failed")
	}
	logrus.Infof("the demand my list totalCount is %d", totalCount)
	//数据组装
	return buildDemandResultList(result, totalCount, l, o)
}

func createDemand(demand *req.RequestDemand) error {
	b, err := json.Marshal(demand.DemandNodeInfo)
	if err != nil {
		logrus.Errorf("demand node info json marshal error:%s", err.Error())
		return errors.New(errors.ServerError, err.Error())
	}
	dbDemand := &domain.Demand{
		DemandName:           demand.DemandName,
		DemandLink:           demand.DemandLink,
		DemandNote:           demand.DemandNote,
		DemandPriorityStatus: demand.DemandPriorityStatus,
		DemandInfo:           string(b),
		BusinessId:           demand.BusinessId,
		PeoplePhone:          demand.PeoplePhone,
	}
	if err := database.CreateDemand(dbDemand); err != nil {
		logrus.Errorf("insert demand param:%v,failed error :%s", demand, err.Error())
		return errors.New(errors.ServerError, "create failed")
	}
	return nil
}

func buildDemandResultList(result []domain.Demand, total int64, limit, offset int) (*resp.ResponseDemandList, error) {
	var r []resp.ResponseDemand
	for _, element := range result {
		demand, err := buildDemandResult(&element)
		if err != nil {
			return nil, err
		}
		r = append(r, *demand)
	}
	//构造返回体
	res := &resp.ResponseDemandList{
		Total:      total,
		DemandList: r,
		Limit:      limit,
		Offset:     offset,
	}
	return res, nil
}

func buildDemandResult(element *domain.Demand) (*resp.ResponseDemand, error) {
	r := new(resp.ResponseDemand)
	var entry resp.ResponseDemand
	var i []resp.ResponseDemandNodeInfo
	entry.DemandName = element.DemandName
	entry.DemandLink = element.DemandLink
	entry.PeoplePhone = element.PeoplePhone
	entry.DemandPriorityStatus = element.DemandPriorityStatus
	entry.DemandNote = element.DemandNote
	entry.BusinessId = element.BusinessId
	entry.DemandId = element.DemandId
	if err := json.Unmarshal([]byte(element.DemandInfo), &i); err != nil {
		logrus.Errorf("demand node info json Unmarshal error:%s", err.Error())
		return nil, errors.New(errors.ServerError, err.Error())
	}
	entry.DemandNodeInfo = i
	logrus.Infof("the demand my list entry is %v", entry)

	return r, nil
}
