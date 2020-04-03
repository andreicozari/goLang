package api

import "encoding/json"

// Book type
type Book struct {
	Title  string
	Name   string
	Author string
	ISBN   string
}

// Marshall the book to json:
func (b Book) ToJSON() []byte {
	ToJson, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}
	return ToJson
}

func FromJSON() {

}
