package main

import "fmt"

type Author struct {
	name string
}

func (a Author) speak(s string) {
	fmt.Printf("%v says: %v", a.name, s)
}

func main() {
	a := Author{}
	a.name = "Rob Pike"
	a.speak("'Concurrency is the key to designing high performance network services!'\n")

	a2 := Author{"Robert Cecil Martin"}
	a2.speak("'Clean up your code!'\n")
}
