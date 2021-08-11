package worker

import (
	"fmt"
	"sync"

	"github.com/fireoxcms/art/job"
	"github.com/fireoxcms/art/task"
)

type Worker struct {
	wait    *sync.WaitGroup
	Jobs    chan job.Job
	Results chan job.Result
	Task    *task.Task
	Done    chan struct{}
}

func NewWorker(jobs chan job.Job, results chan job.Result, task *task.Task) *Worker {
	return &Worker{
		wait:    &sync.WaitGroup{},
		Jobs:    jobs,
		Results: results,
		Task:    task,
		Done:    make(chan struct{}),
	}
}

func (worker *Worker) worker() {
	for j := range worker.Jobs {
		worker.Results <- job.Result{
			Job: j,
			Sum: worker.Task.Process(j.Random),
		}
	}
	worker.wait.Done()
}

func (worker *Worker) WorkerPool(numbers uint) {
	var i uint
	for i = 0; i < numbers; i++ {
		worker.wait.Add(1)
		go worker.worker()
	}
	worker.wait.Wait()
	close(worker.Results)
}

func (worker *Worker) Result() {
	for result := range worker.Results {
		fmt.Printf("Job id %d, input random no %d , sum of digits %d\n",
			result.Job.Id, result.Job.Random, result.Sum)
	}
	worker.Done <- struct{}{}
}
