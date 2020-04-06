package main

import "fmt"

type MyIterface interface {
	method()
}

func main() {

	var myInterface MyIterface

	fmt.Printf("Show the (value; type) of nil interface: (%v , %T) \n",
		myInterface, myInterface)
}
