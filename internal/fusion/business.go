package fusion

import (
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal/database"
	resp "graduation_system_api/internal/domain"
	"graduation_system_api/internal/errors"
	"strconv"
)

func createBusiness(name string) error {
	if err := database.CreateBusiness(name); err != nil {
		logrus.Errorf("insert business name:%s,failed error :%s", name, err.Error())
		return errors.New(errors.ServerError, "insert failed")
	}
	return nil
}

func selectBusiness(limit, offset string) (*resp.ResponseBusinessList, error) {
	l, _ := strconv.Atoi(limit)
	o, _ := strconv.Atoi(offset)
	//select
	result, err := database.SelectBusiness(l, o)
	if err != nil {
		logrus.Errorf("select business list failed error :%s", err.Error())
		return nil, errors.New(errors.ServerError, "select failed")
	}
	logrus.Infof("the business list is %v", result)
	//数据组装
	var r []resp.ResponseBusiness
	for _, element := range result {
		var entry resp.ResponseBusiness
		entry.BusinessID = element.BusinessID
		entry.BusinessName = element.BusinessName
		r = append(r, entry)
	}
	//构造返回体
	res := &resp.ResponseBusinessList{
		Total:  len(r),
		Bus:    r,
		Limit:  l,
		Offset: o,
	}

	return res, nil
}

func deleteBusiness(ids []int) error {
	//id := strings.Split(idStr, ",")
	//var ids []int
	//for _, element := range id {
	//	entry, _ := strconv.Atoi(element)
	//	ids = append(ids, entry)
	//}
	return database.DeleteBusiness(ids)
}
