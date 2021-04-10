package auth

import (
	"errors"
)

func CheckToken() error {
	return errors.New("token timeout")
}
func GetToken(role int, phoneNumber string) string {
	return phoneNumber
}
