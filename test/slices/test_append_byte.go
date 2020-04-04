package main

import "fmt"

func main() {
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)

	// p == []byte{2, 3, 5, 7, 11, 13}

	fmt.Println(p)

	p = append(p, 17)
	fmt.Println(p)

}

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) { // if necessary, reallocate
		// allocate double what's needed, for future growth.
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}
