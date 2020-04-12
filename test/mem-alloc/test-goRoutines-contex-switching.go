package main

import (
	"sync"
	"testing"
)

/*
	Application context switching - communication between two goroutines ==> is much much cheaper,
	  than the Operating System context switching:

	With OS threads it would cost :
	try running this  cmd in cli:
	taskset -c 0 perf bench sched pipe -T
	This produces:
	# Running 'sched/pipe' benchmark:
	# Executed 1000000 pipe operations between two threads
	Total time: 2.935 [sec]
	2.935784 usecs/op
	340624 ops/sec


	GO performance:
	go test -bench=. -cpu=1 \
	test/mem-alloc/test-goRoutines-contex-switching.go
	That’s 0.225 μs, or 92% faster than an OS context switch on my machine, which if you recall took 1.467 μs.
*/

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
