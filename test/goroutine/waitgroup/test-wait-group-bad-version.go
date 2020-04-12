package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	/*sayHello := func(val int) {
		//wg.Add(1)
		defer wg.Done()
		fmt.Println(val)
	}*/

	// the for loop is in the the mainGoR and it's i index is not synchronized inside
	// and the goR is has another / it's own stack, so --> we don't know at each for loop iteration
	// which value will have var i, and even the for loop can finish sooner that all 9 loop iterations:
	for i := 1; i <= 9; i++ {
		wg.Add(1)
		go func() {
			wg.Add(1)
			defer wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()
}
