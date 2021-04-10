package errors

import "testing"

func TestName(t *testing.T) {
	err:= New(UserUndefinedError,"未定义")
	er:= Wrap(ServerError,err,"failed")
	if e,ok:=er.(*Error);ok{
		t.Log(e.ToString())
	}
}
