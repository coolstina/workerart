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
		for i := 0; i < no; i++ {
			pool.AddJobs(rand.Intn(no))
		}
		pool.CloseJob()
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