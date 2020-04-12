package main

import (
	"fmt"
	"time"
)

func main() {
	// Define the wait group counter
	//var waitGroupCounter sync.WaitGroup

	// Add one goroutine to the  WaitGroup counter:
	// I have one goR to wait for:
	//waitGroupCounter.Add(1)

	// Then i need to decrement the counter -->
	//		tell to main goR that there is no need to wait any more for the goR to run
	// Instead of calling directly the go showNTimesAWord , we will use a go call to a closure
	//  for handling the WaitGroup counter decrement
	/*go func() {
		// only when the goR is done --> do with defer --> decrement wg counter with:
		defer waitGroupCounter.Done()
		showNTimesAWord(2, "Apple", ch)

	}()*/

	// use channels:
	ch := make(chan interface{})
	go showNTimesAWord(2, "Apple", ch)
	go showNTimesAWord(2, "Banana", ch)

	// the main goR is the receiver
	// As a receiver you should never close the channel because you don't know if there are still values in channel
	for {
		val, isChannelOpen := <-ch
		if isChannelOpen {
			fmt.Println(val)
		} else {
			fmt.Println("The channel was closed by the sender. There are no more items to receive.")
			break
		}
	}

	//go showNTimesAWord(3, "Banana")

	// get rid of sleep time
	// we don't need any more, a delay here, to wait for the others goRs to execute / to finish
	//time.Sleep(time.Second)
	// And use here a wg.wait() instead a delay with sleep time:
	//waitGroupCounter.Wait()
}

func showNTimesAWord(n int, str string, ch chan interface{}) {
	for i := 1; i <= n; i++ {
		ch <- fmt.Sprintf("%d %v", i, str)
		time.Sleep(200 * time.Millisecond)
	}

	// The sender should be the only part that closes a channel:
	close(ch)
}
