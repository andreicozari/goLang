package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Boot the app")

	http.HandleFunc("/", index)

	http.ListenAndServe(":8080", nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, "Hello from a microservice written in GO lang")
}
