package workerart

// Job structure.
type Job struct {
	list chan interface{}
}

// NewJob initialize job instance.
func NewJob() *Job {
	c := make(chan interface{}, 50)
	return &Job{list: c}
}

// AddJob add job into the channel list.
func (j *Job) AddJob(jobs ...interface{}) {
	for _, job := range jobs {
		j.list <- job
	}
}

// GetJobs get all jobs from channel list.
func (j *Job) GetJobs() <-chan interface{} {
	return j.list
}

// Close job channel list.
func (j *Job) Close() {
	close(j.list)
}
