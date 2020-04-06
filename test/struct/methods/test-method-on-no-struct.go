package main

import "fmt"

type MyInt int

// Absolute value:
func (nr MyInt) abs() MyInt {
	if nr < 0 {
		return nr * -1
	}
	return nr
}

func main() {
	one := MyInt(1)
	fmt.Println(one.abs())

	minusOne := MyInt(-1)
	fmt.Println(minusOne.abs())
}
