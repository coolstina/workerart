![workerart](assets/banner/workerart.jpg)

Fast use worker pool process task.

## Install

```shell script
go get github.com/fishfinal/workerart
```

## What is workerart?

workerart 是一个快速使用协程工作池的实现，而不是当需要有工作池的使用场景时去重复的造轮子，可以提高你的开发效率，同是也不失Go并发处理任务的优雅。workerart 支持：

- 通过选项构建工作池。
- 通过Jobber接口实现你自己特有的Job.
- 自定义任务回调函数。

## Why use workerart?

为了优雅的处理多个任务，当然你自己也可以实现工作池，workerart 仅仅只是让你更快速使用工作池来提高任务执行效率。

## How to use?


```go
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fishfinal/workerart"
)

func main() {
	pool := workerart.NewWorkerPool()

	// Add jobs.
	go func() {
		no := 10000
		for i := 0; i < no; i++ {
			pool.AddJobs(rand.Intn(no))
		}
		pool.CloseJob()
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
```
