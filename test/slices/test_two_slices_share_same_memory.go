package main

import (
	"fmt"
)

// Test same memory address / location of arrays:
// printf %p &

func main() {
	intSlices()
	fmt.Println("\n")
	stringSlices()

	s := []string{"aLiteral"}
	sPointer := s[:]
	fmt.Printf("%v %v %p %p", sPointer, s, &sPointer, &s)

}

func stringSlices() {
	s1 := []string{"A", "B"}
	s2 := []string{"A", "B"}

	// compare s1 and s2 are sharing the same memory location:
	fmt.Printf("s1 = %v, s2 = %v \n", s1, s2)
	fmt.Printf("%p  %p \n", &s1, &s2)

	checkSameMemoryString(s1, s2)
}

func intSlices() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:]
	s3 := []int{2, 3}

	// compare s2 and s3 are sharing the same memory location:
	fmt.Printf("s2 = %v, s3 = %v \n", s2, s3)
	fmt.Printf("%p  %p \n", &s2, &s3)

	checkSameMemory(s2, s3)
}

func checkSameMemory(s1, s2 []int) {
	if &s1 != &s2 {
		fmt.Println("The slices are using different memory address. ")
	} else {
		fmt.Println("The slices are using the same memory address. ")
	}
}

func checkSameMemoryString(s1, s2 []string) {
	if &s1 != &s2 {
		fmt.Println("The slices are using different memory address. ")
	} else {
		fmt.Println("The slices are using the same memory address. ")
	}
}
