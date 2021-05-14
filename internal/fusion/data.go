package fusion

import (
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal/database"
	req "graduation_system_api/internal/domain"
	"graduation_system_api/internal/errors"
	"time"
)
func selectBugNum (typ string) (map[int]int, error) {
	currentTime := time.Now()
	m, _ := time.ParseDuration("-168h")
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())
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