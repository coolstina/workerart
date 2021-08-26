package workerart

type Jobber interface {
	AddJob(val ...interface{})
	GetJobs() <-chan interface{}
	Close()
}
