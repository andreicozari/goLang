package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "hello"

	wg.Add(1)

	say := func(msg string) {
		defer wg.Done()
		fmt.Println(msg)
		salutation = "welcome"
	}

	go say(salutation)

	wg.Wait()
	fmt.Println(salutation)
}
