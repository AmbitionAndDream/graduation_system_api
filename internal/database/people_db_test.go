package database

import (
	"graduation_system_api/internal/database/domain"
	"graduation_system_api/internal/dev"
	"graduation_system_api/internal/global"
	"testing"
)

func TestLogin(t *testing.T) {
	dev.InitDevConf()
	u, err := Login("1399927154", "88888888")
	t.Logf("%+v", u)
	t.Log(err)
	defer global.Close()
}

func TestSelectPeople(t *testing.T) {
	dev.InitDevConf()
	u, err := SelectPeople(2, 1)
	t.Logf("%+v", u)
	t.Log(err)
	defer global.Close()
}

func TestCreatePeople(t *testing.T) {
	dev.InitDevConf()
	err := CreatePeople(&domain.User{
		IsAdmin:     1,
		RoleType:    3,
		PassWord:    "12345678",
		Name:        "hello",
		PhoneNumber: "00000000",
	})
	if err != nil {
		t.Log(err)
		return
	}
}

func TestPeopleDelDB (t *testing.T) {
	dev.InitDevConf()
	err := PeopleDelDB("18329968200")
	if err != nil {
		t.Log(err)
		return
	}
}