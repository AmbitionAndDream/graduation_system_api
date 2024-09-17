package fusion

import (
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal/database"
	"graduation_system_api/internal/database/domain"
	resp "graduation_system_api/internal/domain"
	"graduation_system_api/internal/errors"
	"strconv"
)

func peopleDel(phone string) error {
	return database.PeopleDelDB(phone)
}

func createPeople(u *resp.RequestPeople) error {
	//构建dao
	people := &domain.User{
		Name:        u.Name,
		PhoneNumber: u.Phone,
		IsAdmin:     u.IsAdmin,
		RoleType:    u.RoleType,
		PassWord:    u.PassWord,
	}
	if err := database.CreatePeople(people); err != nil {
		logrus.Errorf("insert people param:%v,failed error :%s", people, err.Error())
		return errors.New(errors.ServerError, "insert failed")
	}
	return nil

}

func selectPeople(limit, offset string) (*resp.ResponsePeopleList, error) {
	l, _ := strconv.Atoi(limit)
	o, _ := strconv.Atoi(offset)
	//select
	result, err := database.SelectPeople(l, o)
	if err != nil {
		logrus.Errorf("select people list failed error :%s", err.Error())
		return nil, errors.New(errors.ServerError, "select failed")
	}
	logrus.Infof("the people list is %v", result)
	//数据组装
	var r []resp.ResponsePeople
	for _, element := range result {
		var entry resp.ResponsePeople
		entry.PeopleID = element.Id
		entry.Name = element.Name
		entry.PhoneNumber = element.PhoneNumber
		entry.IsAdmin = isAdmin(element.IsAdmin)
		entry.RoleType = element.RoleType
		r = append(r, entry)
	}
	//构造返回体
	res := &resp.ResponsePeopleList{
		Total:  len(r),
		User:   r,
		Limit:  l,
		Offset: o,
	}

	return res, nil
}

func isAdmin(flag int) bool {
	if flag == 1 {
		return true
	}
	return false
}


