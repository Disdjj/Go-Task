package DJJPromise

import "github.com/sourcegraph/conc"

// learned from github.com/chebyrash/promise
type GoPool interface {
	Go(f func())
}

type wrap func(f func())

func (wf wrap) Go(f func()) {
	wf(f)
}

func NewGoPool() GoPool {
	return wrap(func(f func()) {
		go f()
	})
}

var c = conc.NewWaitGroup()

func NewConcGoPool() GoPool {
	return wrap(func(f func()) {
		c.Go(f)
	})
}
