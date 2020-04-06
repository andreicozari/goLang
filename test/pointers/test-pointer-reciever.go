package main

import (
	"fmt"
	"math"
)

/*

Pointers and functions:
Here we see the Abs and Scale methods rewritten as functions.
Again, try removing the * from line 23.
Can you see why the behavior changes? What else did you need to change for the example to compile?
*/
type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
