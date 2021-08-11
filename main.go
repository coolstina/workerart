package main

import (
	"fmt"
	"time"

	"github.com/fireoxcms/art/job"
	"github.com/fireoxcms/art/task"
	"github.com/fireoxcms/art/worker"
)

func main() {
	start := time.Now()

	jobs := job.NewJobs(10, 10)
	go jobs.Allocate(100)

	wrk := worker.NewWorker(jobs.Jobs, jobs.Result, task.NewTask())
	go wrk.WorkerPool(50)
	go wrk.Result()

	<-wrk.Done

	finished := time.Now().Sub(start)

	fmt.Println("use time: ", finished)
}
