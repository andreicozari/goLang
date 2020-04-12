package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	//strChan := make(chan string)

	// try with buffered chan:
	strChan := make(chan string, 3)

	go sendABC(strChan)

	for i := 1; i <= 3; i++ {
		receive(strChan)
	}

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

// the receiver:
func receive(strChan chan string) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("\n Received ", <-strChan)
	}()
	wg.Wait()
}
