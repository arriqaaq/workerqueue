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
		quitChan:   make(chan bool),
		wg:         wg,
	}
}

type Worker struct {
	id         int
	jobQueue   chan Job
	workerPool chan chan Job
	quitChan   chan bool
	wg         *sync.WaitGroup
}

func (w *Worker) start() {
	// defer w.wg.Done()
	go func() {

		for w.jobQueue != nil {
			// Add my jobQueue to the worker pool.

			w.workerPool <- w.jobQueue

			select {
			case job, ok := <-w.jobQueue:
				if !ok {
					// fmt.Println("nil job worker", w.id, job, len(w.jobQueue))
					w.jobQueue = nil
					w.wg.Done()
					continue
				}

				// Dispatcher has added a job to my jobQueue.
				job.Execute()

			case <-w.quitChan:
				// We have been asked to stop.
				fmt.Printf("worker%d stopping\n", w.id)
				return
			}

		}
	}()
}

// Close ensured the channel is closed for sending, but waits for all messages to be consumed
func (w *Worker) close() {
	close(w.jobQueue)
}

// Don't call stop, unless explicitly needed, else queued job will fail
func (w *Worker) stop() {
	go func() {
		w.quitChan <- true
	}()
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
		quitChan:   make(chan bool),
		wg:         &sync.WaitGroup{},
		workerMap:  make(map[int]*Worker),
	}
}

type Dispatcher struct {
	name       string
	workerPool chan chan Job
	maxWorkers int
	jobQueue   chan Job
	workerMap  map[int]*Worker
	quitChan   chan bool
	wg         *sync.WaitGroup
	l          sync.Mutex
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.maxWorkers; i++ {
		id := i + 1
		d.wg.Add(1)
		worker := NewWorker(id, d.workerPool, d.wg)
		worker.start()
		d.workerMap[id] = worker
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job, ok := <-d.jobQueue:
			if !ok {
				// fmt.Println("nil job dispatch", job, len(d.jobQueue))
				d.jobQueue = nil
				continue
			}
			workerJobQueue := <-d.workerPool
			// fmt.Println("pushing to queue: ", workerJobQueue)
			workerJobQueue <- job
		case <-d.quitChan:
			// We have been asked to stop.
			// fmt.Printf("dispatcher coming to halt\n")
			d.shutWorkers()
			return
		}
	}
}

func (d *Dispatcher) AddJob(job Job) {
	d.jobQueue <- job
}

func (d *Dispatcher) Stop() {
	// No more Adding jobs to the jobqueue function
	close(d.jobQueue)
	d.quitChan <- true
	d.wg.Wait()
}

func (d *Dispatcher) shutWorkers() {
	for _, worker := range d.workerMap {
		worker.close()
	}
}
