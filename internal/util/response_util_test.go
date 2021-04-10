package util

import (
	"encoding/json"
	"testing"
)

func TestName(t *testing.T) {
	res,_ := json.Marshal(struct {
		Age int `json:"age,omitempty"`
	}{})
	t.Log(string(res))
}
