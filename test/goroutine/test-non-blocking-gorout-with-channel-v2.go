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
	//strChan := make(chan string)

	// try with buffered chan:
	wordsChan := make(chan Word, 3)

	valuesToSend := []Word{NewWord("One"), NewWord("Three"), NewWord("Four")}

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

func receive(wordsChan chan Word) {
	fmt.Println("Received ", <-wordsChan)
}
