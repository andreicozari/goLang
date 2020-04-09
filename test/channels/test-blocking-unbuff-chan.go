package main

import "fmt"

/*
	The go func --> lunches a new anonymous go routine / a closure as anonymous function,
	This goR is pushing a int value of 1 into the unbuffered channel.

	The main goR then read from the channel and prints the value.
*/
func main() {
	unBuffChan := make(chan int)

	go func() {
		// push a int value to the unbuffered channel:
		unBuffChan <- 1
	}()

	// receive / get / read a int value from the unbuffered channel:
	intVal := <-unBuffChan
	fmt.Println(intVal)
}
