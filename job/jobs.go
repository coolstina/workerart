package job

import "math/rand"

type Job struct {
	Id     uint
	Random uint
}

type Result struct {
	Job Job
	Sum uint
}

type Jobs struct {
	Jobs   chan Job
	Result chan Result

	jobsBufferSize   uint
	resultBufferSize uint
}

func (jobs *Jobs) Allocate(number uint) {
	no := int(number)

	var i uint
	for i = 0; i < number; i++ {
		r := rand.Intn(no)

		jobs.Jobs <- Job{
			Id:     i,
			Random: uint(r),
		}
	}

	close(jobs.Jobs)
}

func NewJobs(jobsBufferSize, resultBufferSize uint) *Jobs {
	return &Jobs{
		Jobs:   make(chan Job, jobsBufferSize),
		Result: make(chan Result, jobsBufferSize),

		jobsBufferSize:   jobsBufferSize,
		resultBufferSize: resultBufferSize,
	}
}
