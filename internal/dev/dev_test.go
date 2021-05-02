package dev

import (
	"graduation_system_api/internal/global"
	"testing"
)

func TestGetDevConfigFile(t *testing.T) {
	t.Log(CurrentFile())
}
func TestInitDevConf(t *testing.T) {
	InitDevConf()
	global.Close()
}