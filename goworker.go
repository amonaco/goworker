package goworker

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type Task struct {
	Key  string
	Args interface{}
}

// Worker manages a pool of goroutines to process tasks.
type Worker struct {
	channels []chan *Task
	quit     chan struct{}
	Max      int
	Handler  func(work *Task)
	wg       sync.WaitGroup
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewWorker(max int, handler func(task *Task)) *Worker {
	worker := &Worker{
		Max:     max,
		Handler: handler,
		quit:    make(chan struct{}),
	}
	worker.channels = make([]chan *Task, max)
	for i := 0; i < max; i++ {
		worker.channels[i] = make(chan *Task, 10) // buffered channel
	}
	return worker
}

func (worker *Worker) Start() {
	for i := 0; i < worker.Max; i++ {
		log.Printf("[worker][%d] starting\n", i)
		worker.wg.Add(1)
		go worker.wrapHandler(worker.channels[i], i)
	}
}

func (worker *Worker) Stop() {
	close(worker.quit)
	for _, ch := range worker.channels {
		close(ch)
	}
	worker.wg.Wait()
}

func (worker *Worker) Push(task *Task) {
	id := getRandom(worker.Max)
	worker.channels[id] <- task
}

func (worker *Worker) wrapHandler(c chan *Task, id int) {
	defer worker.wg.Done()
	for {
		select {
		case <-worker.quit:
			return
		case work, ok := <-c:
			if !ok {
				return
			}
			log.Printf("[worker][%d] received task\n", id)
			worker.Handler(work)
		}
	}
}

func getRandom(max int) int {
	return rand.Intn(max)
}
