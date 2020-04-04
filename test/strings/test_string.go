package main

import (
	"fmt"
	//"testing"
)

func main() {
	test_string_slice()
}

func test_string_slice() {
	val := "A random string value"

	poem := `
		The firs line
		The second One !
		The third
		...
	`

	fmt.Println(val[:4])
	fmt.Println(val[0:])
	fmt.Println(val[:len(val)])
	fmt.Println(val[:])
	fmt.Println(val[0:6])
	fmt.Println(val[1:6])

	fmt.Println(poem)
}
