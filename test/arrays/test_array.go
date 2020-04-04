package main

import (
	"fmt"
)

func main() {
	var a [2]string
	a[0] = "item1"
	a[1] = "item2"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	showEvenNums(primes)
}

func showEvenNums(s []int) {
	for _, val := range s {
		if val%2 == 0 {
			fmt.Println(val)
		}
	}
}
