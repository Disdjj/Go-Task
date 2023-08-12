package DJJPromise

import (
	"errors"
	"sync"
)

type Task[T any] struct {
	// return value
	value *T
	// error when occur
	error error

	// channel to wait
	ch chan struct{}

	// once
	once sync.Once
}

// NewTask create a new task instance.
func NewTask[T any](task func() (T, error)) *Task[T] {
	t := &Task[T]{
		ch:   make(chan struct{}),
		once: sync.Once{},
	}
	go t.taskWrapper(task)
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
		t.value = &value
		close(t.ch)
	})
}

func (t *Task[T]) taskWrapper(task func() (T, error)) {
	defer func() {
		if err := recover(); err != nil {
			t.errorSolve(errors.New("task panic"))
		}
	}()

	value, err := task()
	if err != nil {
		t.errorSolve(err)
		return
	}
	println("t pointer", &t)
	t.valueSolve(value)
}

// Await wait for the task to complete.
func (t *Task[T]) Await() (*T, error) {
	<-t.ch
	println("t pointer", &t)
	return t.value, t.error
}
