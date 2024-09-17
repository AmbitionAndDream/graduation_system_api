package dev

import (
	"graduation_system_api/internal/global"
	"path/filepath"
	"runtime"
)

func CurrentFile() string {
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		panic("Can not get current file info")
	}
	return file
}

func GetDevConfigFile() string {
	return filepath.Join(CurrentFile(), "../../../configs/dev/config.toml")
}

func InitDevConf() {
	options := &global.Options{
		Log: true,
		MySQL:   true,
	}
	err := global.InitConfig(GetDevConfigFile(), options)
	if err != nil {
		return
	}
}
