package database

import (
	"graduation_system_api/internal/dev"
	"testing"
)

func TestCreateBusiness(t *testing.T) {
	dev.InitDevConf()
	if err := CreateBusiness("hello"); err != nil {
		t.Log(err)
	}
}
func TestSelectBusiness(t *testing.T) {
	dev.InitDevConf()
	resp, err := SelectBusiness(1,1)
	if err != nil {
		t.Log(err)
	}
	t.Log(resp)
}
func TestDeleteBusiness(t *testing.T) {
	dev.InitDevConf()
	if err := DeleteBusiness([]int{1,2}); err != nil {
		t.Log(err)
	}
}