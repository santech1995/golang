//Simple program to illustrate the use of 'atomic' to avoid a race condition during concurrency

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	var counter int64
	//here counter is the shared variable which is accessed by multiple goroutines.
	//Race condition is when a shared variable is accessed by multiple go
	//routines.
	const gs = 10
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			//write to the counter using atomic
			atomic.AddInt64(&counter, 1) //takes addr of value to be changed
			//and the value bywhich it has to be changed(here ++ => add 1)

			runtime.Gosched() //indicates the system to complete other pending tasks

			//read from the counter using atomic
			fmt.Println("Counter:", atomic.LoadInt64(&counter))

			wg.Done()
		}() //the anonymous func is being called
		fmt.Println("Goroutines:", runtime.NumGoroutine())
		fmt.Println("CPUs:", runtime.NumCPU())
	}
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	wg.Wait() //implies wait until all the execution completes
	fmt.Println("counter:", counter)
}

//OUTPUT:
// CPUs 4
// Goroutines: 1
// Goroutines: 2
// CPUs: 4
// Counter: 1
// Counter: 2
// Goroutines: 3
// CPUs: 4
// Goroutines: 2
// CPUs: 4
// Counter: 3
// Goroutines: 3
// Counter: 4
// CPUs: 4
// Goroutines: 2
// Counter: 5
// CPUs: 4
// Goroutines: 2
// CPUs: 4
// Counter: 6
// Goroutines: 3
// Counter: 7
// CPUs: 4
// Goroutines: 2
// Counter: 8
// CPUs: 4
// Goroutines: 2
// CPUs: 4
// Counter: 9
// Counter: 10
// Goroutines: 3
// CPUs: 4
// Goroutines: 1
// counter: 10
