package task

import "time"

type Task struct{}

func NewTask() *Task {
	return &Task{}
}

func (t *Task) Process(number uint) uint {
	var sum uint
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(1 * time.Second)
	return sum
}
