package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Printf("Boot the app")

	http.HandleFunc("/", index)

	http.ListenAndServe(port(), nil)
}

func index(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, "Hello from a microservice written in GO lang")
}

func port() string {
	port := os.Getenv("GO_SERVICE_PORT")

	if len(port) == 0 {
		port = ":8080"
	}
	return port
}
