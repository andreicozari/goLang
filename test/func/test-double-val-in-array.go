package main

import "fmt"

func doubleAt(arr []int, i int) {
	arr[i] *= 2
}

func main() {
	arr := []int{1, 2, 3, 4}
	doubleAt(arr, 2)
	fmt.Println(arr)
}
