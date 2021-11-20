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
			pool.AddJobStarting(rand.Intn(no))
		}
		pool.AddJobFinished()
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
