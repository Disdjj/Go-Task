package DJJPromise

import (
	"errors"
	"testing"
	"time"
)

func TestBasicRun(t *testing.T) {
	p := NewTask[string](func() (string, error) {
		return "123", nil
	})

	value, err := p.Await()
	if err != nil {
		t.Error(err)
	}

	if value != "123" {
		//print(value)
		t.Error("value is not 123")
	}
}

func TestBasicError(t *testing.T) {
	p := NewTask[string](func() (string, error) {
		return "", errors.New("return")
	})

	_, err := p.Await()
	if err == nil {
		t.Error("")
	}
}

func TestPanic(t *testing.T) {
	p := NewTask[string](func() (string, error) {
		panic(errors.New("return"))
		//return "", errors.New("return")
	})

	_, err := p.Await()
	if err == nil {
		t.Error("")
	}
}

func TestTask_ContinueWith(t *testing.T) {
	preTask := NewTask[string](func() (string, error) {
		time.Sleep(1 * time.Second)
		return "hello world", nil
	})

	task := preTask.ContinueWith(func(t string) (interface{}, error) {
		time.Sleep(1 * time.Second)
		return t + "!", nil
	})

	v, err := task.Await()
	if err != nil {
		t.Error("err in run task")
	}
	if v != "hello world!" {
		t.Error("err result")
	}
}
