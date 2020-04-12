package main

import (
	"fmt"
	"time"
)

type Word struct {
	val    string
	length int
}

func NewWord(val string) Word {
	return Word{val: val}
}

func (word Word) String() string {
	return fmt.Sprintf("{%v, %v}", word.val, word.length)
}

func main() {
	// unbuffered channel for making goroutines synchronized
	//strChan := make(chan string)

	// try with buffered chan:
	// non blocking to send first 3 items, regardless if there is any receiver to read the items
	// on the fourth send --> the GoR will block
	wordsChan := make(chan Word, 3)

	valuesToSend := []Word{NewWord("One"), NewWord("Three"), NewWord("Four")}

	// send concurrently:
	for _, val := range valuesToSend {
		go sendValToChan(val, wordsChan)
	}

	// what if we don't know the number of data in a channel ?
	// Let's say: we don't know how many goroutines was started
	/*for i := 0; i < len(valuesToSend); i++ {
		receive(wordsChan)
	}*/

	// we use select statement:
	// select can listen for data incoming in channel:
	// While our goroutines were running concurrently, they were not running in parallel --> GOMAXPROCS nr of CPU
	select {
	case word := <-wordsChan:
		fmt.Println("Received ", word)
	}
}

func sendValToChan(word Word, wordsChan chan Word) {
	word.length = len(word.val)
	fmt.Println("\n sending ", word)
	time.Sleep(time.Second)
	wordsChan <- word

	// the next line will be executed concurrently with the receiver
	fmt.Println("\n sent ", word)
}

func receive(wordsChan chan Word) {
	fmt.Println("Received ", <-wordsChan)
}
