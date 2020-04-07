package main

import (
	"fmt"
	"strings"
)

var (
	manyWords = "aa bb aa cc aa bb"
	fields    = strings.Fields(manyWords)
	strMap    = make(map[string]int)
)

func main() {
	findWordFrequency()
	findMostRepeatedWord()
}

// Used in text search engines:
func findWordFrequency() {
	for _, word := range fields {
		/*
			count, exists := strMap[word]

			if exists {
				strMap[word] = count + 1
			} else {
				strMap[word] = 1
			}*/

		// can be replaced with this line only:
		// if such a entry by key does not exist then it returns 0 value:
		strMap[word]++

	}

	fmt.Println(strMap)
}

func findMostRepeatedWord() {
	max := 1
	var word string

	for k, count := range strMap {
		if count > max {
			max = count
			word = k
		}
	}

	fmt.Printf("Most repeated: %v : %v ", word, max)
}
