package main

import (
	"fmt"
	"time"
)

/*
A go routine is like a running command in the shell. It runs in background.
It has it's own stack. The stack can shrink and expand as it's required automatically
There can be hundred of thousand of go routines.
Go routines are communicating between them through Channels

Don't communicate by sharing memory, share memory by communicating:
https://www.youtube.com/watch?v=f6kdp27TYZs
Channels are synchronized --> it means that all the input of the channel are sequential

*/

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
