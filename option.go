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
