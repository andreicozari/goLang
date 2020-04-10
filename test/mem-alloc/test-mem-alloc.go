package main

import (
	"fmt"
	"runtime"
	"sync"
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
