package main

import (
	"fmt"
	"time"
)

func main() {
	go a()
	go b()

	time.Sleep(2 * time.Second)
	fmt.Println("\n Main func exit!")
}

func a() {
	for i := 0; i < 100; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Print("O ")
	}
}

func b() {
	for i := 0; i < 100; i++ {
		time.Sleep(50 * time.Millisecond)
		fmt.Print("_ ")
	}
}
