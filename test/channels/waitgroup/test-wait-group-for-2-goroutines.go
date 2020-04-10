package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Define the wait group counter
	var waitGroupCounter sync.WaitGroup

	// Add one goroutine to the  WaitGroup counter:
	// I have one goR to wait for:
	waitGroupCounter.Add(1)

	// Then i need to decrement the counter --> to tell when no need to wait more for the goR:
	// Instead of calling directly the go showNTimesAWord , we will use a go call to a closure
	//  for handling the WaitGroup counter decrement

	go func() {
		showNTimesAWord(2, "Apple")

		// decrement wg counter:
		waitGroupCounter.Done()
	}()

	//go showNTimesAWord(3, "Banana")

	// get rid of sleep time
	// we don't need any more, a delay here, to wait for the others goRs to execute / to finish
	//time.Sleep(time.Second)
	// And use here a wg.wait() instead a delay with sleep time:
	waitGroupCounter.Wait()
}

func showNTimesAWord(n int, str string) {
	for i := 1; i <= n; i++ {
		fmt.Println(i, str)
		time.Sleep(200 * time.Millisecond)
	}
}
