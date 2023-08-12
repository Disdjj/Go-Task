package DJJPromise

import (
	"errors"
	"github.com/sourcegraph/conc"
	"sync"
)

type Task[T any] struct {
	// return value
	value T
	// error when occur
	error error

	// channel to wait
	ch chan struct{}

	// once
	once sync.Once
}

var taskPool = conc.WaitGroup{}

// NewTask create a new task instance.
func NewTask[T any](task func() (T, error)) *Task[T] {
	t := &Task[T]{
		ch:   make(chan struct{}),
		once: sync.Once{},
	}
	taskPool.Go(func() {
		t.taskWrapper(task)
	})
	return t
}

func (t *Task[T]) errorSolve(err error) {
	t.once.Do(func() {
		t.error = err
		close(t.ch)
	})
}

func (t *Task[T]) valueSolve(value T) {

	t.once.Do(func() {
		t.value = value
		close(t.ch)
	})
}

func (t *Task[T]) taskWrapper(task func() (T, error)) {
	defer func() {
		if err := recover(); err != nil {
			t.errorSolve(errors.New("Task Panic"))
		}
	}()

	value, err := task()
	if err != nil {
		t.errorSolve(err)
		return
	}
	t.valueSolve(value)
}

// Await wait for the task to complete.
func (t *Task[T]) Await() (T, error) {
	<-t.ch
	return t.value, t.error
}

func (t *Task[T]) ContinueWith(p func(t T) (interface{}, error)) *Task[interface{}] {
	return NewTask(func() (interface{}, error) {
		await, err := t.Await()
		if err != nil {
			return nil, err
		}
		return p(await)
	})
}
