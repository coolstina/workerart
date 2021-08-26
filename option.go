package workerart

type Option interface {
	apply(*WorkerPool)
}

type optionFunc func(*WorkerPool)

func (o optionFunc) apply(ops *WorkerPool) {
	o(ops)
}

// WithWorkerNumber Use specific workers workersNumber override default value.
func WithWorkerNumber(workersNumber uint) Option {
	return optionFunc(func(ops *WorkerPool) {
		ops.workersNumber = workersNumber
	})
}

// WithWorkerNumber Use specific result channel override default value.
func WithResult(result chan interface{}) Option {
	return optionFunc(func(ops *WorkerPool) {
		ops.result = result
	})
}

// WithWorkerNumber Use specific result channel override default value.
func WithJobber(jobber Jobber) Option {
	return optionFunc(func(ops *WorkerPool) {
		ops.jobs = jobber
	})
}
