package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)

	ch2 := make(chan string)

	go func() {
		for {
			ch1 <- "Every 500ms"
			// every half second:
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			ch2 <- "Every e seconds"
			// every two seconds:
			time.Sleep(time.Second * 2)
		}
	}()

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}

}
