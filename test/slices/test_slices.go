package main

import "fmt"

func main() {
	nums := [6]int{0, 1, 2, 3}

	var slice []int = nums[1:3]

	// the other are initialized with 0 by default:
	fmt.Println(nums)
	fmt.Println(slice)
}
