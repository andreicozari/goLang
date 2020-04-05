package main

import (
	"fmt"
	"strings"
)

func main() {
	manyWords := "aa bb aa cc aa bb aa"
	fields := strings.Fields(manyWords)
	strMap := make(map[string]int)

	for _, word := range fields {
		count, exists := strMap[word]

		if exists {
			strMap[word] = count + 1
		} else {
			strMap[word] = 1
		}
	}

	fmt.Println(strMap)
}
