package workerart

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestWorkerPoolSuite(t *testing.T) {
	suite.Run(t, new(WorkerPoolSuite))
}

type WorkerPoolSuite struct {
	suite.Suite
	workersNumber uint
}

func (suite *WorkerPoolSuite) BeforeTest(suiteName, testName string) {
	suite.workersNumber = 100
}

func (suite *WorkerPoolSuite) Test_NewWorkerPool() {

	pool := NewWorkerPool(
		WithWorkerNumber(suite.workersNumber),
		WithResult(make(chan interface{}, suite.workersNumber)),
		WithJobber(NewJob()),
	)

	// Add jobs.
	go func() {
		no := 100
		for i := 0; i < no; i++ {
			pool.jobs.AddJob(rand.Intn(no))
		}
		pool.jobs.Close()
	}()

	fmt.Printf("%+v\n", pool)

	// Gets all job.
	go func() {
		ts := 0
		for val := range pool.jobs.GetJobs() {
			fmt.Println(val)
			ts++
		}
		fmt.Printf("times: %+v\n", ts)
	}()

	select {
	case <-pool.Done():
		fmt.Printf("done\n")
	case <-time.After(1 * time.Second):
		fmt.Printf("timeout\n")
	}
}
