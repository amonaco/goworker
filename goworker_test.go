package goworker

import (
	"log"
	"testing"
	"time"
)

func handler(task *Task) {
	log.Printf("[handler] received task: %v, %v\n", task.Key, task.Args)
}

func TestWorker(t *testing.T) {
	worker := NewWorker(2, handler)
	worker.Start()

	task1 := &Task{
		Key:  "foo",
		Args: "bar",
	}
	worker.Push(task1)

	task2 := &Task{
		Key:  "baz",
		Args: "quux",
	}
	worker.Push(task2)

	time.Sleep(1000 * time.Millisecond)
	worker.Stop()
}
