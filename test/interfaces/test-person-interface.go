package main

import "fmt"

/* Interface is a collection of methods.
   And each type that is implementing all methods from an interfaces,
   is considering satisfying that interface
*/
type Person interface {
	speak(text string)
}

type Author struct {
	name string
}

/*
There is no explicit declaration of intent, no "implements" keyword.
Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.
*/
func (a Author) speak(text string) {
	fmt.Println(text)
}

func main() {
	var person Person = Author{"Rob Pike"}
	person.speak("Go is fun !")
}
