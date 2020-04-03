package main

import "fmt"

var count = 0

func main() {
	fmt.Println(findMax())
}

func findMax() int {
	nums := []int{4, 6, 1, 45, 11, 123, 0, 123, 11, 6}

	max := nums[0]

	// _ omit the index from the range:
	for _, value := range nums[1:] {
		count++
		if value > max {
			max = value
		}
	}

	fmt.Println(count)

	return max
}
