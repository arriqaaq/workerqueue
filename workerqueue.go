package workerqueue

import (
	"fmt"
	"sync"
	"time"
)

const (
	DefaultChannelSize = 100
)

type Job interface {
	Execute() error
	Name() string
}

func newJob(name string, delay time.Duration) Job {
	return &job{name, delay}
}

// Job holds the attributes needed to perform unit of work.
type job struct {
	name  string
	delay time.Duration
}

func (j *job) Execute() error {
	time.Sleep(j.delay)
	return nil
}

func (j *job) Name() string {
	return j.name
}

// NewWorker creates takes a numeric id and a channel w/ worker pool.
func NewWorker(id int, workerPool chan chan Job, wg *sync.WaitGroup) *Worker {
	return &Worker{
		id:         id,
		jobQueue:   make(chan Job),
		workerPool: workerPool,
		wg:         wg,
	}
}

type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	wg         *sync.WaitGroup
}

func (w *Worker) start() {
	// defer w.wg.Done()
	go func() {
		defer func() {
			w.wg.Done()
		}()

		w.workerPool <- w.jobQueue

		for job := range w.jobQueue {
			w.workerPool <- w.jobQueue
			job.Execute()
		}
	}()
}

// Close ensured the channel is closed for sending, but waits for all messages to be consumed
func (w *Worker) close() {
	close(w.jobQueue)
}

// NewDispatcher creates, and returns a new Dispatcher object.
func NewDispatcher(name string, maxWorkers int) *Dispatcher {
	workerPool := make(chan chan Job, maxWorkers)
	jobQueue := make(chan Job, DefaultChannelSize)

	return &Dispatcher{
		name:       name,
		jobQueue:   jobQueue,
		maxWorkers: maxWorkers,
		workerPool: workerPool,
		wg:         &sync.WaitGroup{},
		doneCh:     make(chan bool),
	}
}

type Dispatcher struct {
	name       string
	workerPool chan chan Job
	maxWorkers int
	jobQueue   chan Job
	wg         *sync.WaitGroup
	doneCh     chan bool
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		id := i + 1
		d.wg.Add(1)
		worker := NewWorker(id, d.workerPool, d.wg)
		worker.start()
	}
	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for job := range d.jobQueue {
		workerJobQueue := <-d.workerPool
		workerJobQueue <- job
	}
	// close(d.workerPool)
	// for i := 0; i < d.maxWorkers; i++ {
	for {
		select {
		case worker, ok := <-d.workerPool:
			if ok {
				close(worker)
			} else {
				d.doneCh <- true
				return
			}

		}
	}

}

func (d *Dispatcher) AddJob(job Job) {
	d.jobQueue <- job
}

func (d *Dispatcher) Stop() {
	// No more Adding jobs to the jobqueue function
	close(d.jobQueue)
	d.wg.Wait()
	close(d.workerPool)
	<-d.doneCh
}
