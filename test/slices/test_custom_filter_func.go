package main

import "fmt"

func main() {
	items := []int{2, 3, 5, 7, 11}
	fmt.Println(Filter(items, numIsEven))
}

// Filter returns a new slice holding only
// the elements of s that satisfy fn()
func Filter(s []int, fn func(int) bool) []int {
	var p []int // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

func numIsEven(n int) bool {
	return n%2 == 0
}
