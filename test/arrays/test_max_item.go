package main

import "fmt"

var count = 0

func main() {
	fmt.Println(findMax())
}

func findMax() int {

	// ann array with implicit size of 6
	nums := []int{11, 123, 0, 123, 11, 6}

	fmt.Println("The length of array is ", len(nums))

	nums[2] = 1233

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
