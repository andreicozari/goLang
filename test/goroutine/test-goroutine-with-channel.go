package main

import (
	"fmt"
)

func main() {
	//strChan := make(chan string)

	// try with buffered chan:
	strChan := make(chan string, 3)

	go sendABC(strChan)

	fmt.Println("\n Received ", <-strChan)
	fmt.Println("\n Received ", <-strChan)
	fmt.Println("\n Received ", <-strChan)

	/*for ; ; {
		fmt.Println("\n Received ",  <- strChan)

		for _, r := range `12345` {
			fmt.Printf("\r%c", r)
			time.Sleep(1000 * time.Millisecond)
		}
	}*/
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
