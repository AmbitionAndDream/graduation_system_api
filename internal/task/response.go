package task

import "fmt"

type TaskResponse []error

func (r TaskResponse) Success() bool {
	err := r.Error()
	return err == nil
}
func (r TaskResponse) Error() error {
	for i, err := range r {
		if err != nil {
			return fmt.Errorf("task:%d error :%s", i+1, err)
		}
	}
	return nil
}
