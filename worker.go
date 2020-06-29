package goworker

import (
	"log"
	"math/rand"
	"time"
)

type Task struct {
	Key  string
	Args interface{}
}

type Worker struct {
	channels [](chan *Task)
	quit     chan bool
	Max      int
	Handler  func(work *Task)
}

func NewWorker(max int, handler func(task *Task)) *Worker {
	worker := &Worker{
		Max:     max,
		Handler: handler,
	}

	worker.channels = make([](chan *Task), max)
	worker.quit = make(chan bool)

	for i := 0; i < max; i++ {
		worker.channels[i] = make(chan *Task) //, 10)
	}

	return worker
}

func (worker *Worker) Start() {
	i := 0
	for i = 0; i < worker.Max; i++ {
		log.Printf("[worker][%d] starting\n", i)
		go worker.wrapHandler(worker.channels[i], i)
	}
}

func (worker *Worker) Stop() {
	worker.quit <- true
}

func (worker *Worker) Push(task *Task) {
	lane := getRandom(worker.Max)
	worker.channels[lane] <- task
}

func (w *Worker) wrapHandler(c chan *Task, lane int) {
	for {
		select {
		case <-w.quit:
			return
		case work := <-c:
			log.Printf("[worker][%d] received task\n", lane)
			w.Handler(work)
		}
	}
}

func getRandom(max int) int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max = max - 1
	return rand.Intn(max-min+1) + min
}
