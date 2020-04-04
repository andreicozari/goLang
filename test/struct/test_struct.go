package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{3, 4}
	fmt.Println(v)

	// pointer to v:
	p := &v
	p.X = 1e9
	fmt.Println(v)

}
