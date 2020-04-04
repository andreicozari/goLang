package main

import "fmt"

var description = `Slices are like references to arrays
A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.`

func main() {

	// initialize but not all items of the array:
	nums := [6]int{0, 1, 2, 3}

	var slice []int = nums[1:3]

	// the other are initialized with 0 by default:
	fmt.Println(nums)
	fmt.Println(slice)
}
