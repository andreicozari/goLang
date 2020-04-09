package main

import (
	"fmt"
	"time"
)

func main() {
	strChan := make(chan string)

	// try with buffered chan:
	//strChan := make(chan string, 3)

	go sendValToChan("A", strChan)
	go sendValToChan("B", strChan)
	go sendValToChan("C", strChan)

	read(strChan)
	read(strChan)
	read(strChan)

}

func sendValToChan(val string, strChan chan string) {
	fmt.Println("\n sending ", val)
	time.Sleep(2 * time.Second)
	strChan <- val
	fmt.Println("\n sent ", val)
}

func read(strChan chan string) {
	fmt.Println("Received ", <-strChan)
}
