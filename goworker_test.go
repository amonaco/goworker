package goworker

import (
	"sync"
	"testing"
	"time"
)

func TestWorker(t *testing.T) {
	var mu sync.Mutex
	var results []string

	handler := func(task *Task) {
		mu.Lock()
		results = append(results, task.Key)
		mu.Unlock()
	}

	worker := NewWorker(2, handler)
	worker.Start()

	task1 := &Task{Key: "foo", Args: "bar"}
	task2 := &Task{Key: "baz", Args: "quux"}
	worker.Push(task1)
	worker.Push(task2)

	time.Sleep(500 * time.Millisecond)
	worker.Stop()

	mu.Lock()
	defer mu.Unlock()
	if len(results) != 2 {
		t.Errorf("expected 2 tasks processed, got %d", len(results))
	}

}	}
}
