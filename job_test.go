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
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestJobSuite(t *testing.T) {
	suite.Run(t, new(JobSuite))
}

type JobSuite struct {
	suite.Suite
	Job *Job
}

func (suite *JobSuite) BeforeTest(suiteName, testName string) {
	suite.Job = NewJob()
}

func (suite *JobSuite) Test_AddJob() {
	suite.Job.AddJob(func() {
		fmt.Println("First job")
	})
	suite.Job.AddJob(func() {
		fmt.Println("Second job")
	})

	assert.Len(suite.T(), suite.Job.list, 2)
}

func (suite *JobSuite) Test_GetJobs() {
	suite.Test_AddJob()

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	go func() {
		for job := range suite.Job.GetJobs() {
			f := job.(func())
			f()
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Printf("done\n")
	case <-time.After(200 * time.Millisecond):
		fmt.Printf("timeout\n")
	}
}
