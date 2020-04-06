package main

import "fmt"

func my_func(args ...int) int {
	sum := 0
	for _, v := range args {
		sum = sum + v
	}

	return sum
}

func main() {
	arr := []int{2, 4}
	sum := my_func(arr...)
	fmt.Println("Sum is ", sum)
}
