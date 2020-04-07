package main

import (
	"fmt"
	"math"
)

/*
	Interface values:
Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

(value, type)
An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.
*/
type MyInterface interface {
	M()
}

type MyType struct {
	S string
}

// implement the interface method M() :
func (t *MyType) M() {
	fmt.Println(t.S)
}

type MyFloat float64

func (f MyFloat) M() {
	fmt.Println(f)
}

func main() {
	var i MyInterface

	i = &MyType{"Hello"}
	describe(i)
	i.M()

	i = MyFloat(math.Pi)
	describe(i)
	i.M()
}

func describe(i MyInterface) {
	fmt.Printf("(%v, %T)\n", i, i)
}
