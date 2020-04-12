package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

/*
A newly minted goroutine is given a few kilobytes, which is almost always
enough. When it isn’t, the run-time grows (and shrinks) the memory for
storing the stack automatically, allowing many goroutines to live in a
modest amount of memory. The CPU overhead averages about three cheap
instructions per function call. It is practical to create hundreds of thousands
of goroutines in the same address space. If goroutines were just threads,
system resources would run out at a much smaller number.

*/

/*
	1e = 10
	1e4 = 10 at power 4 = 10 * 4 = 1000
	If you create 1000 goR -> it takes les than 1Mb memory per total.

	With OS threads it would cost :
	try running this  cmd in cli:
	taskset -c 0 perf bench sched pipe -T

	This produces:
	# Running 'sched/pipe' benchmark:
	# Executed 1000000 pipe operations between two threads
	Total time: 2.935 [sec]
	2.935784 usecs/op
	340624 ops/sec


	Application context switching - communication between two goroutines ==> is much much cheaper,
	  than the Operating System context switching.
*/

func main() {

	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}

	var wg sync.WaitGroup

	noop := func() { wg.Done(); <-c }

	const numGoroutines = 1e4
	wg.Add(numGoroutines)

	before := memConsumed()
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}

	wg.Wait()

	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})
	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}
	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin)
	wg.Wait()
}
