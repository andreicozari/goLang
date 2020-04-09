package main

import (
	"fmt"
	"time"
)

func main() {
	//strChan := make(chan string)

	// try with buffered chan:
	strChan := make(chan string, 3)

	valuesToSend := []string{"A", "B", "C"}

	for _, val := range valuesToSend {
		go sendValToChan(val, strChan)
	}

	for i := 0; i < len(valuesToSend); i++ {
		receive(strChan)
	}
}

func sendValToChan(val string, strChan chan string) {
	fmt.Println("\n sending ", val)
	time.Sleep(2 * time.Second)
	strChan <- val

	// the next line will be executed concurrently with the receiver
	fmt.Println("\n sent ", val)
}

func receive(strChan chan string) {
	fmt.Println("Received ", <-strChan)
}
