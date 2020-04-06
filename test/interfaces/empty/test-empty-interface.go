package main

import "fmt"

/*

The empty interface
The interface type that specifies zero methods is known as the empty interface:
interface{}
An empty interface may hold values of any type. (Every type implements at least zero methods.)
Empty interfaces are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.

*/

func main() {
	var i interface{}

	// print the nil interface:
	fmt.Println(i)

	i = 1
	fmt.Println(i)

	i = "stringValue"
	fmt.Println(i)

	// see the Print method that takes an array of empty interfaces
	// this means that it accept any type
	fmt.Print()

}
