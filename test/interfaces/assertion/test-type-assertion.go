package main

import "fmt"

func main() {
	var i interface{} = "stringValue"

	nr, ok := i.(int)
	fmt.Println(nr, ok)

	float, ok := i.(float64)
	fmt.Println(float, ok)

	str, ok := i.(string)
	fmt.Println(str, ok)

	intNr := i.(int) // panic
	fmt.Println(intNr)
}
