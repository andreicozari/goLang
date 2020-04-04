package main

import "fmt"

type Coordinates struct {
	Lat, Long float64
}

var m map[string]Coordinates

var mapLiteral = map[string]Coordinates{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	m = make(map[string]Coordinates)
	m["Go Labs"] = Coordinates{
		40.68433, -74.39967,
	}
	fmt.Println(m["Go Labs"])

	fmt.Println(mapLiteral["Google"])
}
