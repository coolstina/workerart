package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fishfinal/workerart"
)

func main() {
	pool := workerart.NewWorkerPool()

	// Add jobs.
	go func() {
		no := 10000
		for i := 0; i < no; i++ {
			pool.AddJobs(rand.Intn(no))
		}
		pool.CloseJob()
	}()

	go pool.WorkersProcessing()

	go func() {
		once := 0
		for val := range pool.Results() {
			fmt.Println(val)
			once++
		}
		fmt.Printf("once: %+v\n", once)
		pool.Done()
	}()

	select {
	case <-pool.Finished():
		fmt.Printf("done\n")
	case <-time.After(5 * time.Second):
		fmt.Printf("timeout\n")
	}
}
