package main

import (
	"fmt"
	"time"
)

func main() {
	unBuffChan := make(chan int)

	nrOfItems := 2

	// push 2 values into the unbuffered channel:
	go sendItemsToChan(nrOfItems, unBuffChan)

	// read 2 values from the unbuffered channel:
	readItemsFromChan(nrOfItems, unBuffChan)
}

func sendItemsToChan(nrOfItems int, intChan chan int) {
	for i := 0; i < nrOfItems; i++ {
		// wait / sleep one second:
		time.Sleep(time.Second)
		intChan <- i
	}
}

func readItemsFromChan(nrOfItems int, intChan chan int) {
	for i := 0; i < nrOfItems; i++ {
		fmt.Println(<-intChan)
	}
}
