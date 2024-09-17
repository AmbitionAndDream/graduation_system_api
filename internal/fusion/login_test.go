package fusion

import (
	"graduation_system_api/internal/dev"
	"graduation_system_api/internal/global"
	"testing"
)

func TestGetFusionHandler(t *testing.T) {
	dev.InitDevConf()
	role,err:=login("1399927154","88888888")
	t.Logf("%d",role)
	t.Log(err)
	defer global.Close()
}
