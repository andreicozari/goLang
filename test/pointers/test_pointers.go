package main

import "fmt"

//A pointer holds the memory address of a value.

func main() {
	i, j := 8, 10

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 16         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 2    // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
