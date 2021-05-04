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
	r:=3
	switch r {
	case 1:
		t.Log("heihei")
	case 2:
		t.Log("hahah")
	case 3:
		t.Log("lalal")
	default:
		t.Log("huhuhu")
	}
}

func TestCheckToken2(t *testing.T) {
	sli:=[]int{1,2,3,4,5,6,7,8,9,0,11}
	t.Log(sli[1:])
	t.Log(sli[0:])
	t.Log(sli[9:10])
	t.Log(sli[1:5+1])
	t.Log(sli[len(sli):])
}
