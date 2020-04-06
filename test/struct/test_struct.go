package main

import "fmt"

type Coordinate struct {
	X, Y int
}

func main() {
	fmt.Println(Coordinate{1, 2})

	v := Coordinate{3, 4}
	fmt.Println(v)

	// pointer to v:
	p := &v
	p.X = 1e9
	fmt.Println(v)

}
