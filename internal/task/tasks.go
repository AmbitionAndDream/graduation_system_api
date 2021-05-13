package task

import (
	"context"
	"sync"
)

type Tasks interface {
	AddTask(t Task)
	AddTaskFunc(fun TaskFun)
	Run(ctx context.Context)
	Wait() error
	Cancel()
}

type tasks struct {
	response   TaskResponse
	mutex      sync.Mutex
	tasks      []Task
	cancel     bool
	wg         sync.WaitGroup
	cancelFunc func()
}

func (t *tasks) AddTask(task Task) {
	t.tasks = append(t.tasks, task)
}
func (t *tasks) AddTaskFunc(fun TaskFun) {
	task := NewTask(fun)
	t.AddTask(task)
}
func (t *tasks) Run(ctx context.Context) {
	ctx1, cancelFunc := context.WithCancel(ctx)
	t.cancelFunc = cancelFunc
	t.response = make([]error, len(t.tasks))
	t.mutex = sync.Mutex{}
	t.wg = sync.WaitGroup{}
	t.wg.Add(len(t.tasks))
	for i, task := range t.tasks {
		go func(i int, task Task) {
			defer t.wg.Done()
			task.Run(ctx1)
			err := task.Wait()
			t.mutex.Lock()
			if err != nil && !t.cancel {
				t.response[i] = err
				t.cancel = true
				t.Cancel()
			}
			t.mutex.Unlock()
		}(i, task)
	}
}
func (t *tasks) Wait() error {
	t.wg.Wait()
	response := t.response
	return response.Error()
}
func (t *tasks) Cancel() {
	t.cancelFunc()
}

func NewTasks() Tasks {
	return &tasks{}
}
