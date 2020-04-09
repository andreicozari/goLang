package main

import (
	"fmt"
	"time"
)

func main() {
	//strChan := make(chan string)

	// try with buffered chan:
	strChan := make(chan string, 3)

	go sendABC(strChan)

	waitOneSec()
	fmt.Println("\n Received ", <-strChan)

	waitOneSec()
	fmt.Println("\n Received ", <-strChan)

	waitOneSec()
	fmt.Println("\n Received ", <-strChan)
}

func sendABC(strChan chan string) {

	fmt.Println("\n sending a")
	strChan <- "a"
	fmt.Println("\n sent a")
	//time.Sleep(500 * time.Millisecond)

	fmt.Println("\n sending b")
	strChan <- "b"
	fmt.Println("\n send b")
	//time.Sleep(500 * time.Millisecond)

	fmt.Println("\n sending c")
	strChan <- "c"
	fmt.Println("\n sent c")
	//time.Sleep(500 * time.Millisecond)
}

func waitOneSec() {
	for _, r := range `123` {
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("\r%c", r)
	}
}
