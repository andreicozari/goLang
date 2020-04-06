package main

import "fmt"

/*
Nil interface values:
A nil interface value holds neither value nor concrete type.
Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.
*/

type MyInterface interface {
	method()
}

func main() {

	var myInterface MyInterface

	fmt.Printf("Show the (value; type) of nil interface: (%v , %T) \n",
		myInterface, myInterface)
}
