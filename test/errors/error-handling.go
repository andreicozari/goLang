package main

import (
	"errors"
	"fmt"
)

type errorString struct {
	message string
}

func New(message string) error {
	return &errorString{message}
}

func (e *errorString) Error() string {
	return e.message
}

func main() {
	fmt.Println(New("This is a message from new error!"))

	fmt.Println(errors.New("EOF"))
}
