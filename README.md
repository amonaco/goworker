### goworker

I use this in a couple of projects that need a simple fan-out worker pool.
Personal project, use this at your own risk

Sample usage:
```
package main

import (
	"log"
	"time"

	goworker "github.com/amonaco/goworker"
)

func handler(task *goworker.Task) {
	log.Printf("[handler] received task: %v, %v\n", task.Key, task.Args)
}

func main() {

	worker := goworker.NewWorker(2, handler)
	worker.Start()

	task1 := &goworker.Task{
		Key:  "foo",
		Args: "bar",
	}
	worker.Push(task1)

	task2 := &goworker.Task{
		Key:  "baz",
		Args: "quux",
	}
	worker.Push(task2)

	time.Sleep(1000 * time.Millisecond)
	worker.Stop()

}
```
