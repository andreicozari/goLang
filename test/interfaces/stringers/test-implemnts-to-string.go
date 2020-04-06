package main

import "fmt"

/*
Stringers:
One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}
A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
*/

type Person struct {
	name string
	age  int
}

// Like implement toString in java:
func (p Person) String() string {
	return fmt.Sprintf(" Person -> name: %v, age: %v \n", p.name, p.age)
}

func main() {
	p := Person{name: "Bob", age: 35}
	fmt.Println(p)
}
