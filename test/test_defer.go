package main

import "fmt"

func main() {
	defer fmt.Println("Second -  Executed after last line in function ")

	fmt.Println("First !")
}
