package main

import (
	"fmt"
)

func main() {

	sequenceChannel := make(chan interface{}, 24)

	go push(sequenceChannel)

	for char := range sequenceChannel {
		fmt.Print(char)
	}
}

func push(ch chan interface{}) {
	for i := 0; i < cap(ch); i++ {
		// the select is used to listen the operation over the channels:
		// only one case will be chosen at a time/ like switch operation:
		select {
		case ch <- "a":
		case ch <- "b":
		case ch <- "c":
		case ch <- "d":
		case ch <- "e":
		case ch <- "f":
		case ch <- "1":
		case ch <- "2":
		case ch <- "3":
		case ch <- "4":
		case ch <- "5":
		case ch <- "6":
		}
	}

	// Notify the receiver that there are no more values in the channel:
	close(ch)
}
