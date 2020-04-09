package main

import (
	"fmt"
	"time"
)

func main() {
	//strChan := make(chan string)

	// try with buffered chan:
	strChan := make(chan string, 3)

	go sendValToChan("A", strChan)
	go sendValToChan("B", strChan)
	go sendValToChan("C", strChan)

	receive(strChan)
	receive(strChan)
	receive(strChan)

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
