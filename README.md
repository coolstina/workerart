![workerart](assets/banner/workerart.jpg)

Fast use worker pool process task.

## Install

```shell script
go get github.com/coolstina/workerart
```

## What is workerart?

Workerart is a quick implementation of using coroutine work pools, rather than repeating the wheel when you need to use workpools, to improve your development efficiency without losing the elegance of Go concurrent processing tasks. Workerart support:

- Build the working pool with options.
- Implement your own specific jobs through the Jobber interface.
- Custom task callback functions.

## Why use workerart?

While you can implement workpools yourself in order to gracefully handle multiple tasks, WorkerArt simply lets you use workpools more quickly to improve your task performance.

## How to use?


### [Quick try](./example/simple/main.go)

```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/coolstina/workerart"
)

func main() {
	// Step1: Initialize the worker pool.
	pool := workerart.NewWorkerPool()

	// Step2: Add a job task to the work pool.
	go func() {
		no := 10000
		for i := 1; i <= no; i++ {
			pool.AddJobStarting(rand.Intn(no))
		}
		pool.AddJobFinished()
	}()

	// Step3: Workers processing work.
	go pool.WorkersProcessing()

	// Step4: Receive worker pool processed result.
	go func() {
		once := 0
		for val := range pool.Results() {
			fmt.Println(val)
			once++
		}
		fmt.Printf("once: %+v\n", once)

		// Notice: Notifies the work pool that all work tasks are complete.
		pool.Done()
	}()

	// Step5: Wait worker pool that all work tasks are complete.
	select {
	case <-pool.Finished():
		fmt.Printf("done\n")
	case <-time.After(5 * time.Second):
		fmt.Printf("timeout\n")
	}
}
```

Note: In the above usage sample, the worker does nothing, and the workerart does not add complexity to the usage. By default, the task does nothing but return the added work. Look at the `taskCallback` value when `NewWorkerPool` is initialized:

```go
// NewWorkerPool initialize worker pool instance.
func NewWorkerPool(ops ...Option) *WorkerPool {
	pool := &WorkerPool{
		wait:          &sync.WaitGroup{},
		workersNumber: 100,
		result:        make(chan interface{}, 100),
		errors:        make(chan error, 100),
		done:          make(chan struct{}),

		// Default no processing,
		// this is，you must implement your own handlers.
		taskCallback: func(job interface{}) (interface{}, error) {
			return job, nil
		},

		jobs: NewJob(),
	}

	for _, o := range ops {
		o.apply(pool)
	}

	return pool
}
```

To make a task meaningful, you must implement a TaskCallback function that implements the logical processing that the task is specifically intended to perform. Take a look at the following example [Definition TaskCallback function](#[Definition TaskCallback function](./example/taskcallback/main.go)).

### [Definition TaskCallback function](./example/taskcallback/main.go)

```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/coolstina/workerart"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func main() {
	// Step1: Initialize the worker pool.
	pool := workerart.NewWorkerPool(
		workerart.WithTaskCallback(process()),
	)

	// Step2: Add a job task to the work pool.
	go func() {
		no := 100000
		for i := 1; i <= no; i++ {
			u := &User{
				Id:   i,
				Name: fmt.Sprintf("callme_%d", i),
				Age:  rand.Intn(40),
			}
			pool.AddJobStarting(u)
		}
		pool.AddJobFinished()
	}()

	// Step3: Workers processing work.
	go pool.WorkersProcessing()

	// Step4: Receive worker pool processed result.
	go func() {
		once := 0
		for val := range pool.Results() {
			fmt.Printf("receive: %+v\n", val)
			once++
		}
		fmt.Printf("once: %+v\n", once)

		// Notice: Notifies the work pool that all work tasks are complete.
		pool.Done()
	}()

	// Step5: Wait worker pool that all work tasks are complete.
	select {
	case <-pool.Finished():
		fmt.Printf("done\n")
	case <-time.After(5 * time.Second):
		fmt.Printf("timeout\n")
	}
}

func process() workerart.TaskCallback {
	return func(val interface{}) (interface{}, error) {
		user, ok := val.(*User)
		if ok {
			user.Age += 2
			return user, nil
		}
		return nil, nil
	}
}
```