// Copyright 2021 helloshaohua <wu.shaohua@foxmail.com>;
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package workerart

import (
	"sync"
)

// WorkerPool worker pool structure.
type WorkerPool struct {
	// Work to be done by workers in a worker pool.
	jobs Jobber

	// Wait for the workers to finish all the work.
	wait *sync.WaitGroup

	// Workers workersNumber.
	workersNumber uint

	// Default no processing,
	// this is，you must implement your own handlers.
	taskCallback TaskCallback

	// Result of the channel.
	result chan interface{}

	done chan struct{}

	errors chan error
}

// AddJobs Add the jobs to worker pool.
func (pool *WorkerPool) AddJobs(jobs ...interface{}) {
	pool.jobs.AddJob(jobs...)
}

// AddJobs Add the jobs to worker pool.
func (pool *WorkerPool) CloseJob() {
	pool.jobs.Close()
}

// WorkersProcessing Workers processing the task work.
func (pool *WorkerPool) WorkersProcessing() {

	// Create multiple workers to complete the work.
	for i := uint(0); i < pool.workersNumber; i++ {
		pool.wait.Add(1)

		// Worker task processing.
		go pool.worker()
	}

	// Wait for all worker processing to complete.
	pool.wait.Wait()

	// Close the result channel when all worker processing is complete.
	close(pool.result)
}

// Results The result of receiving worker processing.
func (pool *WorkerPool) Results() <-chan interface{} {
	return pool.result
}

// Done Whether the worker pool has completed all work.
func (pool *WorkerPool) Done() {
	pool.done <- struct{}{}
}


// Finished Whether the worker pool has completed all work.
func (pool *WorkerPool) Finished() <-chan struct{} {
	return pool.done
}

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

// worker task processing.
func (pool *WorkerPool) worker() {

	// Get job from channel.
	for job := range pool.jobs.GetJobs() {

		// Task processing.
		task, err := pool.taskCallback(job)

		// If there is an error, an error is sent to the channel.
		if err != nil {
			pool.errors <- err

			continue
		}

		// Send the processing results to the channel.
		pool.result <- task
	}

	// Current worker processing completed.
	pool.wait.Done()
}
