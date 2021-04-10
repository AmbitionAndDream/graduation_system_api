package errors

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type Error struct {
	code  int
	error error
}

func (e *Error) Error() string {
	return e.error.Error()
}

func New(code int, msg string) error {
	return &Error{
		error: errors.New(msg),
		code:  code,
	}
}

func Wrap(code int, err error, msg string) error {
	return &Error{
		error: errors.Wrap(err, msg),
		code:  code,
	}
}

func (e *Error) Code() int {
	return e.code
}
func (e *Error) ToString() string {
	b, _ := json.Marshal(map[string]interface{}{"code:": e.code, "error msg:": e.Error()})
	return string(b)
}


