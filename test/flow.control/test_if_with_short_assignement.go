package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Seed the random number generator using the current time (nanoseconds since epoch):
	rand.Seed(time.Now().UnixNano())

	//if val := rand.Intn(122); val < 100 {
	if val := rand.Intn(200); val < 100 {
		fmt.Printf("%v Is less than 100", val)
	} else {
		fmt.Printf("%v Is greater than 100", val)
	}
}
