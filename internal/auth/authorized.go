package auth

import (
	"encoding/base64"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal/errors"
	"sync"
	"time"
)

var expired map[string]int64
var lock sync.RWMutex

func init() {
	expired = make(map[string]int64, 100)
}

func setExpired(token string) {
	lock.Lock()
	expired[token] = time.Now().Add(time.Hour * 5).Unix()
	lock.Unlock()
}

func getExpired(token string) (int64, bool) {
	lock.RLock()
	t, exits := expired[token]
	lock.RUnlock()
	if exits {
		return t, true
	} else {
		return 0, false
	}
}

func CheckToken(t string) error {
	expired, exit := getExpired(t)
	if !exit {
		return errors.New(errors.TokenInvalidError, "token invalid error")
	}
	date, err := base64.StdEncoding.DecodeString(t)
	if err != nil {
		return err
	}
	token := make(map[string]interface{})
	if err = json.Unmarshal(date, &token); err != nil {
		return err
	}
	logrus.Infof("the token decode is :%v", token)

	if expired > time.Now().Unix() {
		setExpired(t)
	} else {

		return errors.New(errors.TokenExpiredError, "token expired error")
	}

	return nil
}

//func hastRight(t string) bool {
//	expired, exit := getExpired(t)
//	return exit &&
//	//expired,ok:=token["expired"]
//	//logrus.Info(reflect.ValueOf(expired).Kind())
//	//if reflect.ValueOf(expired).Kind()==reflect.Float64{
//	//	return ok && expired.(float64)>float64(time.Now().Unix())
//	//}
//	//return false
//}

func GetToken(role int, phoneNumber string) string {
	logrus.Infof("get token role:%d,phoneNumebr:%s", role, phoneNumber)
	tokenMap := make(map[string]interface{})
	tokenMap["role"] = role
	tokenMap["phoneNumber"] = phoneNumber
	//token["expired"]=time.Now().Add(time.Second*5).Unix()
	b, _ := json.Marshal(tokenMap)
	logrus.Infof("the token%v, encode is :%s", tokenMap, base64.StdEncoding.EncodeToString(b))
	token := base64.StdEncoding.EncodeToString(b)
	//设置过期时间
	setExpired(token)
	return token
}
