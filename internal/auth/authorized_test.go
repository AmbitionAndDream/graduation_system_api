package auth

import (
	"testing"
	"time"
)

func TestGetToken(t *testing.T) {
	token:=GetToken(2,"13999271543")
	t.Log(token)
	time.Sleep(6*time.Second)
	if err :=CheckToken(token);err !=nil{
		t.Log(err)
	}
}
func TestCheckToken(t *testing.T) {
	t.Log(time.Now().UnixNano()/1000000)
	t.Log(time.Now().Unix())
}