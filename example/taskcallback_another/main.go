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
	pool := workerart.NewWorkerPool()

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

	// Notes: Set the callback function separately
	pool.SetTaskCallback(process())

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
