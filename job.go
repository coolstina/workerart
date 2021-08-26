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
