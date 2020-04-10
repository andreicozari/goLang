package main

import (
	"fmt"
	"time"
)

func main() {
	go showNTimesAWord(2, "Apple")
	go showNTimesAWord(3, "Banana")

	time.Sleep(time.Second)
}

func showNTimesAWord(n int, str string) {
	for i := 1; i <= n; i++ {
		fmt.Println(i, str)
		time.Sleep(200 * time.Millisecond)
	}
}
