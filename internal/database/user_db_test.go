package database

import (
	"graduation_system_api/internal/dev"
	"graduation_system_api/internal/global"
	"testing"
)

func TestLogin(t *testing.T) {
	dev.InitDevConf()
	u,err:=Login("1399927154","88888888")
	t.Logf("%+v",u)
	t.Log(err)
	defer global.Close()
}