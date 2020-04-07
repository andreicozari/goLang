package main

import (
	"fmt"
)

// function that return double values:
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0.0, fmt.Errorf("Sqrt of negative value %v can't be parsed", x)
	}

	return 0, nil
}

func main() {
	handleSqrt(9)  // return 3
	handleSqrt(-4) // return error
}

func handleSqrt(n float64) {
	res, err := Sqrt(n)

	// nil -->  builtin type to mean : no value, like null in java or NONE java script
	if err != nil {
		fmt.Printf("ERROR: %v \n", err)
	} else {
		fmt.Printf("Sqrt of %v is \n", res)
	}
}
