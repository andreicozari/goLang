package main

import "fmt"

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
