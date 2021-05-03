package fusion

import (
	"graduation_system_api/internal/dev"
	"testing"
)

func TestNewFusionHandler_HandlePeopleEvent(t *testing.T) {
	dev.InitDevConf()
	resp,err:=selectPeople("2","1")
	if err!=nil{
		t.Log(err)
		return
	}
	t.Log(resp)
}
