package auth

import (
	"strings"
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
	a:=map[int]string{}
	a[1]="3"
	a[1]="4"
	t.Log(a)
	b:=strings.Split("",",")
	t.Log(b)
	t.Log(len(b))
}
