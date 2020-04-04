package main

import (
	"fmt"
	"math/rand"
)

func main() {

	if val := rand.Intn(122); val < 100 {
		fmt.Printf("%v Is less than 100", val)
	} else {
		fmt.Printf("%v Is greater than 100", val)
	}
}
