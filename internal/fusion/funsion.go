package fusion

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal/auth"
	"graduation_system_api/internal/domain"
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/global"
	"net/http"
	"reflect"
	"strconv"
)

type newFusionHandler struct {
}

func GetFusionHandler() *newFusionHandler {
	return &newFusionHandler{}
}

func (f *newFusionHandler) HandlerLoginEvent(ctx *gin.Context) (resp interface{}, err error) {
	type User struct {
		PhoneNumber string `json:"phoneNumber" binding:"required"`
		PassWord    string `json:"passWord" binding:"required"`
	}
	u := new(User)
	if err = ctx.ShouldBind(u); err != nil {
		logrus.Errorf("parse user login params failed ,error: %s", err.Error())
		err = errors.New(http.StatusBadRequest, "url 参数有误")
		return
	}

	//登陆
	var role int
	if role, err = login(u.PhoneNumber, u.PassWord); err != nil {
		return nil, err
	}
	//生成token
	token := auth.GetToken(role, u.PhoneNumber)
	resp = struct {
		Token string `json:"token"`
	}{Token: token}
	return
}

func (f *newFusionHandler) HandleDemandEvent(ctx *gin.Context) (resp interface{}, err error) {
	action := ctx.Param("action")
	switch action {
	case global.Create:
		demand := new(domain.RequestDemand)
		if err = ctx.ShouldBind(demand); err != nil {
			logrus.Errorf("demand create param invalid ,the error is %s", err.Error())
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return nil, createDemand(demand)
	case global.Item:
		item := new(domain.RequestDemandItem)
		if err = ctx.ShouldBind(item); err != nil {
			logrus.Errorf("demand update item param invalid ,the error is %s", err.Error())
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return nil, updateItem(item)
	case global.SetTime:
		itemSetTime := new(domain.RequestDemandSetTime)
		if err = ctx.ShouldBind(itemSetTime); err != nil {
			logrus.Errorf("demand item setTime param invalid ,the error is %s", err.Error())
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return nil, updateItemTime(itemSetTime)
	case global.Solve:
		solve := new(domain.RequestDemandItem)
		if err = ctx.ShouldBind(solve); err != nil {
			logrus.Errorf("demand item solve param invalid ,the error is %s", err.Error())
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return nil, itemSolve(solve)
	case global.SelectDemandMyList:
		limit := ctx.Query("limit")
		offset := ctx.Query("offset")
		demandPhone := ctx.Query("phone")
		if !checkParam(limit, offset, demandPhone) {
			logrus.Errorf("demand my list param invalid ,the param limit: %s,offset:%s,demandPhone:%s", limit, offset, demandPhone)
			logrus.Errorf("HandleDemandEvent url param error ,url:%s", ctx.Request.URL.String())
			return nil, errors.New(errors.ParamInvalidError, "url 参数有误")
		}
		return selectDemandMyList(limit, offset, demandPhone)
	case global.Select:
		limit := ctx.Query("limit")
		offset := ctx.Query("offset")
		demandPhone := ctx.Query("phone")
		isAll := ctx.Query("is_all")
		status := ctx.Query("status")
		if !checkParam(limit, offset, demandPhone, isAll, status) {
			logrus.Errorf("demand list param invalid ,the param limit: %s,offset:%s,demandPhone:%s,isAll:%s,status:%s", limit, offset, demandPhone, isAll, status)
			logrus.Errorf("HandleDemandEvent url param error ,url:%s", ctx.Request.URL.String())
			return nil, errors.New(http.StatusBadRequest, "url 参数有误")
		}
		if (isAll != "true" && isAll != "false") || (status != "0" && status != "1") {
			logrus.Errorf("HandleDemandEvent url param error ,url:%s", ctx.Request.URL.String())
			return nil, errors.New(http.StatusBadRequest, "url 参数有误")
		}
		return selectDemandList(limit, offset, demandPhone, isAll, status)
	case global.Delete:
		ids := new(domain.RequestDeleteDemandId)
		if err = ctx.ShouldBind(ids); err != nil {
			logrus.Errorf("demand del param invalid ,the param is %v", ids)
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return nil, deleteDemand(ids.DemandId)
	case global.SelectPoolList:
		return selectPoolList()
	case global.Detail:
		demandId := ctx.Query("demand_id")
		if !checkParam(demandId) {
			logrus.Errorf("demand detail param invalid ,the param demand_id:%s", demandId)
			logrus.Errorf("HandleDemandEvent url param error ,url:%s", ctx.Request.URL.String())
			return nil, errors.New(http.StatusBadRequest, "url 参数有误")
		}
		did, _ := strconv.Atoi(demandId)
		return selectDemandById(did)
	default:
		logrus.Errorf("HandleDemandEvent url param error ,url:%s", ctx.Request.URL.String())
		return nil, errors.New(http.StatusBadRequest, "url 参数有误")
	}
}
func (f *newFusionHandler) HandleBugEvent(ctx *gin.Context) (resp interface{}, err error) {
	action := ctx.Param("action")
	switch action {
	case global.Create:
		bugCreate := new(domain.RequestBug)
		if err = ctx.ShouldBind(bugCreate); err != nil {
			logrus.Errorf("bug create param invalid ,the error is %s", err.Error())
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return nil, createBug(bugCreate)
	case global.Solve:
		bugSolve := new(domain.RequestBugSolve)
		if err = ctx.ShouldBind(bugSolve); err != nil {
			logrus.Errorf("bug solve param invalid ,the error is %s", err.Error())
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return nil, solveBug(bugSolve)
	case global.Select:
		//参数获取
		bug, err := getBugParam(ctx)
		if err != nil {
			return nil, err
		}
		return selectBugList(bug)
	default:
		logrus.Errorf("HandleBugEvent url param error ,url:%s", ctx.Request.URL.String())
		return nil, errors.New(http.StatusBadRequest, "url 参数有误")
	}
}
func getBugParam(ctx *gin.Context) (*domain.BugList, error) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	isAssign := ctx.Query("is_assign")
	status := ctx.Query("status")
	beginTime := ctx.Query("begin_time")
	endTime := ctx.Query("end_time")
	solveType := ctx.Query("solve_type")
	t := ctx.Query("type")
	opportunity := ctx.Query("opportunity")
	priorityStatus := ctx.Query("priority_status")
	phone := ctx.Query("user_id")
	if !checkParam(limit, offset, isAssign, status, solveType, beginTime, endTime, t, opportunity, priorityStatus, phone) {
		logrus.Errorf("HandleDemandEvent url param error ,url:%s", ctx.Request.URL.String())
		return nil, errors.New(http.StatusBadRequest, "url 参数有误")
	}
	l, _ := strconv.Atoi(limit)
	o, _ := strconv.Atoi(offset)
	i, _ := strconv.Atoi(isAssign)
	s, _ := strconv.Atoi(status)
	st, _ := strconv.Atoi(solveType)
	ty, _ := strconv.Atoi(t)
	op, _ := strconv.Atoi(opportunity)
	ps, _ := strconv.Atoi(priorityStatus)
	bt, _ := strconv.ParseInt(beginTime, 10, 64)
	et, _ := strconv.ParseInt(endTime, 10, 64)
	return &domain.BugList{
		Limit:          l,
		Offset:         o,
		Status:         s,
		BeginTime:      bt,
		SolveType:      st,
		Type:           ty,
		Opportunity:    op,
		IsAssign:       i,
		PriorityStatus: ps,
		EndTime:        et,
		PeoplePhone:    phone,
	}, nil
}

// 一个kind  对应一个Handle 多个action
func (f *newFusionHandler) HandleBusinessEvent(ctx *gin.Context) (resp interface{}, err error) {
	action := ctx.Param("action")
	switch action {
	case global.Create:
		name := &domain.RequestBusinessName{}
		if err := ctx.BindJSON(name); err != nil {
			logrus.Errorf("business create param invalid ,the error is %s", err.Error())
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return nil, createBusiness(name.Name)
	case global.Delete:
		ids := new(domain.RequestDeleteBusinessId)
		if err = ctx.ShouldBind(ids); err != nil {
			logrus.Errorf("business del param invalid ,the error is %s", err.Error())
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return nil, deleteBusiness(ids.BusinessId)
	case global.Select:
		limit := ctx.Query("limit")
		offset := ctx.Query("offset")
		if !checkParam(limit, offset) {
			logrus.Errorf("business list param invalid ,the param limit: %s,offset:%s", limit, offset)
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return selectBusiness(limit, offset)
	default:
		logrus.Errorf("HandleBusinessEvent url param error ,url:%s", ctx.Request.URL.String())
		return nil, errors.New(http.StatusBadRequest, "url 参数有误")
	}
}

func (f *newFusionHandler) HandlePeopleEvent(ctx *gin.Context) (resp interface{}, err error) {
	// del
	action := ctx.Param("action")
	switch action {
	case global.Select:
		limit := ctx.Query("limit")
		offset := ctx.Query("offset")
		if !checkParam(limit, offset) {
			logrus.Errorf("people list param invalid ,the param limit: %s,offset:%s", limit, offset)
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return selectPeople(limit, offset)
	case global.Create:
		u := new(domain.RequestPeople)
		if err = ctx.ShouldBind(u); err != nil {
			logrus.Errorf("parse people list params failed ,error: %s", err.Error())
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		return nil, createPeople(u)
	case global.Delete:
		phone := &domain.RequestPeoplePhone{}
		if err := ctx.BindJSON(phone); err != nil {
			logrus.Errorf("people del params invalid, the error is %s", err.Error())
			return nil, errors.New(errors.ParamInvalidError, "param invalid error")
		}
		if err = peopleDel(phone.Phone); err != nil {
			logrus.Errorf("people del error: %s", err.Error())
			return
		}
		return "people del successful", nil
	default:
		logrus.Errorf("HandlePeopleEvent url param error ,url:%s", ctx.Request.URL.String())
		return nil, errors.New(http.StatusBadRequest, "url 参数有误")
	}
}

func checkParam(date ...interface{}) bool {
	if len(date) == 0 {
		return false
	}
	k := reflect.ValueOf(date[0]).Kind()
	switch k {
	case reflect.String:
		for _, element := range date {
			if len(element.(string)) == 0 {
				return false
			}
		}
	}
	return true
}
