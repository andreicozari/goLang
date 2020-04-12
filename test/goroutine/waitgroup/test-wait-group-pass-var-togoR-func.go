package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	sayHello := func(val int) {
		wg.Add(1)
		defer wg.Done()
		fmt.Println(val)
	}

	for i := 1; i <= 9; i++ {
		go sayHello(i)
	}
	wg.Wait()
}
