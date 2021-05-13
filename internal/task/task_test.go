package task

import (
	"context"
	"github.com/pkg/errors"
	"testing"
	"time"
)

func TestTasks (t *testing.T){
	fun1 := func() error {
		time.Sleep(1 * time.Second)
		return nil
	}
	fun2 := func() error {
		time.Sleep(2 * time.Second)
		return nil
	}
	fun3 := func() error {
		time.Sleep(3 * time.Second)
		return nil
	}
	fun4 := func() error {
		time.Sleep(4 * time.Second)
		return nil
	}
	fun5 := func() error {
		time.Sleep(5 * time.Second)
		return errors.New("task error")
	}
	tasks := NewTasks()
	tasks.AddTaskFunc(fun1)
	tasks.AddTaskFunc(fun2)
	tasks.AddTaskFunc(fun3)
	tasks.AddTaskFunc(fun4)
	tasks.AddTaskFunc(fun5)
	ctx := context.Background()
	_, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()
	tasks.Run(ctx)
	if err :=tasks.Wait();err !=nil{
		t.Log(err.Error())

	}

}
