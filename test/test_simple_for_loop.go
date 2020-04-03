package main

import "fmt"

//The init and post statements are optional.
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	///////////////////////////////////////
	message := `For is Go's "while"
			At that point you can drop the semicolons: C's while is spelled for in Go.`

	fmt.Println(message)

	fmt.Println(sum)
	sum = 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

}
