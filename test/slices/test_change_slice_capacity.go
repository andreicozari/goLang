package main

import "fmt"

/*
Creating a slice with make
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

a := make([]int, 5)  // len(a)=5
To specify a capacity, pass a third argument to make:

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
*/

func main() {
	a := make([]int, 5)
	printS("a", a) //  l5 c5 0 0 0 0 0

	b := make([]int, 0, 5)
	printS("b", b) // l0 c5

	// Capacity changes only if the first index is changed and it's diff than 0
	c := b[:2]
	printS("c", c) // l2 c5 0 0

	// Capacity changes here:
	d := c[2:5]
	printS("d", d) // l3 c3 0 0 0
}

func printS(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
