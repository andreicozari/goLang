package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

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
	wordsChan := make(chan Word, 3)

	valuesToSend := []Word{NewWord("One"), NewWord("Two"), NewWord("Three")}

	// send concurrently:
	for _, val := range valuesToSend {
		go sendValToChan(val, wordsChan)
	}

	for i := 0; i < len(valuesToSend); i++ {
		receive(wordsChan)
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

func receive(strChan chan Word) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("\n Received ", <-strChan)
	}()
	wg.Wait()
}
