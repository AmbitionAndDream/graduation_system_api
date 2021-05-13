package task

import "context"

type TaskFun func() error

type Task interface {
	Run(ctx context.Context)
	Wait() error
}
type task struct {
	ctx    context.Context
	err    error
	run    TaskFun
	finish chan struct{}
}

func (t *task) Run(ctx context.Context) {
	t.ctx = ctx
	go func() {
		t.err = t.run()
		t.finish <- struct{}{}
	}()
}
func (t *task) Wait() error {
	defer close(t.finish)
	select {
	case <-t.ctx.Done():
		return t.ctx.Err()
	case <-t.finish:
		return t.err
	}
}

func NewTask(f TaskFun) Task {
	return &task{
		run:    f,
		finish: make(chan struct{}),
	}
}
