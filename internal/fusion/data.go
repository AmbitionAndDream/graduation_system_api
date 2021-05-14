package fusion

import (
	"context"
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal/database"
	req "graduation_system_api/internal/domain"
	"graduation_system_api/internal/errors"
	"graduation_system_api/internal/task"
	"time"
)
func selectBugNum (typ string) (map[int]int, error) {
	currentTime := time.Now()
	m, _ := time.ParseDuration("-168h")
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	beginTime := endTime.Add(m).UnixNano() / 1000000

	params := req.BugList{ Limit: 0, Offset: 0, BeginTime: beginTime, EndTime: endTime.UnixNano() / 1000000 }
	b, err := database.SelectBugAll(&params)
	if err != nil {
		logrus.Errorf("select bug failed param:%v,error:%s", b, err.Error())
		return nil, errors.New(errors.ServerError, "select bug failed")
	}
	countArr := make(map[int]int)
	for _, element := range b {
		var countType int
		if typ == "1" {
			countType = element.Type
		}else{
			countType = element.SolveType
		}
		if countType != 0{
			if _, ok := countArr[countType]; !ok{
				countArr[countType] = 1
			}else {
				count := countArr[countType]
				countArr[countType] = count + 1
			}
		}

	}
	return countArr, nil
}

func selectDemandNum () (map[int64]int64, error) {
	currentTime := time.Now()
	T1, _ := time.ParseDuration("-24h")
	T2, _ := time.ParseDuration("-48h")
	T3, _ := time.ParseDuration("-72h")
	T4, _ := time.ParseDuration("-96h")
	T5, _ := time.ParseDuration("-120h")
	T6, _ := time.ParseDuration("-144h")
	T7, _ := time.ParseDuration("-168h")
	end := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())

	res := make(map[int64]int64)
	time1 := end.Add(T1).UnixNano() / 1000000
	time2 := end.Add(T2).UnixNano() / 1000000
	time3 := end.Add(T3).UnixNano() / 1000000
	time4 := end.Add(T4).UnixNano() / 1000000
	time5 := end.Add(T5).UnixNano() / 1000000
	time6 := end.Add(T6).UnixNano() / 1000000
	time7 := end.Add(T7).UnixNano() / 1000000

	fun1 := func() error {
		count, err := database.SelectDemandByTime(time1, end.UnixNano()/1000000)
		if err != nil {
			logrus.Errorf("select demand count failed param:%v,error:%s", count, err.Error())
			return  errors.New(errors.ServerError, "select demand failed")
		}
		res[time1] = count
		return nil
	}
	fun2 := func() error {
		count, err := database.SelectDemandByTime(time2, time1)
		if err != nil {
			logrus.Errorf("select demand count failed param:%v,error:%s", count, err.Error())
			return  errors.New(errors.ServerError, "select demand failed")
		}
		res[time2] = count
		return nil
	}
	fun3 := func() error {
		count, err := database.SelectDemandByTime(time3, time2)
		if err != nil {
			logrus.Errorf("select demand count failed param:%v,error:%s", count, err.Error())
			return  errors.New(errors.ServerError, "select demand failed")
		}
		res[time3] = count
		return nil
	}
	fun4 := func() error {
		count, err := database.SelectDemandByTime(time4, time3)
		if err != nil {
			logrus.Errorf("select demand count failed param:%v,error:%s", count, err.Error())
			return  errors.New(errors.ServerError, "select demand failed")
		}
		res[time4] = count
		return nil
	}
	fun5 := func() error {
		count, err := database.SelectDemandByTime(time5, time4)
		if err != nil {
			logrus.Errorf("select demand count failed param:%v,error:%s", count, err.Error())
			return  errors.New(errors.ServerError, "select demand failed")
		}
		res[time5] = count
		return nil
	}
	fun6 := func() error {
		count, err := database.SelectDemandByTime(time6, time5)
		if err != nil {
			logrus.Errorf("select demand count failed param:%v,error:%s", count, err.Error())
			return  errors.New(errors.ServerError, "select demand failed")
		}
		res[time6] = count
		return nil
	}
	fun7 := func() error {
		count, err := database.SelectDemandByTime(time7, time6)
		if err != nil {
			logrus.Errorf("select demand count failed param:%v,error:%s", count, err.Error())
			return  errors.New(errors.ServerError, "select demand failed")
		}
		res[time7] = count
		return nil
	}

	tasks := task.NewTasks()
	tasks.AddTaskFunc(fun1)
	tasks.AddTaskFunc(fun2)
	tasks.AddTaskFunc(fun3)
	tasks.AddTaskFunc(fun4)
	tasks.AddTaskFunc(fun5)
	tasks.AddTaskFunc(fun6)
	tasks.AddTaskFunc(fun7)

	ctx := context.Background()
	_, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()
	tasks.Run(ctx)
	if err :=tasks.Wait();err !=nil{
		logrus.Infof("demand num error: %s", err.Error())
		return nil, errors.New(errors.ServerError, "select demand failed")
	}
	return res, nil
}