package main

import "fmt"

func main() {
	fmt.Println(findMax())
}

func findMax() int {
	arr := []int{4, 6, 1, 45, 11, 123, 0, 123, 11, 6}

	max := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}
