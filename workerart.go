package workerart

import "sync"

// WorkerPool worker pool structure.
type WorkerPool struct {
	// Work to be done by workers in a worker pool.
	jobs Jobber

	// Wait for the workers to finish all the work.
	wait *sync.WaitGroup

	// Workers workersNumber.
	workersNumber uint

	// Result of the channel.
	result chan interface{}

	done chan struct{}
}

func (pool *WorkerPool) Done() <-chan struct{} {
	return pool.done
}

// NewWorkerPool initialize worker pool instance.
func NewWorkerPool(ops ...Option) *WorkerPool {
	pool := &WorkerPool{
		wait:          &sync.WaitGroup{},
		workersNumber: 100,
		result:        make(chan interface{}, 100),
	}

	for _, o := range ops {
		o.apply(pool)
	}

	return pool
}
