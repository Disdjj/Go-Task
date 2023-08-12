# Go-Task

programing Go in Async way.

## Usage

```go
package main

import (
	task "github.com/Disdjj/Go-Task"
)

func main() {
	p := task.NewTask[string](func() (string, error) {
		return "123", nil
	})
	res, err := p.Await()
	if err != nil {
		panic(err)
	}
	println(res)
}

```

## Attention

This Project is use for learning purpose, not for production.
