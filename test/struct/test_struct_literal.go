package main

import "fmt"

type Line struct {
	X, Y int
}

var (
	v1 = Line{1, 2}  // has type Line
	v2 = Line{X: 1}  // Y:0 is implicit
	v3 = Line{}      // X:0 and Y:0
	p  = &Line{1, 2} // has type *Line
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
