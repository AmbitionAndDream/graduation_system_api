package auth

import (
	"encoding/base64"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"graduation_system_api/internal/errors"
	"reflect"
	"time"
)

func CheckToken(t string) error {
	date,err:=base64.StdEncoding.DecodeString(t)
	if err !=nil{
		return err
	}
	token := make(map[string]interface{})
	if err =json.Unmarshal(date,&token);err !=nil{
		return err
	}
	logrus.Infof("the token decode is :%v",token)

	if !hastRight(token){
		return errors.New(700,"token expired error")
	}

	return nil
}
func hastRight(token map[string]interface{})bool{
	expired,ok:=token["expired"]
	logrus.Info(reflect.ValueOf(expired).Kind())
	if reflect.ValueOf(expired).Kind()==reflect.Float64{
		return ok && expired.(float64)>float64(time.Now().Unix())
	}
	return false
}
func GetToken(role int, phoneNumber string) string {
	logrus.Infof("get token role:%d,phoneNumebr:%s",role,phoneNumber)
	token := make(map[string]interface{})
	token["role"]=role
	token["phoneNumber"]=phoneNumber
	token["expired"]=time.Now().Add(time.Second*5).Unix()
	b,_:=json.Marshal(token)
	logrus.Infof("the token%v, encode is :%s",token,base64.StdEncoding.EncodeToString(b))
	return base64.StdEncoding.EncodeToString(b)

}
