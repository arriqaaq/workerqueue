package workerqueue

import (
	// "math/rand"
	"testing"
	"time"
)

func TestDispatcher_Run_Test_10m(t *testing.T) {
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
	time.Sleep(1 * time.Millisecond)
	d.Stop()
}

// func BenchmarkDispatcher1(b *testing.B) {
// 	// run the Fib function b.N times
// 	for n := 0; n < b.N; n++ {
// 		delay := time.Duration(rand.Intn(1000)) * time.Millisecond
// 		b.Log(b.N, delay)
// 		jobs := []Job{
// 			newJob("1", delay),
// 			newJob("2", delay),
// 			newJob("3", delay),
// 			newJob("4", delay),
// 			newJob("5", delay),
// 			newJob("6", delay),
// 			newJob("7", delay),
// 			newJob("8", delay),
// 			newJob("9", delay),
// 			newJob("10", delay),
// 		}

// 		d := NewDispatcher("dp_1", 5)
// 		d.Run()

// 		for _, j := range jobs {
// 			d.AddJob(j)
// 		}
// 		time.Sleep(1 * time.Millisecond)
// 		d.Stop()
// 	}
// }
