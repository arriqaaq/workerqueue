package workerqueue

import (
	"reflect"
	"sync"
	"testing"
	"time"
)

func Test_job_Execute(t *testing.T) {
	type fields struct {
		Name  string
		Delay time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{"1", fields{"j1", 1 * time.Millisecond}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &job{
				name:  tt.fields.Name,
				delay: tt.fields.Delay,
			}
			if err := j.Execute(); (err != nil) != tt.wantErr {
				t.Errorf("job.Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewWorker(t *testing.T) {
	type args struct {
		id         int
		workerPool chan chan Job
		wg         *sync.WaitGroup
	}
	tests := []struct {
		name string
		args args
		want *Worker
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewWorker(tt.args.id, tt.args.workerPool, tt.args.wg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWorker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWorker_start(t *testing.T) {
	type fields struct {
		id         int
		jobQueue   chan Job
		workerPool chan chan Job
		quitChan   chan bool
		wg         *sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Worker{
				id:         tt.fields.id,
				jobQueue:   tt.fields.jobQueue,
				workerPool: tt.fields.workerPool,
				quitChan:   tt.fields.quitChan,
				wg:         tt.fields.wg,
			}
			w.start()
		})
	}
}

func TestWorker_close(t *testing.T) {
	type fields struct {
		id         int
		jobQueue   chan Job
		workerPool chan chan Job
		quitChan   chan bool
		wg         *sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Worker{
				id:         tt.fields.id,
				jobQueue:   tt.fields.jobQueue,
				workerPool: tt.fields.workerPool,
				quitChan:   tt.fields.quitChan,
				wg:         tt.fields.wg,
			}
			w.close()
		})
	}
}

func TestWorker_stop(t *testing.T) {
	type fields struct {
		id         int
		jobQueue   chan Job
		workerPool chan chan Job
		quitChan   chan bool
		wg         *sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Worker{
				id:         tt.fields.id,
				jobQueue:   tt.fields.jobQueue,
				workerPool: tt.fields.workerPool,
				quitChan:   tt.fields.quitChan,
				wg:         tt.fields.wg,
			}
			w.stop()
		})
	}
}

func TestNewDispatcher(t *testing.T) {
	type args struct {
		name       string
		maxWorkers int
	}
	tests := []struct {
		name string
		args args
		want *Dispatcher
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDispatcher(tt.args.name, tt.args.maxWorkers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDispatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_Run(t *testing.T) {
	type fields struct {
		name       string
		workerPool chan chan Job
		maxWorkers int
		jobQueue   chan Job
		workerMap  map[int]*Worker
		quitChan   chan bool
		wg         *sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				name:       tt.fields.name,
				workerPool: tt.fields.workerPool,
				maxWorkers: tt.fields.maxWorkers,
				jobQueue:   tt.fields.jobQueue,
				workerMap:  tt.fields.workerMap,
				quitChan:   tt.fields.quitChan,
				wg:         tt.fields.wg,
			}
			d.Run()
		})
	}
}

func TestDispatcher_dispatch(t *testing.T) {
	type fields struct {
		name       string
		workerPool chan chan Job
		maxWorkers int
		jobQueue   chan Job
		workerMap  map[int]*Worker
		quitChan   chan bool
		wg         *sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				name:       tt.fields.name,
				workerPool: tt.fields.workerPool,
				maxWorkers: tt.fields.maxWorkers,
				jobQueue:   tt.fields.jobQueue,
				workerMap:  tt.fields.workerMap,
				quitChan:   tt.fields.quitChan,
				wg:         tt.fields.wg,
			}
			d.dispatch()
		})
	}
}

func TestDispatcher_AddJob(t *testing.T) {
	type fields struct {
		name       string
		workerPool chan chan Job
		maxWorkers int
		jobQueue   chan Job
		workerMap  map[int]*Worker
		quitChan   chan bool
		wg         *sync.WaitGroup
	}
	type args struct {
		job Job
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				name:       tt.fields.name,
				workerPool: tt.fields.workerPool,
				maxWorkers: tt.fields.maxWorkers,
				jobQueue:   tt.fields.jobQueue,
				workerMap:  tt.fields.workerMap,
				quitChan:   tt.fields.quitChan,
				wg:         tt.fields.wg,
			}
			d.AddJob(tt.args.job)
		})
	}
}

func TestDispatcher_Stop(t *testing.T) {
	type fields struct {
		name       string
		workerPool chan chan Job
		maxWorkers int
		jobQueue   chan Job
		workerMap  map[int]*Worker
		quitChan   chan bool
		wg         *sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				name:       tt.fields.name,
				workerPool: tt.fields.workerPool,
				maxWorkers: tt.fields.maxWorkers,
				jobQueue:   tt.fields.jobQueue,
				workerMap:  tt.fields.workerMap,
				quitChan:   tt.fields.quitChan,
				wg:         tt.fields.wg,
			}
			d.Stop()
		})
	}
}

func TestDispatcher_shutWorkers(t *testing.T) {
	type fields struct {
		name       string
		workerPool chan chan Job
		maxWorkers int
		jobQueue   chan Job
		workerMap  map[int]*Worker
		quitChan   chan bool
		wg         *sync.WaitGroup
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dispatcher{
				name:       tt.fields.name,
				workerPool: tt.fields.workerPool,
				maxWorkers: tt.fields.maxWorkers,
				jobQueue:   tt.fields.jobQueue,
				workerMap:  tt.fields.workerMap,
				quitChan:   tt.fields.quitChan,
				wg:         tt.fields.wg,
			}
			d.shutWorkers()
		})
	}
}

func TestDispatcher_Run_Test_1(t *testing.T) {
	delay := 10 * time.Millisecond
	jobs := []Job{
		newJob("1", delay),
		newJob("2", delay),
		newJob("3", delay),
		newJob("4", delay),
		newJob("5", delay),
		newJob("6", delay),
		newJob("7", delay),
		newJob("8", delay),
		newJob("9", delay),
		newJob("10", delay),
	}

	d := NewDispatcher("dp_1", 5)
	d.Run()

	for _, j := range jobs {
		d.AddJob(j)
	}
	// time.Sleep(1 * time.Second)
	d.Stop()
}
